package intx

import (
	"context"
)

type CancelOrdersRequest struct {
	Portfolio  string `json:"portfolio"`
	Instrument string `json:"instrument"`
}

type CancelOrdersResponse struct {
	Order   []Order
	Request *CancelOrdersRequest
}

func (c Client) CancelOrders(ctx context.Context, request *CancelOrdersRequest) (*CancelOrdersResponse, error) {
	path := "/orders"

	queryParams := appendQueryParam("", "portfolio", request.Portfolio)

	if request.Instrument != "" {
		queryParams = appendQueryParam(queryParams, "instrument", request.Instrument)
	}

	response := &CancelOrdersResponse{Request: request}

	if err := delete(ctx, c, path, queryParams, nil, &response.Order); err != nil {
		return nil, err
	}

	return response, nil
}
