package intx

import (
	"context"
)

type ListPortfoliosRequest struct{}

type ListPortfoliosResponse struct {
	Portfolios []*Portfolio           `json:"portfolios"`
	Request    *ListPortfoliosRequest `json:"request"`
}

func (c Client) ListPortfolios(ctx context.Context, request *ListPortfoliosRequest) (*ListPortfoliosResponse, error) {
	path := "/portfolios"

	response := &ListPortfoliosResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, request, &response.Portfolios); err != nil {
		return nil, err
	}

	return response, nil
}
