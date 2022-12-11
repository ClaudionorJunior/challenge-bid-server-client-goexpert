package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Bid struct {
	Usdbrl struct {
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

type BidModel struct {
	gorm.Model
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	Timestamp  string
	CreateDate string
}

const URL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
const REQUEST_MAX_DURATION = 5 * time.Second
const DB_MAX_TIMEOUT = 10 * time.Millisecond

func main() {
	http.HandleFunc("/cotacao", fetchBid)
	http.ListenAndServe(":8080", nil)

}

func fetchBid(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&BidModel{})

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, REQUEST_MAX_DURATION)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var allBid Bid
	err = json.Unmarshal(body, &allBid)
	if err != nil {
		panic(err)
	}

	_, gormCancel := context.WithTimeout(context.Background(), DB_MAX_TIMEOUT)
	defer gormCancel()
	fmt.Println("aqui", allBid)
	db.Debug().Create(&BidModel{
		Code:       allBid.Usdbrl.Code,
		Codein:     allBid.Usdbrl.Codein,
		Name:       allBid.Usdbrl.Name,
		High:       allBid.Usdbrl.High,
		Low:        allBid.Usdbrl.Low,
		VarBid:     allBid.Usdbrl.VarBid,
		PctChange:  allBid.Usdbrl.PctChange,
		Bid:        allBid.Usdbrl.Bid,
		Ask:        allBid.Usdbrl.Ask,
		Timestamp:  allBid.Usdbrl.Timestamp,
		CreateDate: allBid.Usdbrl.CreateDate,
	})

	w.Write(body)
}
