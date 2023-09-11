package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig
	Redis    RedisConfig
	Server   ServerConfig
	Otp      OtpConfig
	Logger   LoggerConfig
	Jwt      JwtConfig
	Password Password
}

type ServerConfig struct {
	Port    int
	Host    string
	RunMode string
}

type PostgresConfig struct {
	Port            int
	Host            string
	User            string
	Password        string
	DbName          string
	SSlMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Port               int
	Host               string
	Password           string
	Db                 int
	PoolSize           int
	MinIdleConnections int
	PoolTimeout        time.Duration
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
}

type LoggerConfig struct {
	Level    string
	FilePath string
	Encoding string
	Logger   string
}
type JwtConfig struct {
	Secret                     string
	RefreshSecret              string
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
}

type Password struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

type OtpConfig struct {
	Digits     int
	ExpireTime time.Duration
	Limiter    time.Duration
}

func getConfigPath(env string) (path string) {
	if env == "docker" {
		return "/app/config/config-docker.yml"
	} else if env == "production" {
		return "./config/config-docker.yml/config-production.yml"
	} else {
		return "../config/config-development.yml"
	}
}

func loadConfig(fileP, fileExtension string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileExtension)
	v.SetConfigName(fileP)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config File NotFound Error")
		}
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func GetConfig() (cfg *Config) {
	fileP := getConfigPath("APP_ENV")
	v, err := loadConfig(fileP, "yml")
	if err != nil {
		log.Fatalf("cant load config\n %s", err.Error())
		return
	}
	cfg, err = ParseConfig(v)
	if err != nil {
		log.Fatalf("Error in parse cfg\n %s", err.Error())
		return
	}
	return cfg
}
