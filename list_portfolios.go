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

	var portfolios []*Portfolio

	if err := get(ctx, c, path, emptyQueryParams, request, &portfolios); err != nil {
		return nil, err
	}

	response := &ListPortfoliosResponse{
		Portfolios: portfolios,
		Request:    request,
	}

	return response, nil
}
