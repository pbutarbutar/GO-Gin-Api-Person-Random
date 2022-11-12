package main

import (
	"api-random-user/config"
	httpDelivery "api-random-user/src/app/person/delivery"
	"api-random-user/src/app/person/repository"
	"api-random-user/src/app/person/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()
	Port := os.Getenv("API_PORT")

	apiCfg := config.NewApiPersonConfig()
	personRepo := repository.NewPersonUseCase(apiCfg)
	personUseCase := usecase.NewPersonUseCase(personRepo)
	httpDelivery.NewPersonHandler(router, personUseCase)

	log.Fatal(router.Run(Port))
}
