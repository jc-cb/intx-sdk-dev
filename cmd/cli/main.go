package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coinbase-samples/intx-sdk-go"
	"net/http"
	"os"
)

func main() {
	credentials := &intx.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := intx.NewClient(credentials, http.Client{})

	ctx := context.Background()
	response, err := client.ListOpenOrders(ctx, &intx.ListOpenOrdersRequest{
		Portfolio:  "018a04e1-2006-7e1c-912e-23c2eb75f06f",
		Instrument: "BTC-PERP",
	})
	if err != nil {
		fmt.Println("Error listing portfolios:", err)
		return
	}

	jsonResponse, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling response to JSON:", err)
		return
	}
	fmt.Println(string(jsonResponse))
}
