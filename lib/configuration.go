package lib

import (
	"encoding/json"
	"io/ioutil"
)

const (
	CONFIG_FILENAME = "config.json"
)

type Configuration struct {
	OutputPath   string
	TemplatePath string
	ContentPath  string
	AssetsPath   string
}

func (c *Configuration) Read() error {
	bytes, err := ioutil.ReadFile(CONFIG_FILENAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &c)
}
