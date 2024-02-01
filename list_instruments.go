package intx

import (
	"context"
)

type ListInstrumentsRequest struct{}

type ListInstrumentsResponse struct {
	Instruments []*Instrument           `json:"instruments"`
	Request     *ListInstrumentsRequest `json:"request"`
}

func (c Client) ListInstruments(ctx context.Context, request *ListInstrumentsRequest) (*ListInstrumentsResponse, error) {
	path := "/instruments"

	response := &ListInstrumentsResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Instruments); err != nil {
		return nil, err
	}

	return response, nil
}
