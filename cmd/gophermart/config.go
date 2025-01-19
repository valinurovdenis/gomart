package main

import (
	"flag"
	"os"
	"reflect"
)

type Config struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	DatabaseURI          string `env:"DATABASE_URI"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	LogLevel             string `env:"LOG_LEVEL"`
	SecretKey            string `env:"SECRET_KEY"`
	UpdateThreads        string `env:"UPDATE_THREADS"`
	AccrualRetries       int    `env:"ACCRUAL_RETRIES"`
	AccrualDelay         int    `env:"ACCRUAL_DELAY"`
	AccrualTimeout       int    `env:"ACCRUAL_TIMEOUT"`
}

func parseFlags(config *Config) {
	flag.StringVar(&config.RunAddress, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&config.DatabaseURI, "d", "", "database uri")
	flag.StringVar(&config.AccrualSystemAddress, "r", "", "accrual system address")
	flag.StringVar(&config.LogLevel, "l", "info", "log level")
	flag.StringVar(&config.SecretKey, "k", "SECRET_KEY", "secret key")
	flag.StringVar(&config.UpdateThreads, "t", "UPDATE_THREADS", "update threads number")
	flag.IntVar(&config.AccrualRetries, "x", 3, "number of retries to accrual service")
	flag.IntVar(&config.AccrualDelay, "y", 500, "delay in ms between retries to accrual service")
	flag.IntVar(&config.AccrualTimeout, "z", 1000, "timeout in ms to accrual service")
	flag.Parse()
}

func (config *Config) updateFromEnv() {
	v := reflect.Indirect(reflect.ValueOf(config))
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		var envName string
		if envName = field.Tag.Get("env"); envName == "" {
			continue
		}
		if envVal := os.Getenv(envName); envVal != "" {
			v.Field(i).SetString(envVal)
		}
	}
}
