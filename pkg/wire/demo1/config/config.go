package config

import "fmt"

type Config struct {
	DSN string `json:"dsn"`
}

func NewDefaultConfig(path string) *Config {
	fmt.Println("config path:", path)
	// read and parse
	return &Config{
		DSN: "root:123456@/goctest",
	}
}
