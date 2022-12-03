package api

import (
	"net/http"

	"github.com/conghaile/coincrowd-API/db"
)

type coinsResponse struct {
	Coins []string `json:"coins"`
}

type coinDataResponse struct {
	Time   int    `json:"time"`
	Source string `json:"source"`
}

type coinDataListResponse struct {
	Coins []coinDataResponse
}

type timeFrame struct {
	Weeks int64
}

type APIServer struct {
	listenAddr string
	store      db.Storage
}

type APIError struct {
	Error string
}

type APIFunc func(http.ResponseWriter, *http.Request) error
