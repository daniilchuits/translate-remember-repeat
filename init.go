package main

import (
	"fmt"
	"log"
	"os"
	"translate-remember-repeat/internal/domain"

	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Error loading '.env': %s", err))
	}

	switch {
	case os.Getenv("REDIS_CNN_STR") == "":
		log.Panicf("%s: 'REDIS_CNN_STR'", domain.EmptyVariableInEnv)
	case os.Getenv("DEEPL_AUTH_KEY") == "":
		log.Panicf("%s: 'DEEPL_AUTH_KEY'", domain.EmptyVariableInEnv)
	}
}
