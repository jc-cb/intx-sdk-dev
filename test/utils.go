package test

import (
	"encoding/json"
	"fmt"
	"github.com/coinbase-samples/intx-sdk-go"
	"net/http"
	"os"
)

func newLiveTestClient() (*intx.Client, error) {

	credentials, err := loadCredentialsFromEnv()
	if err != nil {
		return nil, err
	}

	client := intx.NewClient(credentials, http.Client{})
	return client, nil

}

func loadCredentialsFromEnv() (*intx.Credentials, error) {

	credentials := &intx.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("INTX_CREDENTIALS")), credentials); err != nil {
		return nil, fmt.Errorf("unable to deserialize intx credentials JSON: %w", err)
	}

	return credentials, nil
}
