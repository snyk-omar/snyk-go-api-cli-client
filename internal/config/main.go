package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type YamlConfig struct {
	Config struct {
		Org    string `yaml:"org"`
		Apikey string `yaml:"apikey"`
	}
}

func verifyFileName(fileName string) {
	fileExtension := filepath.Ext(fileName)

	if fileExtension != ".yaml" {
		panic("File extension is not supported: " + fileExtension)
	}
}

func Config(fileName string) YamlConfig {

	verifyFileName(fileName)

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n\n", err)
	}

	var yamlConfig YamlConfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n\n", err)
	}

	return yamlConfig
}
