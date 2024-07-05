package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CurrencyDateValue struct {
	Date  time.Time
	Value string
}

type Currency struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	println("Starting server on port 8080")
	// Open a connection to the SQLite database.
	db, err := sql.Open("sqlite3", "./trades.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	defer db.Close()

	// Initialize the database.
	if err := InitializeDB(db); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// first request takes more than 200ms
	GetCurrencyCurrentValue(900 * time.Millisecond)

	http.HandleFunc("/", insertHandler(db))
	http.ListenAndServe(":8080", nil)
}

func insertHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/cotacao" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.Println("Received request")
		defer log.Println("Request ended")

		cash, error := GetCurrencyCurrentValue(200 * time.Millisecond)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		persistCurrencyValue(db, cash)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cash.USDBRL.Bid)

	}
}

func InitializeDB(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS trade (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"DAY" DATE DEFAULT CURRENT_TIMESTAMP, 
		"VALUE" REAL NOT NULL
	);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}
	return nil
}

func GetCurrencyCurrentValue(unit time.Duration) (*Currency, error) {
	ctx, cancel := context.WithTimeout(context.Background(), unit)

	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var apiResponse Currency
	error = json.Unmarshal(body, &apiResponse)

	return &apiResponse, error

}

func persistCurrencyValue(db *sql.DB, cash *Currency) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	insertSQL := `INSERT INTO trade (VALUE) VALUES (?)`
	statement, err := db.PrepareContext(ctx, insertSQL)
	if err != nil {
		log.Fatalf("failed to prepare statement: %v", err)
	}
	defer statement.Close()

	_, err = statement.ExecContext(ctx, cash.USDBRL.Bid)
	if err != nil {
		log.Fatalf("failed to execute statement: %v", err)
	}

	return nil

}
