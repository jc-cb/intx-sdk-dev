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

	var instruments []*Instrument

	if err := get(ctx, c, path, emptyQueryParams, nil, &instruments); err != nil {
		return nil, err
	}

	response := &ListInstrumentsResponse{
		Instruments: instruments,
		Request:     request,
	}

	return response, nil
}
