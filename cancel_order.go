package intx

import (
	"context"
	"fmt"
)

type CancelOrderRequest struct {
	Portfolio string `json:"portfolio"`
	Id        string `json:"id"`
}

type CancelOrderResponse struct {
	Order   *Order
	Request *CancelOrderRequest
}

func (c Client) CancelOrder(ctx context.Context, request *CancelOrderRequest) (*CancelOrderResponse, error) {
	path := fmt.Sprintf("/orders/%s", request.Id)

	queryParams := appendQueryParam("", "portfolio", request.Portfolio)

	response := &CancelOrderResponse{Request: request}

	if err := delete(ctx, c, path, queryParams, nil, &response.Order); err != nil {
		return nil, err
	}

	return response, nil
}
