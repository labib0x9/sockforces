package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Service string
	Addr    string
	Port    int
}

var (
	configuration *Config
	once          sync.Once
)

func loadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		if !os.IsNotExist(err) {
			log.Panic(err)
		}
	}

	fn := func(name string) string {
		value := os.Getenv(name)
		if value == "" {
			log.Panic(name)
		}
		return value
	}

	port, err := strconv.Atoi(fn("PORT"))
	if err != nil {
		log.Fatalln(err)
	}

	configuration = &Config{
		Service: fn("SERVICE"),
		Addr:    fn("ADDR"),
		Port:    port,
	}
}

func GetConfig() *Config {
	once.Do(func() {
		loadConfig()
	})
	return configuration
}
