package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coinbase-samples/intx-sdk-go"
	"net/http"
	"os"
)

//func main() {
//	credentials := &intx.Credentials{}
//	if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	client := intx.NewClient(credentials, http.Client{})
//
//	ctx := context.Background()
//
//	//response, err := client.GetAssetBalance(ctx, &intx.GetAssetBalanceRequest{
//	//	Portfolio: "018a04e1-2006-7e1c-912e-23c2eb75f06f",
//	//	Asset:     "BTC",
//	//})
//
//	response, err := client.ListOpenOrders(ctx, &intx.ListOpenOrdersRequest{
//		Portfolio: "018a04e1-2006-7e1c-912e-23c2eb75f06f",
//	})
//
//	if err != nil {
//		fmt.Println("Error listing portfolios:", err)
//		return
//	}
//
//	jsonResponse, err := json.MarshalIndent(response, "", "  ")
//	if err != nil {
//		fmt.Println("Error marshaling response to JSON:", err)
//		return
//	}
//	fmt.Println(string(jsonResponse))
//}

//func main() {
//	credentials := &intx.Credentials{}
//	if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	client := intx.NewClient(credentials, http.Client{})
//
//	ctx := context.Background()
//
//	clientOrderId, err := uuid.NewUUID()
//	if err != nil {
//		log.Fatal("Error generating UUID:", err)
//	}
//
//	response, err := client.CreateOrder(ctx, &intx.CreateOrderRequest{
//		ClientOrderId: clientOrderId.String(),
//		Side:          "BUY",
//		Instrument:    "b3469e0b-222c-4f8a-9f68-1f9e44d7e5e0",
//		Type:          "LIMIT",
//		Tif:           "GTC",
//		Size:          "0.1",
//		Price:         "43000",
//		Portfolio:     "018a04e1-2006-7e1c-912e-23c2eb75f06f",
//	})
//
//	if err != nil {
//		fmt.Println("Error listing portfolios:", err)
//		return
//	}
//
//	jsonResponse, err := json.MarshalIndent(response, "", "  ")
//	if err != nil {
//		fmt.Println("Error marshaling response to JSON:", err)
//		return
//	}
//	fmt.Println(string(jsonResponse))
//}

//
//func main() {
//	credentials := &intx.Credentials{}
//	if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	client := intx.NewClient(credentials, http.Client{})
//
//	ctx := context.Background()
//
//	response, err := client.SetMarginOverride(ctx, &intx.SetMarginOverrideRequest{
//		Portfolio:      "018a04e1-2006-7e1c-912e-23c2eb75f06f",
//		MarginOverride: "0.1",
//	})
//
//	if err != nil {
//		fmt.Println("Error listing portfolios:", err)
//		return
//	}
//
//	jsonResponse, err := json.MarshalIndent(response, "", "  ")
//	if err != nil {
//		fmt.Println("Error marshaling response to JSON:", err)
//		return
//	}
//	fmt.Println(string(jsonResponse))
//}

func main() {
	credentials := &intx.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := intx.NewClient(credentials, http.Client{})

	ctx := context.Background()
	response, err := client.ListPortfolios(ctx, &intx.ListPortfoliosRequest{})
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
