package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go_fiber/internal/app"
	"go_fiber/logger"
	"os"
)

func main() {
	env := os.Getenv("env")

	fmt.Print(env)
	if env == "dev" {
		err := godotenv.Load(".env.dev")
		if err != nil {
			logger.Error("Error loading .env.dev file :: ", zap.Error(err))
			panic("Error loading .env.dev file")
		}
	} else if env == "production" {
		err := godotenv.Load(".env")
		if err != nil {
			logger.Error("Error loading .env file :: ", zap.Error(err))
			panic("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")
	appInstance := app.Initialize()

	if err := appInstance.Listen(":" + port); err != nil {
		logger.Error("Error starting application :: ", zap.Error(err))
		panic(err)
	}
}
