package intx

import (
	"context"
	"fmt"
)

type GetPortfolioFillsRequest struct {
	Portfolio     string `json:"portfolio"` // required
	OrderId       string `json:"order_id,omitempty"`
	ClientOrderId string `json:"client_order_id,omitempty"`
	RefDatetime   string `json:"ref_datetime,omitempty"`
	ResultLimit   int    `json:"result_limit,omitempty"`
	ResultOffset  int    `json:"result_offset,omitempty"`
	TimeFrom      string `json:"time_from,omitempty"`
	Pagination    *PaginationParams
}

type GetPortfolioFillsResponse struct {
	Pagination PaginationParams          `json:"pagination"`
	Results    []Fill                    `json:"results"`
	Request    *GetPortfolioFillsRequest `json:"request"`
}

func (c Client) GetPortfolioFills(ctx context.Context, request *GetPortfolioFillsRequest) (*GetPortfolioFillsResponse, error) {
	path := fmt.Sprintf("/portfolios/%s/fills", request.Portfolio)

	queryParams := appendQueryParam("", "portfolios", request.Portfolio)

	if request.OrderId != "" {
		queryParams = appendQueryParam(queryParams, "order_id", request.OrderId)
	}
	if request.ClientOrderId != "" {
		queryParams = appendQueryParam(queryParams, "client_order_id", request.ClientOrderId)
	}
	if request.RefDatetime != "" {
		queryParams = appendQueryParam(queryParams, "ref_datetime", request.RefDatetime)
	}
	if request.ResultLimit != 0 {
		queryParams = appendQueryParam(queryParams, "result_limit", fmt.Sprint(request.ResultLimit))
	}
	if request.ResultOffset != 0 {
		queryParams = appendQueryParam(queryParams, "result_offset", fmt.Sprint(request.ResultOffset))
	}
	if request.TimeFrom != "" {
		queryParams = appendQueryParam(queryParams, "time_from", fmt.Sprint(request.TimeFrom))
	}

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &GetPortfolioFillsResponse{Request: request}

	if err := get(ctx, c, path, queryParams, nil, response); err != nil {
		return nil, err
	}

	return response, nil
}
