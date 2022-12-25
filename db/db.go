package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=posts password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

// function that queries all coins on record - returns string slice

func (s *PostgresStore) GetAllCoins() (*Coins, error) {
	query := "SELECT DISTINCT coin FROM analyzed ORDER BY coin"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var coin string
	var allCoins Coins
	for rows.Next() {
		rows.Scan(&coin)
		allCoins.Coins = append(allCoins.Coins, coin)

	}

	return &allCoins, nil

}

// function that queries all data on record for a given coin during given period of time - returns coin slice

func (s *PostgresStore) GetCoinData(coin string, timeframe int64) (*CoinDataList, error) {
	query := "SELECT date, source FROM analyzed WHERE coin = ($1) AND date >= ($2)"
	rows, err := s.db.Query(query, coin, timeframe)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var coinData CoinData
	var coinDataList CoinDataList
	for rows.Next() {
		err := rows.Scan(&coinData.Time, &coinData.Source)
		if err != nil {
			return nil, err

		}

		coinDataList.Coins = append(coinDataList.Coins, coinData)
	}

	return &coinDataList, nil

}

func (s *PostgresStore) GetSourceCoins(source string) (*Coins, error) {
	query := "SELECT DISTINCT coin FROM analyzed WHERE source = ($1) ORDER BY coin"
	rows, err := s.db.Query(query, source)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var coin string
	var sourceCoins Coins
	for rows.Next() {
		err := rows.Scan(&coin)
		if err != nil {
			return nil, err
		}
		sourceCoins.Coins = append(sourceCoins.Coins, coin)
	}

	return &sourceCoins, nil
}
