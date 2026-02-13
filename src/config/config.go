package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct{
	Server ServerConfig
	Postgres PostgresConfig
	Redis RedisConfig
	Password PasswordConfig
	Otp      OtpConfig
	Cors CorsConfig
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

type OtpConfig struct {
	ExpireTime time.Duration
	Digits     int
	Limiter    time.Duration
}

type ServerConfig struct{
	Port string
	RunMode string
}
type PostgresConfig struct{
	Host string
	Port string
	User string
	Password string
	DbName string
	SSLMode bool
}

type RedisConfig struct{
	Host string
	Port string
	Password string
	Db string
 	DialTimeout time.Duration 
  	ReadTimeout time.Duration
	WriteTimeout time.Duration 
  	PoolSize int 
  	PoolTimeOut time.Duration 
	IdleTimeout time.Duration
  	IdleCheckFrequency time.Duration
}

type CorsConfig struct{
	AllowOrigin string
}

func GetConfig() *Config{
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil{
		log.Fatalf("Error in load config: %v", err)
	}

	cfg, err := ParsConfig(v)
	if err != nil{
		log.Fatalf("Error in pars config: %v",err)
	}

	return cfg
}

func ParsConfig(v *viper.Viper) (*Config, error){
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil{
		log.Printf("Unable to pars config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error){
	v := viper.New()
	v.SetConfigFile(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil{
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok{
			return nil, errors.New("config file not found...")
		}
		return nil, err
	}

	return v, nil
}

func getConfigPath(env string) string{
	switch {
	case env == "docker":
		return "config/config-docker"
	case env == "production":
		return "config/config-production"
	default:
		return "../config/config-development" 
	} 
}