package intx

import (
	"context"
)

type CreatePortfolioRequest struct {
	Name string `json:"name"`
}

type CreatePortfolioResponse struct {
	Portfolio *Portfolio              `json:"portfolio"`
	Request   *CreatePortfolioRequest `json:"request"`
}

func (c Client) CreatePortfolio(ctx context.Context,
	request *CreatePortfolioRequest) (*CreatePortfolioResponse, error) {

	path := "/portfolios"

	response := &CreatePortfolioResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response.Portfolio); err != nil {
		return nil, err
	}

	return response, nil
}
