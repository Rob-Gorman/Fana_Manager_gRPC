package configs

import (
	"errors"
	"fmt"
	"os"

	"sovereign/utils"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load()
	utils.HandleErr(err, "Error loading .env file")
}

func getEnvVar(envKey string) string {
	val, ok := os.LookupEnv(envKey)
	if !ok {
		errmsg := missingEnvVarMsg(envKey)
		utils.HandleErr(errors.New("missing .env key"), errmsg)
	}
	return val
}

// return type []interface... expected by Sprintf
// (Does NOT want type []string...)
func getEnvVars(envVars ...string) (result []interface{}) {
	for _, envKey := range envVars {
		val := getEnvVar(envKey)
		result = append(result, val)
	}
	return result
}

func missingEnvVarMsg(variableName string) string {
	return fmt.Sprintf("No %s variable found in environment. Verify .env file.", variableName)
}
