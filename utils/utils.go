package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatal(msg)
	}
}

func LoadDotEnv() {
	err := godotenv.Load()
	HandleErr(err, "Error loading .env file")
}

func Port() int {
	envKey := "PORT"
	portString, ok := os.LookupEnv(envKey)
	if !ok {
		return 8080
	}

	port, err := strconv.Atoi(portString)
	errmsg := missingEnvVarMsg(envKey)
	HandleErr(err, errmsg)

	return port
}

func DBURI() string {
	envKey := "DB_URI"
	uri, ok := os.LookupEnv(envKey)
	if !ok {
		errmsg := missingEnvVarMsg(envKey)
		HandleErr(errors.New("missing .env key"), errmsg)
	}
	return uri
}

func DBName() string {
	envKey := "DB_NAME"
	db, ok := os.LookupEnv(envKey)
	if !ok {
		errmsg := missingEnvVarMsg(envKey)
		HandleErr(errors.New("missing .env key"), errmsg)
	}
	return db
}

func missingEnvVarMsg(variableName string) string {
	return fmt.Sprintf("No %s variable found in environment. Verify .env file.", variableName)
}
