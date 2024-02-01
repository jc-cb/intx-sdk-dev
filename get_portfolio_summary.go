package intx

import (
	"context"
	"fmt"
)

type GetPortfolioSummaryRequest struct {
	Portfolio string `json:"portfolio"`
}

type GetPortfolioSummaryResponse struct {
	Summary *Summary                    `json:"summary"`
	Request *GetPortfolioSummaryRequest `json:"request"`
}

func (c Client) GetPortfolioSummary(ctx context.Context,
	request *GetPortfolioSummaryRequest) (*GetPortfolioSummaryResponse, error) {
	path := fmt.Sprintf("/portfolios/%s/summary", request.Portfolio)

	response := &GetPortfolioSummaryResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Summary); err != nil {
		return nil, err
	}

	return response, nil
}
