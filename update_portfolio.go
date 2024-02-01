package intx

import (
	"context"
	"errors"
	"fmt"
)

type UpdatePortfolioRequest struct {
	Name      string `json:"name"`
	Portfolio string `json:"portfolio"`
}

type UpdatePortfolioResponse struct {
	Portfolio *Portfolio              `json:"portfolio"`
	Request   *UpdatePortfolioRequest `json:"request"`
}

func (c Client) UpdatePortfolio(ctx context.Context, request *UpdatePortfolioRequest) (*UpdatePortfolioResponse, error) {
	if request == nil {
		return nil, errors.New("Update Portfolio request is nil")
	}

	path := fmt.Sprintf("/portfolios/%s", request.Portfolio)

	response := &UpdatePortfolioResponse{Request: request}

	if err := put(ctx, c, path, emptyQueryParams, request, &response.Portfolio); err != nil {
		return nil, err
	}

	return response, nil
}
