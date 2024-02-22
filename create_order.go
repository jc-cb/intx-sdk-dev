package intx

import (
	"context"
)

type CreateOrderRequest struct {
	ClientOrderId string  `json:"client_order_id"`
	Side          string  `json:"side"`
	Size          string  `json:"size"`
	Tif           string  `json:"tif"`
	Instrument    string  `json:"instrument"`
	Type          string  `json:"type"`
	Price         string  `json:"price,omitempty"`
	StopPrice     *string `json:"stop_price,omitempty"`
	ExpireTime    *string `json:"expire_time,omitempty"`
	Portfolio     string  `json:"portfolio"`
	User          *string `json:"user,omitempty"`
	StpMode       *string `json:"stp_mode,omitempty"`
	PostOnly      *bool   `json:"post_only,omitempty"`
}

type CreateOrderResponse struct {
	Order   *Order              `json:"order"`
	Request *CreateOrderRequest `json:"request"`
}

func (c Client) CreateOrder(ctx context.Context,
	request *CreateOrderRequest) (*CreateOrderResponse, error) {

	path := "/orders"

	response := &CreateOrderResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response.Order); err != nil {
		return nil, err
	}

	return response, nil
}
