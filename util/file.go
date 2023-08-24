package util

import (
	"os"

	"gopkg.in/yaml.v3"
)

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func Unmarshall(path string, obj interface{}) error {
	yamlFile, err := os.ReadFile(path)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, obj)
	}
	return err
}
