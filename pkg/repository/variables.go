package repository

import (
	"os"

	util "github.com/beenson/URL_Shortener/pkg/utils"
)

var (
	Host_address        string
	Default_code_length int
	Maximum_tries       int
	Cache_code_ttl      int
)

func Init() {
	Host_address = os.Getenv("HOST_ADDRESS")

	// Parse
	util.ConvertEnvToInt(&Default_code_length, "DEFAULT_CODE_LENGTH", 5)
	util.ConvertEnvToInt(&Maximum_tries, "MAXIMUM_CODE_TRIES", 5)
	util.ConvertEnvToInt(&Cache_code_ttl, "CACHE_CODE_TTL", 600)
}
