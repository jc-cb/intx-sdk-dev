package intx

import (
	"context"
	"fmt"
)

type ModifyOrderRequest struct {
	Id            string `json:"id"`
	Portfolio     string `json:"portfolio"`
	ClientOrderId string `json:"client_order_id"`
	Price         string `json:"price,omitempty"`
	StopPrice     string `json:"stop_price,omitempty"`
	Size          string `json:"size,omitempty"`
}

type ModifyOrderResponse struct {
	Order   *Order
	Request *ModifyOrderRequest
}

func (c Client) ModifyOrder(ctx context.Context, request *ModifyOrderRequest) (*ModifyOrderResponse, error) {
	path := fmt.Sprintf("/orders/%s", request.Id)

	type modifyOrderBody struct {
		Portfolio     string `json:"portfolio"`
		ClientOrderId string `json:"client_order_id"`
		Price         string `json:"price,omitempty"`
		StopPrice     string `json:"stop_price,omitempty"`
		Size          string `json:"size,omitempty"`
	}

	body := modifyOrderBody{
		Portfolio:     request.Portfolio,
		ClientOrderId: request.ClientOrderId,
		Price:         request.Price,
		StopPrice:     request.StopPrice,
		Size:          request.Size,
	}

	response := &ModifyOrderResponse{Request: request}

	if err := put(ctx, c, path, emptyQueryParams, body, &response.Order); err != nil {
		return nil, err
	}

	return response, nil
}
