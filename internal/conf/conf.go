package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Conf struct {
	Old struct {
		Hostname        string
		PermalinkPrefix string `yaml:"permalink-prefix"`
	}
	New struct {
		Hostname        string
		PermalinkPrefix string `yaml:"permalink-prefix"`
	}
	OutputPath string `yaml:"output-path"`
	Auth       struct {
		Username string
		Password string
	}
}

func ParseConf(confPath string) {
	// open the file and save as byte stream
	confBytes, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(confBytes, &Conf)
	if err != nil {
		panic(err)
	}

	if Conf.Old.PermalinkPrefix[0] != '/' {
		Conf.Old.PermalinkPrefix = "/" + Conf.Old.PermalinkPrefix
	}
	if Conf.New.PermalinkPrefix[0] != '/' {
		Conf.New.PermalinkPrefix = "/" + Conf.New.PermalinkPrefix
	}

	if Conf.Old.PermalinkPrefix[len(Conf.Old.PermalinkPrefix)-1] != '/' {
		Conf.Old.PermalinkPrefix += "/"
	}
	if Conf.New.PermalinkPrefix[len(Conf.New.PermalinkPrefix)-1] != '/' {
		Conf.New.PermalinkPrefix += "/"
	}
}
