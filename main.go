package main

import (
	"log"

	"github.com/MarcoManjarrez/VinylStore.git/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/createAccount", handlers.CreateAccount)
	r.POST("/login", handlers.Login)

	protected := r.Group("/api")
	protected.Use(handlers.CheckAuthorization()) //"Proteje" a los endpoints. Usa a checkAuth para ver si el token que uso el usuario es correcto
	{
		protected.GET("/vinyls", handlers.GetVinylRegistry)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
