package config

import (
	"log"
	"os"

	"github.com/TheTeemka/hhChat/pkg/validator"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBString   string
}

func MustLoad() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("could not load .env file: %s", err)
	}

	cfg := &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		DBString:   os.Getenv("PSQL_DBSTRING"),
	}

	v := validator.New()
	if cfg.Validate(v); !v.Valid() {
		log.Fatal(v)
	}

	return cfg
}

func (cfg *Config) Validate(v *validator.Validator) {
	v.CheckWithRules("Config DBString", cfg.DBString, validator.IsNotEmpty)

	v.CheckWithRules("Server Port", cfg.ServerPort, validator.IsNotEmpty)
	v.CheckWithRules("Server Port", cfg.ServerPort[1:], validator.IsInt(0))
}
