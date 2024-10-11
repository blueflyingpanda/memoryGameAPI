package main

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var BotDB *gorm.DB

// @title           Player API
// @version         1.0
// @description     This is a sample server for player management.
// @host      d5dsv84kj5buag61adme.apigw.yandexcloud.net
// @BasePath  /
// @schemes https
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

	initAPI(8080)
}
