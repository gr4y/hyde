package lib

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type BuildContext struct {
	Configuration Configuration
	Template      *template.Template
}

func (bc *BuildContext) Build() error {
	t, err := bc.parseTemplates()
	if err != nil {
		return err
	}
	bc.Template = t
	contents, err := bc.getContents()
	if err != nil {
		return err
	}
	for _, content := range contents {
		if err := content.Write(*bc); err != nil {
			return err
		}
	}
	return nil
}

func (bc *BuildContext) getContents() ([]Content, error) {
	var contents []Content
	files, err := filepath.Glob(fmt.Sprintf("%s/**/*", bc.Configuration.ContentPath))
	for _, f := range files {
		content := Content{}
		err := content.InitFromFile(f)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error occured while building `%s`\n\r  %s", f, err))
		}
		contents = append(contents, content)
	}
	return contents, err
}

func (bc *BuildContext) parseTemplates() (*template.Template, error) {
	templateList := []string{}
	if err := filepath.Walk(bc.Configuration.TemplatePath, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			templateList = append(templateList, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return template.ParseFiles(templateList...)
}
