package intx

import (
	"context"
)

type ListAssetsRequest struct{}

type ListAssetsResponse struct {
	Assets  []*Asset           `json:"assets"`
	Request *ListAssetsRequest `json:"request"`
}

func (c Client) ListAssets(ctx context.Context, request *ListAssetsRequest) (*ListAssetsResponse, error) {
	path := "/assets"

	var assets []*Asset

	if err := get(ctx, c, path, emptyQueryParams, nil, &assets); err != nil {
		return nil, err
	}

	response := &ListAssetsResponse{
		Assets:  assets,
		Request: request,
	}

	return response, nil
}
