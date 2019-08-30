package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Hosts []Host `yaml:",flow"`
	Keys  []Key  `yaml:",flow"`
}

type Host struct {
	Name       string
	Connection struct {
		Hostname string
		Port     string
	}
	Credentials struct {
		User string
		Key  string
	}
}

type Key struct {
	Name string
	Path string
}

var config = loadConfig()

func loadConfig() Config {
	var config = Config{}
	path := filepath.Join(os.Getenv("HOME"), ".essh")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}

	configPath := filepath.Join(path, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		file, _ := os.Create(configPath)
		_, _ = file.WriteString("\x68\x6f\x73\x74\x73\x3a\x0a\x20\x20\x2d\x20\x6e\x61\x6d\x65\x3a\x20\x22\x68\x6f\x73\x74\x4e\x61\x6d\x65\x22\x0a\x20\x20\x20\x20\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x3a\x0a\x20\x20\x20\x20\x20\x20\x68\x6f\x73\x74\x6e\x61\x6d\x65\x3a\x20\x22\x31\x2e\x31\x2e\x31\x2e\x31\x22\x0a\x20\x20\x20\x20\x20\x20\x70\x6f\x72\x74\x3a\x20\x22\x32\x32\x22\x0a\x20\x20\x20\x20\x63\x72\x65\x64\x65\x6e\x74\x69\x61\x6c\x73\x3a\x0a\x20\x20\x20\x20\x20\x20\x75\x73\x65\x72\x3a\x20\x22\x75\x73\x65\x72\x22\x0a\x20\x20\x20\x20\x20\x20\x6b\x65\x79\x3a\x20\x22\x6b\x65\x79\x4e\x61\x6d\x65\x22\x0a\x6b\x65\x79\x73\x3a\x0a\x20\x20\x2d\x20\x6e\x61\x6d\x65\x3a\x20\x22\x6b\x65\x79\x4e\x61\x6d\x65\x22\x0a\x20\x20\x20\x20\x70\x61\x74\x68\x3a\x20\x22\x7e\x2f\x6b\x65\x79\x73\x2f\x6b\x65\x79\x66\x69\x6c\x65\x2e\x70\x65\x6d\x22")
		_ = file.Close()
	}
	file, err := ioutil.ReadFile(configPath)

	err = yaml.UnmarshalStrict(file, &config)
	if err != nil {
		fmt.Println("Failed to Parse Configuration File: ", err)
	}

	return config
}

func GetHostByName(hostname string) (Host, int) {
	for index, host := range config.Hosts {
		if host.Name == hostname {
			return host, index
		}
	}
	return Host{}, -1
}

func GetKeyByName(keyname string) (Key, int) {
	for index, key := range config.Keys {
		if key.Name == keyname {
			return key, index
		}
	}
	return Key{}, -1
}
