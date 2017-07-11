package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigurationData struct {
	Address string `yaml:"http.address"`
}

func Read(path string) (*ConfigurationData, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var configurationData ConfigurationData
	if err = yaml.Unmarshal(content, &configurationData); err != nil {
		return nil, err
	}
	return &configurationData, nil
}
