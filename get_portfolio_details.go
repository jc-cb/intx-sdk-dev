package intx

import (
	"context"
	"fmt"
)

type GetPortfolioDetailsRequest struct {
	Portfolio string `json:"portfolio"`
}

type GetPortfolioDetailsResponse struct {
	Details *Details                    `json:"details"`
	Request *GetPortfolioDetailsRequest `json:"request"`
}

func (c Client) GetPortfolioDetails(ctx context.Context,
	request *GetPortfolioDetailsRequest) (*GetPortfolioDetailsResponse, error) {
	path := fmt.Sprintf("/portfolios/%s/detail", request.Portfolio)

	response := &GetPortfolioDetailsResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Details); err != nil {
		return nil, err
	}

	return response, nil
}
