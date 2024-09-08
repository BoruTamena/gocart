package main

import (
	"log"

	"github.com/BoruTamena/infra/repository"
	"github.com/BoruTamena/internal/core/service"
	"github.com/BoruTamena/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	Router := gin.Default()

	db, err := repository.NewDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	cart_rep := repository.NewCartRepository(db)

	cart_service := service.NewCartService(cart_rep)

	cart_handler := handler.NewCartHandler(Router, cart_service)

	cart_handler.InitHandler()

	defer db.Close()

	// serving the server on port 8000
	if err := Router.Run(":8000"); err != nil {

		log.Print(err.Error())
	}

}
