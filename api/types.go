package api

import "net/http"

type coins struct {
	Coins []string `json:"coins"`
}

type coinData struct {
	Time   int    `json:"time"`
	Coin   string `json:"coin"`
	Source string `json:"source"`
}

type coinDataList struct {
	Coins []coinData
}

type APIServer struct {
	listenAddr string
}

type APIError struct {
	Error string
}

type APIFunc func(http.ResponseWriter, *http.Request) error
