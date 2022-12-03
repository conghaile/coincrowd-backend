package db

import "database/sql"

type PostgresStore struct {
	db *sql.DB
}

type Coins struct {
	Coins []string
}

type CoinData struct {
	Time   int
	Source string
}

type CoinDataList struct {
	Coins []CoinData
}

type Storage interface {
	GetAllCoins() (*Coins, error)
	GetCoinData(coin string, timeframe int64) (*CoinDataList, error)
	GetSourceCoins(source string) (*Coins, error)
}
