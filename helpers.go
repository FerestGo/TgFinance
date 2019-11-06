package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
)

type Config map[string]string

func getConfig() Config {
	configFile := flag.String("env", ".env", "environment file path")
	flag.Parse()
	log.Println("CONFIG", *configFile)
	config, _ := godotenv.Read(*configFile)
	return config
}
