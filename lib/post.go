package lib

// import (
// 	"bytes"
// 	"errors"
// 	"fmt"
// 	"github.com/gernest/front"
// 	"io/ioutil"
// 	"os"
// 	"time"
// )

// var (
// 	ErrEmptyBytes = errors.New("byte array is empty")
// )

// type Post struct {
// 	Date     time.Time
// 	Title    string
// 	Template string
// 	Content  string
// 	Matter   map[string]interface{}
// }

// func (p *Post) ReadFromFile(filename string) error {
// 	matter := front.NewMatter()
// 	matter.Handle("---", front.YAMLHandler)

// 	fileBytes, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return err
// 	}

// 	// If fileBytes is empty, an corresponding error will be returned
// 	fileReader := bytes.NewReader(fileBytes)
// 	if fileReader.Len() <= 0 {
// 		return ErrEmptyBytes
// 	}

// 	front, body, err := matter.Parse(fileReader)
// 	// p.Date = time.Parse(front["date"])
// 	p.Title = front["title"].(string)
// 	p.Template = front["layout"].(string)
// 	p.Content = body
// 	p.Matter = front
// 	return err
// }

// func (p *Post) Build(context Context) error {
// 	data := &struct {
// 		Title   string
// 		Content string
// 	}{Title: p.Title, Content: p.Content}

// 	if err := context.Template.ExecuteTemplate(os.Stdout, fmt.Sprintf("%s.html", p.Template), data); err != nil {
// 		return err
// 	}
// 	return nil
// }
