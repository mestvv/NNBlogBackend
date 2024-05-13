package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env" env-required:"true"`
	HttpServer HttpServer `yaml:"http_server"`
	Database   Database   `yaml:"database"`
	Limiter    Limiter    `yaml:"limiter"`
}

type HttpServer struct {
	Port           string        `yaml:"port" env-default:"8080"`
	Timeout        time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout    time.Duration `yaml:"iddle_timeout" env-default:"60s"`
	SwaggerEnabled bool          `yaml:"swagger_enabled" env-default:"false"`
}

type Database struct {
	Net                string        `yaml:"net" env-default:"tcp"`
	Server             string        `yaml:"server" env-required:"true"`
	DBName             string        `yaml:"db_name" env-required:"true"`
	User               string        `yaml:"user" env:"mysql_user" env-required:"true"`
	Password           string        `yaml:"password" env:"mysql_password" env-required:"true"`
	TimeZone           string        `yaml:"time_zone"`
	Timeout            time.Duration `yaml:"timeout" env-default:"2s"`
	MaxIdleConnections int           `yaml:"max_idle_connections" env-default:"40"`
	MaxOpenConnections int           `yaml:"max_open_connections" env-default:"40"`
}

type Limiter struct {
	RPS   int           `yaml:"rps" env-default:"10"`
	Burst int           `yaml:"burst" env-default:"20"`
	TTL   time.Duration `yaml:"ttl" env-default:"10m"`
}

func MustLoad(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file doesn't exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can not read config: %s", err)
	}

	return &cfg
}
