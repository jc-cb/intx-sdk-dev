package intx

import (
	"context"
	"fmt"
)

type GetPortfolioBalancesRequest struct {
	Portfolio string `json:"portfolio"`
}

type GetPortfolioBalancesResponse struct {
	Balances []Balance                    `json:"balances"`
	Request  *GetPortfolioBalancesRequest `json:"request"`
}

func (c Client) GetPortfolioBalances(ctx context.Context,
	request *GetPortfolioBalancesRequest) (*GetPortfolioBalancesResponse, error) {
	path := fmt.Sprintf("/portfolios/%s/balances", request.Portfolio)

	response := &GetPortfolioBalancesResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Balances); err != nil {
		return nil, err
	}

	return response, nil
}
