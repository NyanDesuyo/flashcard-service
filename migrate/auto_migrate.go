package main

import (
	"github.com/NyanDesuyo/flashcard-service/config"
	"github.com/NyanDesuyo/flashcard-service/models"
)

func init() {
	config.LoadEnv()
	config.ConnectMainPostgres()
}

func main() {
	config.MainDB.AutoMigrate(&models.User{})
}
