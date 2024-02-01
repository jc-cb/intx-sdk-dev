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

	response := &ListAssetsResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Assets); err != nil {
		return nil, err
	}

	return response, nil
}
