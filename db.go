package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID       uint   `gorm:"primarykey" json:"-"`
	TgId     uint   `gorm:"unique;not null" json:"-"`
	Username string `gorm:"unique;not null"`
	Name     string `gorm:"not null"`
}

type Player struct {
	gorm.Model `json:"-"`
	Login      string `gorm:"unique;not null"`
	Password   string `gorm:"not null" json:"-"`
	Score      uint   `gorm:"not null;default:0"`
}

func initDB(host, dbName, dbUser, dbPass string, port int) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require TimeZone=Europe/Moscow",
		host, port, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(&Player{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}
