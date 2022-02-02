package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := newSolution()
	err := app.run()
	if err != nil {
		log.Fatalln(err)
	}
}

func newSolution() *solution {
	return &solution{}
}

func (sol *solution) run() error {
	err := sol.setupRouter()
	if err != nil {
		return err
	}
	return nil
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
