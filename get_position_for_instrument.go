package intx

import (
	"context"
	"fmt"
)

type GetInstrumentPositionRequest struct {
	Portfolio  string `json:"portfolio"`
	Instrument string `json:"instrument"`
}

type GetInstrumentPositionResponse struct {
	Positions []Position                    `json:"positions"`
	Request   *GetInstrumentPositionRequest `json:"request"`
}

func (c Client) GetInstrumentPosition(ctx context.Context,
	request *GetInstrumentPositionRequest) (*GetInstrumentPositionResponse, error) {
	path := fmt.Sprintf("/portfolios/%s/positions/%s", request.Portfolio, request.Instrument)

	response := &GetInstrumentPositionResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, response); err != nil {
		return nil, err
	}

	return response, nil
}
