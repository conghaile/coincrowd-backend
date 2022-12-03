package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/conghaile/coincrowd-API/db"
	"github.com/gorilla/mux"
)

const weekSeconds int64 = 604800

func NewAPIServer(listenAddr string, store db.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/all", makeHTTPHandlerFunc(s.CoinsHandler))
	router.HandleFunc("/coindata/{coin}", makeHTTPHandlerFunc(s.CoinDataHandler))
	router.HandleFunc("/sourcedata/{source}", makeHTTPHandlerFunc(s.SourceHandler))

	log.Println("Server running at port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// Endpoint that returns JSON of all coins on record to frontend
// /all

func (s *APIServer) CoinsHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		coinlist, err := s.store.GetAllCoins()
		if err != nil {
			return err
		}

		coinlistResponse := convertToCoinsResponse(coinlist)

		WriteJSON(w, http.StatusOK, coinlistResponse)
		return nil

	}
	return fmt.Errorf("Method not allowed: %s", r.Method)
}

// Endpoint that returns JSON of all data for a given coin to frontend
// /{coin}

func (s *APIServer) CoinDataHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		coin := mux.Vars(r)["coin"]
		var timeframe timeFrame
		err := json.NewDecoder(r.Body).Decode(&timeframe)
		if err != nil {
			return err
		}

		timeframeSeconds := (time.Now().UnixMilli() / 1000) - (timeframe.Weeks * weekSeconds)

		coinData, err := s.store.GetCoinData(coin, timeframeSeconds)
		if err != nil {
			return err
		}
		coinDataResponse := convertToCoinDataListResponse(coinData)

		WriteJSON(w, http.StatusOK, coinDataResponse)
		return nil
	}

	return fmt.Errorf("Method not allowed: %s", r.Method)
}

// Endpoint that returns all coins from a given source to frontend
// /{source}

func (s *APIServer) SourceHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		source := mux.Vars(r)["source"]

		sourceData, err := s.store.GetSourceCoins(source)
		if err != nil {
			return err
		}
		sourceDataResponse := convertToCoinsResponse(sourceData)

		WriteJSON(w, http.StatusOK, sourceDataResponse)

		return nil
	}

	return fmt.Errorf("Method not allowed: %s", r.Method)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandlerFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func convertToCoinsResponse(coins *db.Coins) *coinsResponse {
	return &coinsResponse{
		Coins: coins.Coins,
	}
}

func convertToCoinDataResponse(coinData *db.CoinData) *coinDataResponse {
	return &coinDataResponse{
		Time:   coinData.Time,
		Source: coinData.Source,
	}
}

func convertToCoinDataListResponse(coinDataList *db.CoinDataList) *coinDataListResponse {
	cdlr := new(coinDataListResponse)
	for _, coinData := range coinDataList.Coins {
		cdr := convertToCoinDataResponse(&coinData)
		cdlr.Coins = append(cdlr.Coins, *cdr)
	}

	return cdlr
}
