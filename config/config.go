package config

import (
	"log"
	"net/url"
	"os"
)

func MustGetenv(name string) string {
	val := os.Getenv(name)
	if val == "" {
		log.Fatalf("FATAL: Environment variable %s is not set!", name)
	}
	return val
}

func GetenvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val != "" {
		return val
	}
	return defaultValue
}

func RequireInProduction(name, defaultValue string) string {
	if IsProd {
		return MustGetenv(name)
	}
	return GetenvWithDefault(name, defaultValue)
}

var AppEnv = GetenvWithDefault("APP_ENV", "dev")
var IsProd = AppEnv == "prod"

var DatabaseUrl = RequireInProduction("DATABASE_URL", "./dev.db")
var DatabaseToken = RequireInProduction("DATABASE_TOKEN", "")

func DatabaseSettings() (driver, connString string) {
	if IsProd {
		qs := url.Values{
			"authToken": {DatabaseToken},
		}

		return "turso-serverless", DatabaseUrl + "?" + qs.Encode()
	}

	return "turso", DatabaseUrl
}
