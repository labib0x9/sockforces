package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Github struct {
	AppId          int64
	InstalationId  int64
	PrivateKeyPath string
	Org            string
}

type Config struct {
	Service string
	Addr    string
	Port    int
	Github  *Github
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

	gAppId, err := strconv.ParseInt(fn("APP_ID"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	gInsId, err := strconv.ParseInt(fn("INSTALLAR_ID"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	configuration = &Config{
		Service: fn("SERVICE"),
		Addr:    fn("ADDR"),
		Port:    port,
		Github: &Github{
			AppId:          gAppId,
			InstalationId:  gInsId,
			PrivateKeyPath: fn("PRIVATE_KEY_PATH"),
			Org: fn("ORGANIZATION"),
		},
	}
}

func GetConfig() *Config {
	once.Do(func() {
		loadConfig()
	})
	return configuration
}
