package test

import (
	"context"
	"github.com/coinbase-samples/intx-sdk-go"
	"testing"
	"time"
)

func TestGetAsset(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.GetAsset(ctx, &intx.GetAssetRequest{
		Asset: "BTC",
	})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if response.AssetDetail == nil {
		t.Fatal("expected asset detail in get")
	}

}
