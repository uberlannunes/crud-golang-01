package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uberlannunes/crud-golang-01/configuration/database/mongodb"
	"github.com/uberlannunes/crud-golang-01/configuration/logger"
	"github.com/uberlannunes/crud-golang-01/controller/routes"
	"github.com/uberlannunes/crud-golang-01/setup"
)

func main() {

	logger.Info("starting application...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init dependencies

	// mongodb.InitConnection()
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, errors=%s", err.Error())
		return
	}

	fmt.Printf("db=%v", database)
	fmt.Println()

	userController := setup.InitDependencies(database)
	// userService := service.NewUserDomainService()
	// userController := controller.NewUserController(userService)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
