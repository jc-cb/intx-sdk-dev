package intx

import (
	"context"
)

type TransferBetweenPortfoliosRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Asset  string `json:"asset"`
	Amount string `json:"amount"`
}

type TransferBetweenPortfoliosResponse struct {
	Success bool                              `json:"success"`
	Request *TransferBetweenPortfoliosRequest `json:"request"`
}

func (c Client) TransferBetweenPortfolios(ctx context.Context, request *TransferBetweenPortfoliosRequest) (*TransferBetweenPortfoliosResponse, error) {

	path := "/portfolios/transfer"

	response := &TransferBetweenPortfoliosResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
