package intx

import (
	"context"
	"fmt"
)

type ListTransfersRequest struct {
	Portfolios   string `json:"portfolios"'`
	TimeFrom     string `json:"time_from"`
	TimeTo       string `json:"time_to"`
	ResultLimit  int    `json:"result_limit"`
	ResultOffset int    `json:"result_offset"`
	Status       string `json:"status"`
	Type         string `json:"type"`
	Pagination   *PaginationParams
}

type ListTransfersResponse struct {
	Pagination PaginationSubset      `json:"pagination"`
	Transfers  []Transfer            `json:"results"`
	Request    *ListTransfersRequest `json:"request"`
}

func (c Client) ListTransfers(ctx context.Context, request *ListTransfersRequest) (*ListTransfersResponse, error) {
	path := "/transfers"

	queryParams := appendQueryParam("", "portfolios", request.Portfolios)

	if request.TimeFrom != "" {
		queryParams = appendQueryParam(queryParams, "time_from", request.TimeFrom)
	}
	if request.TimeTo != "" {
		queryParams = appendQueryParam(queryParams, "time_to", request.TimeTo)
	}
	if request.ResultLimit != 0 {
		queryParams = appendQueryParam(queryParams, "result_limit", fmt.Sprint(request.ResultLimit))
	}
	if request.ResultOffset != 0 {
		queryParams = appendQueryParam(queryParams, "result_offset", fmt.Sprint(request.ResultOffset))
	}
	if request.Status != "" {
		queryParams = appendQueryParam(queryParams, "status", request.Status)
	}
	if request.Type != "" {
		queryParams = appendQueryParam(queryParams, "type", request.Type)
	}

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &ListTransfersResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Transfers); err != nil {
		return nil, err
	}

	return response, nil
}
