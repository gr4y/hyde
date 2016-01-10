package lib

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/gernest/front"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"
)

const (
	SEPARATOR      = "---"
	SLUG_SEPARATOR = "-"
)

var (
	rControl   = regexp.MustCompile("[\u0000-\u001f]")
	rSpecial   = regexp.MustCompile("[\\s~`!@#\\$%\\^&\\*\\(\\)\\-_\\+=\\[\\]\\{\\}\\|\\;:\"'<>,\\.\\?\\/]+")
	rRepeatSep = regexp.MustCompile(SLUG_SEPARATOR + "{2,}")
	rEdgeSep   = regexp.MustCompile("^" + SLUG_SEPARATOR + "+|" + SLUG_SEPARATOR + "+$")
	// Errors
	ErrEmptyBytes = errors.New("Can't create content, because file was empty.")
	ErrEmptyTitle = errors.New("The `title` field of content can't be empty.")
	ErrNoTemplate = errors.New("The `template` field of content can't be empty.")
)

type Content struct {
	Title      string
	Slug       string
	ContentRaw string
	Content    string
	Template   string
	Date       time.Time
}

func (c *Content) InitFromFile(path string) error {
	b, err := ioutil.ReadFile(path)
	err = c.InitFromBytes(b)
	return err
}

func (c *Content) InitFromBytes(b []byte) error {
	if len(b) == 0 {
		return ErrEmptyBytes
	}
	front, body, err := parseYAMLFrontMatter(b)
	err = c.initWithFrontMatterAndContent(front, body)
	return err
}

// Builds and Writes Content to Disk
func (c *Content) Write(bc BuildContext) error {
	data := &struct {
		Content       Content
		Configuration Configuration
	}{Content: *c, Configuration: bc.Configuration}

	if f, err := os.Create(fmt.Sprintf("%s/%s.html", bc.Configuration.OutputPath, c.Slug)); err != nil {
		defer f.Close()
		w := bufio.NewWriter(f)
		if err := bc.Template.ExecuteTemplate(w, fmt.Sprintf("%s.html", c.Template), data); err != nil {
			return err
		}
	}

	return nil
}

/************** PRIVATE METHODS **************/
func (c *Content) initWithFrontMatterAndContent(fm map[string]interface{}, body string) error {
	err := c.initTitle(fm)
	err = c.initSlug(fm)
	err = c.initTemplate(fm)
	err = c.initDate(fm)
	err = c.initContent(body)
	return err
}

func (c *Content) initTitle(fm map[string]interface{}) error {
	title := fm["title"].(string)
	if len(title) > 0 {
		c.Title = title
	} else {
		return ErrEmptyTitle
	}
	return nil
}

func (c *Content) initTemplate(fm map[string]interface{}) error {
	template := fm["layout"].(string)
	if len(template) > 0 {
		c.Template = template
	} else {
		return ErrNoTemplate
	}
	return nil
}

func (c *Content) initSlug(fm map[string]interface{}) error {
	if fm["slug"] != nil {
		c.Slug = fm["slug"].(string)
	} else {
		s, err := slugize(c.Title)
		if err != nil {
			return err
		}
		c.Slug = s
	}
	return nil
}

func (c *Content) initDate(fm map[string]interface{}) error {
	return nil
}

func (c *Content) initContent(body string) error {
	if len(body) > 0 {
		c.ContentRaw = body
		c.Content = string(convertMarkdownToHTML(body))
	} else {

	}
	return nil
}

/************** HELPER METHODS **************/

// Parsing the YAML Frontmatter
func parseYAMLFrontMatter(b []byte) (map[string]interface{}, string, error) {
	matter := front.NewMatter()
	matter.Handle(SEPARATOR, front.YAMLHandler)
	reader := bytes.NewReader(b)
	return matter.Parse(reader)
}

// Convert a date (string) into a time.Time object
func parseDate(date string) (time.Time, error) {
	//	time.Parse(JEKYLL, value)
	return time.Now().UTC(), nil
}

// Generate a slug from a string, removing umlauts and weird characters.
func slugize(title string) (string, error) {
	str := escapeDiacritic(title)
	str = rControl.ReplaceAllString(str, SLUG_SEPARATOR)
	str = rSpecial.ReplaceAllString(str, SLUG_SEPARATOR)
	str = rRepeatSep.ReplaceAllString(str, SLUG_SEPARATOR)
	str = rEdgeSep.ReplaceAllString(str, "")
	return strings.ToLower(str), nil
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func escapeDiacritic(str string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, str)
	return result
}

// Converts the content into markdown
func convertMarkdownToHTML(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return string(html)
}
