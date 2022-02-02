package main

import "github.com/Sagleft/uexchange-go"

type solution struct {
	Config appConfig

	ExchangeClient *uexchange.Client
}

type appConfig struct {
	Exchange exchangeConfig `json:"exchange"`
}

type exchangeConfig struct {
	PublicKey string `json:"pubkey"`
	Password  string `json:"password"`
}
