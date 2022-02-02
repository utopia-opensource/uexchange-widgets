package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := newSolution()
	err := checkErrors(
		app.connectExchange,
		app.setupRouter,
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func newSolution() *solution {
	return &solution{}
}

func (sol *solution) setupRouter() error {
	router := router{
		Gin: gin.Default(),
	}
	err := router.setupRoutes()
	if err != nil {
		return err
	}
	return router.Gin.Run()
}

func (sol *solution) connectExchange() error {
	//sol.ExchangeClient
	// TODO
	return nil
}
