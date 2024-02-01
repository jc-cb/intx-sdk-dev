package intx

import (
	"context"
	"fmt"
)

type GetPortfolioPositionsRequest struct {
	Portfolio string `json:"portfolio"`
}

type GetPortfolioPositionsResponse struct {
	Positions []Position                    `json:"positions"`
	Request   *GetPortfolioPositionsRequest `json:"request"`
}

func (c Client) GetPortfolioPositions(ctx context.Context,
	request *GetPortfolioPositionsRequest) (*GetPortfolioPositionsResponse, error) {
	path := fmt.Sprintf("/portfolios/%s/positions", request.Portfolio)

	response := &GetPortfolioPositionsResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Positions); err != nil {
		return nil, err
	}

	return response, nil
}
