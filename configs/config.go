package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

//Config...
type Config struct {
	HTTPPort         string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
}

//Load...
func Load() (c Config) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	c = Config{
		HTTPPort:         cast.ToString(env("HTTP_PORT")),
		PostgresHost:     cast.ToString(env("POSTGRES_HOST")),
		PostgresPort:     cast.ToInt(env("POSTGRES_PORT")),
		PostgresDatabase: cast.ToString(env("POSTGRES_DATABASE")),
		PostgresUser:     cast.ToString(env("POSTGRES_USER")),
		PostgresPassword: cast.ToString(env("POSTGRES_PASSWORD")),
	}

	return
}

func (c *Config) PostgresURL() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDatabase)
}

func env(key string) interface{} {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}

	log.Fatal(fmt.Sprintf("couldn't find environment key: %v", key))

	return nil
}
