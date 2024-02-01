package intx

import (
	"context"
	"fmt"
)

type GetPortfolioRequest struct {
	Portfolio string `json:"portfolio"`
}

type GetPortfolioResponse struct {
	Portfolio *Portfolio           `json:"portfolio"`
	Request   *GetPortfolioRequest `json:"request"`
}

func (c Client) GetPortfolio(ctx context.Context, request *GetPortfolioRequest) (*GetPortfolioResponse, error) {
	path := fmt.Sprintf("/portfolios/%s", request.Portfolio)

	response := &GetPortfolioResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Portfolio); err != nil {
		return nil, err
	}

	return response, nil
}
