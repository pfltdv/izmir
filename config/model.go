package config

type IzmirConfig struct {
}

type AWSConfig struct {
}

type GCPConfig struct {
}

type Config struct {
	Izmir *IzmirConfig `yaml:"izmir"`
	AWS   *AWSConfig   `yaml:"aws"`
	GCP   *GCPConfig   `yaml:"gcp"`
}
