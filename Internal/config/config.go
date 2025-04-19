package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config{
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == ""{
		flags := flag.String("config" , "", "path to the config file")
		flag.Parse()

		configPath = *flags

		if configPath == ""{
            log.Fatal("Config path is not set..")
		}	
	}

	_, err := os.Stat(configPath)
	if os.IsNotExist(err){
		log.Fatalf("config files doest not exist: %s", configPath)
	}

	var cfg Config

	er := cleanenv.ReadConfig(configPath, &cfg)
	if er == nil{
		log.Fatalf("can not read the config file %s", err.Error())
	}

	return &cfg

}