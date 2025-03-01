package env

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type ValidEnv struct {
	CredentialTokenA string `env:"CREDENTIALS_TOKEN_A"`
}

func ValidateEnv() *ValidEnv {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	//exhaustruct:ignore
	vEnv := ValidEnv{}

	vEnvMeta := reflect.TypeOf(vEnv)

	for i := 0; i < vEnvMeta.NumField(); i++ {
		field := vEnvMeta.Field(i)
		fieldEnvName := field.Tag.Get("env")

		fieldEnvValue := os.Getenv(fieldEnvName)

		if len(fieldEnvValue) == 0 {
			panic(fmt.Sprintf("Field %s is required in your .env file", field.Name))
		}

		fieldValue := reflect.ValueOf(&vEnv).Elem().Field(i)
		fieldValue.Set(reflect.ValueOf(fieldEnvValue))
	}

	return &vEnv
}
