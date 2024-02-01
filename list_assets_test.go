package intx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestListAssets(t *testing.T) {
	credentials := &Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := NewClient(credentials, http.Client{})

	ctx := context.Background()

	response, err := client.ListAssets(ctx, nil)
	if err != nil {
		t.Errorf("ListAssets returned an error: %v", err)
	}
	if len(response.Assets) == 0 {
		t.Error("Expected non-empty assets")
	}

}
