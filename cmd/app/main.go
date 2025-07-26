package app

import (
	"log"

	"github.com/Agu-GC/bitso/internal/infraestructure/myHanddler"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("Levantando server HTTP")
	router := gin.Default()

	router.GET("/", myHanddler.Home)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting gin server")
		panic("Error starting gin server " + err.Error())
	}
	log.Print("Server started")

}
