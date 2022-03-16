package repository

import (
	"log"
	"os"
	"strconv"
)

var (
	Host_address        string
	Default_code_length int
)

func Init() {
	Host_address = os.Getenv("HOST_ADDRESS")

	// Parse
	var err error
	Default_code_length, err = strconv.Atoi(os.Getenv("DEFAULT_CODE_LENGTH"))
	if err != nil {
		log.Fatal("DEFAULT_CODE_LENGTH should be integer")
	}
}
