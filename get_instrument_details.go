package intx

import (
	"context"
	"fmt"
)

type GetInstrumentRequest struct {
	Instrument string `json:"instrument"`
}

type GetInstrumentResponse struct {
	InstrumentDetail *Instrument           `json:"instrument_detail"`
	Request          *GetInstrumentRequest `json:"request"`
}

func (c Client) GetInstrument(ctx context.Context, request *GetInstrumentRequest) (*GetInstrumentResponse, error) {
	path := fmt.Sprintf("/instruments/%s", request.Instrument)

	response := &GetInstrumentResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.InstrumentDetail); err != nil {
		return nil, err
	}

	return response, nil
}
