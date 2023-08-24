package config

import (
	"fmt"
	"os"

	"github.com/pfltdv/izmir/util"
)

const CONFIG_ENV_PARAM string = "IZMIR_CONFIG"

func Load(path string) (*Config, error) {
	if path == "" {
		path = os.Getenv(CONFIG_ENV_PARAM)
	}
	if path == "" {
		return nil, fmt.Errorf("invalid configuration")
	}
	config := &Config{}
	err := util.Unmarshall(path, config)
	if err != nil {
		return nil, err
	}
	return config, err
}
