package main

import (
	"errors"
	"log"

	"github.com/Sagleft/uexchange-go"
	"github.com/gin-gonic/gin"
)

func main() {
	app := newSolution()
	err := checkErrors(
		app.parseConfig,
		app.initGin,
		app.connectExchange,
		app.setupRoutes,
		app.runGin,
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func newSolution() *solution {
	return &solution{}
}

func (sol *solution) initGin() error {
	sol.Gin = gin.Default()
	return nil
}

func (sol *solution) runGin() error {
	return sol.Gin.Run()
}

func (sol *solution) connectExchange() error {
	//sol.ExchangeClient
	// TODO
	return nil
}

func (sol *solution) getExchangePairs() ([]uexchange.PairsDataContainer, error) {
	if sol.ExchangeClient == nil {
		return nil, errors.New("exchange client is not set")
	}
	return sol.ExchangeClient.GetPairs()
}
