package config

import (
	"log"
	"os"
)

var (
	DBURL         string
	SessionSecret string
	ServerPort    string
)

func Load() {
	DBURL = os.Getenv("DATABASE_URL")
	if DBURL == "" {
		log.Fatal("Missing DATABASE_URL in environment")
	}

	SessionSecret = os.Getenv("SESSION_SECRET")
	if SessionSecret == "" {
		log.Fatal("Missing SESSION_SECRET in environment")
	}

	ServerPort = os.Getenv("PORT")
}
