package intx

import (
	"context"
	"fmt"
)

type GetAssetRequest struct {
	Asset string `json:"asset"`
}

type GetAssetResponse struct {
	AssetDetail *Asset
	Request     *GetAssetRequest
}

func (c Client) GetAsset(ctx context.Context, request *GetAssetRequest) (*GetAssetResponse, error) {
	path := fmt.Sprintf("/assets/%s", request.Asset)

	response := &GetAssetResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.AssetDetail); err != nil {
		return nil, err
	}

	return response, nil
}
