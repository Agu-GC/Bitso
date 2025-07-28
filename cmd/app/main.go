package main

import (
	"log"

	"github.com/Agu-GC/Bitso/internal/application"
	"github.com/Agu-GC/Bitso/internal/infraestructure"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("Levantando server HTTP")
	router := gin.Default()

	userHanddler := infraestructure.UserHanddler{
		UserService: &application.UserService{},
	}

	router.GET("/user/:id", userHanddler.GetUser)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting gin server")
		panic("Error starting gin server " + err.Error())
	}
	log.Print("Server started")

}
