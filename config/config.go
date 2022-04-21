package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config ...
type Config struct {
	Environment string         `yaml:"environment"`
	Server      ServerConfig   `yaml:"server"`
	Postgres    PostgresConfig `yaml:"postgres"`
	Logger      Logger         `yaml:"logger"`
	TelegramBot TelegramBot    `yaml:"telegramBot"`
	File        FileConfig     `yaml:"file"`
}

// Server config ...
type ServerConfig struct {
	AppVersion        string `yaml:"appVersion"`
	Port              string `yaml:"port"`
	SSL               bool   `yaml:"ssl"`
	CtxDefaultTimeout int    `yaml:"ctxDefaultTimeout"`
}

// Postgresql config ...
type PostgresConfig struct {
	PostgresqlHost     string `yaml:"postgresqlHost"`
	PostgresqlPort     string `yaml:"postgresqlPort"`
	PostgresqlUser     string `yaml:"postgresqlUser"`
	PostgresqlDbname   string `yaml:"postgresqlDbname"`
	PostgresqlPassword string `yaml:"postgresqlPassword"`
	PostgresqlSSLMode  bool   `yaml:"postgresqlSSLMode"`
}

// Logger config ...
type Logger struct {
	LogLevel string `yaml:"logLevel"`
}

// Telegram bot config
type TelegramBot struct {
	Token string `yaml:"token"`
	Url   string `yaml:"url"`
}

// File config ...
type FileConfig struct {
	Receive []string `yaml:"receive"`
}

var c Config = Config{}

// Load loads environment vars and intflates Config
func Load() Config {

	// Open config file
	file, err := os.Open("./telegram_message_config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decode := yaml.NewDecoder(file)
	if err = decode.Decode(&c); err != nil {
		log.Fatal(err)
	}

	return c
}

func CreateFile() error {
	if _, err := os.Stat("receiveMedia"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("receiveMedia", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	for _, filename := range c.File.Receive {
		if _, err := os.Stat("receiveMedia/" + filename); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir("receiveMedia/"+filename, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}
