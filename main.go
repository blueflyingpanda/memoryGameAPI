package main

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var BotDB *gorm.DB
var DB *gorm.DB

// @title           Player API
// @version         1.0
// @description     This is a sample server for player management.
// @host      localhost:8080
// @BasePath  /
// @schemes http
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	botDBHost := os.Getenv("BOT_DB_HOST")
	botDBName := os.Getenv("BOT_DB_NAME")
	botDBUser := os.Getenv("BOT_DB_USER")
	botDBPass := os.Getenv("BOT_DB_PASS")

	botDBPort, err := strconv.Atoi(os.Getenv("BOT_DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	BotDB = initDB(botDBHost, botDBName, botDBUser, botDBPass, botDBPort)

	DBHost := os.Getenv("DB_HOST")
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	DBPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	DB = initDB(DBHost, DBName, DBUser, DBPass, DBPort)

	initAPI(8080)
}
