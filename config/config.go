package config

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Server struct {
	Address string
	Port string

}


type Config struct {
	Server Server `yaml:"http_server" env-required`
	DBUsername string `env:"DB_USERNAME" env-required`
	DBPassword string `env:"DB_PASSWORD" env-required`
	DBName string `env:"DB_NAME" env-required`
	JWTSecret string `env:"JWT_SECRET" env-required`
	JWTExpiryInSeconds int64 `env:"JWT_EXPIRY" env-required`
}
func InitConfig() (*Config, error) {

	godotenv.Load()

	// getting configuration path from .env file
	configurationPath := os.Getenv("CONFIG_PATH");

	// getting configuration path from -config flag
	if configurationPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		configurationPath = *flags

		if configurationPath == "" {
			log.Fatal("Config file path not found!!!")
		}
		
	}

	// checking if config file exists at the path or not
	_, err:= os.Stat(configurationPath); if os.IsNotExist(err) {
		log.Fatalf("Config file at %s does not exist!!!", configurationPath);
	}

	// mapping configuration to config

	var config Config

	err = cleanenv.ReadConfig(configurationPath, &config)

	if err!= nil {
		log.Fatalf("Error in config file %s: %v", configurationPath, err)
    }

	err = cleanenv.ReadEnv(&config)

	if err!= nil {
        log.Fatalf("Error in environment variables: %v", err)
    }

    if config.Server.Address == "" {
        return nil, errors.New("hostAddr is required in the config")
    }

	if config.Server.Port == "" {
        return nil, errors.New("port is required in the config")
    }

	if config.DBUsername == "" {
        return nil, errors.New("dbUsername is required in the config")
    }

	if config.DBPassword == "" {
        return nil, errors.New("dbPassword is required in the config")
    }

	if config.DBName == "" {
        return nil, errors.New("dbName is required in the config")
    }

	if config.JWTSecret == "" {
        return nil, errors.New("jwtSecret is required in the config")
    }

	if config.JWTExpiryInSeconds <= 0 {
        return nil, errors.New("jwtExpiryInSeconds is required in the config and should be a positive integer")
    }

	return &config, nil

}

