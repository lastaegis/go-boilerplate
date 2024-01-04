package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go-boilerplate/routing"
	"log"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil{
		log.Panic(err)
	}
	
	// Init Echo
	echoFramework := echo.New()
	routing.RoutingAPI(echoFramework)
}