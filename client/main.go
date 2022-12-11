package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	bidStructs "github.com/ClaudionorJunior/challenge-bid-server-client-goexpert/bid-structs"
)

const MAX_DURATION = time.Duration(time.Millisecond * 300)
const URL = "http://localhost:8080/cotacao"
const FILE_NAME = "cotacao.txt"

func main() {
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), MAX_DURATION)
	defer cancel()

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var parsedData bidStructs.Bid
	err = json.Unmarshal(responseData, &parsedData)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(FILE_NAME)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error to create"+FILE_NAME+": %v\n", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", parsedData.Usdbrl.Bid))
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error to write"+FILE_NAME+": %v\n", err)
	}
}
