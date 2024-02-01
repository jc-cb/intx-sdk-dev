package intx

import (
	"context"
	"fmt"
)

type GetAssetBalanceRequest struct {
	Portfolio string `json:"portfolio"`
	Asset     string `json:"asset"`
}

type GetAssetBalanceResponse struct {
	Balance *Balance                `json:"balance"`
	Request *GetAssetBalanceRequest `json:"request"`
}

func (c Client) GetAssetBalance(ctx context.Context,
	request *GetAssetBalanceRequest) (*GetAssetBalanceResponse, error) {
	path := fmt.Sprintf("/portfolios/%s/balances/%s", request.Portfolio, request.Asset)

	response := &GetAssetBalanceResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Balance); err != nil {
		return nil, err
	}

	return response, nil
}
