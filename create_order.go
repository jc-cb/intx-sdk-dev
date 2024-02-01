package intx

import (
	"context"
	"errors"
)

type CreateOrderResponse struct {
	Order   *Order              `json:"order"`
	Request *CreateOrderRequest `json:"request"`
}

func (c Client) CreateOrder(ctx context.Context, request *CreateOrderRequest) (*CreateOrderResponse, error) {
	if request == nil {
		return nil, errors.New("create order request is nil")
	}

	path := "/orders"

	response := &CreateOrderResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response.Order); err != nil {
		return nil, err
	}

	return response, nil
}
