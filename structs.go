package main

type solution struct {
	Config appConfig
}

type appConfig struct {
	Exchange exchangeConfig `json:"exchange"`
}

type exchangeConfig struct {
	PublicKey string `json:"pubkey"`
	Password  string `json:"password"`
}
