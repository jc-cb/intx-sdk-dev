package intx

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

func TestGetAsset(t *testing.T) {
	credentials := &Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
		t.Fatalf("Error unmarshalling credentials: %v", err)
	}

	client := NewClient(credentials, http.Client{})

	ctx := context.Background()

	request := &GetAssetRequest{
		Asset: "BTC",
	}

	response, err := client.GetAsset(ctx, request)
	if err != nil {
		t.Errorf("GetAsset returned an error: %v", err)
	}

	if response.AssetDetail == nil {
		t.Error("Expected non-nil AssetDetail")
	}
}
