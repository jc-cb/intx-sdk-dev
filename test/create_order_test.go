package test

import (
	"context"
	"github.com/coinbase-samples/intx-sdk-go"
	"github.com/google/uuid"
	"log"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	clientOrderId, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("Error generating UUID:", err)
	}

	response, err := client.CreateOrder(
		ctx,
		&intx.CreateOrderRequest{
			Portfolio:     client.Credentials.PortfolioId,
			ClientOrderId: clientOrderId.String(),
			Side:          "BUY",
			Instrument:    "b3469e0b-222c-4f8a-9f68-1f9e44d7e5e0",
			Type:          "LIMIT",
			Tif:           "GTC",
			Size:          "0.1",
			Price:         "43000",
		})
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("expected a not nil response")
	}

	if len(response.Order.OrderId) == 0 {
		t.Fatal("expected an activity id")
	}

}
