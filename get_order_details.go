package intx

import (
	"context"
	"fmt"
)

type GetOrderDetailsRequest struct {
	Portfolio string `json:"portfolio"`
	Id        string `json:"id"`
}

type GetOrderDetailsResponse struct {
	Order   *Order
	Request *GetOrderDetailsRequest
}

func (c Client) GetOrderDetails(ctx context.Context, request *GetOrderDetailsRequest) (*GetOrderDetailsResponse, error) {
	path := fmt.Sprintf("/orders/%s", request.Id)

	queryParams := appendQueryParam("", "portfolio", request.Portfolio)

	response := &GetOrderDetailsResponse{Request: request}

	if err := get(ctx, c, path, queryParams, nil, &response.Order); err != nil {
		return nil, err
	}

	return response, nil
}
