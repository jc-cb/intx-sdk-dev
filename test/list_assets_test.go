package test

import (
	"context"
	"github.com/coinbase-samples/intx-sdk-go"
	"testing"
	"time"
)

func TestListAssets(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newLiveTestClient()
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.ListAssets(ctx, &intx.ListAssetsRequest{})

	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal(err)
	}

	if len(response.Assets) == 0 {
		t.Fatal("expected assets in get")
	}
}
