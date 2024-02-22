package intx

import (
	"context"
)

type CreatePortfolioTransferRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Asset  string `json:"asset"`
	Amount string `json:"amount"`
}

type CreatePortfolioTransferResponse struct {
	Success bool                            `json:"success"`
	Request *CreatePortfolioTransferRequest `json:"request"`
}

func (c Client) CreatePortfolioTransfer(ctx context.Context, request *CreatePortfolioTransferRequest) (*CreatePortfolioTransferResponse, error) {

	path := "/portfolios/transfer"

	response := &CreatePortfolioTransferResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
