package intx

import (
	"context"
	"fmt"
)

type GetInstrumentQuoteRequest struct {
	Instrument string `json:"instrument"`
}

type GetInstrumentQuoteResponse struct {
	InstrumentQuote *Quote                     `json:"instrument_quote"`
	Request         *GetInstrumentQuoteRequest `json:"request"`
}

func (c Client) GetInstrumentQuote(ctx context.Context,
	request *GetInstrumentQuoteRequest) (*GetInstrumentQuoteResponse, error) {
	path := fmt.Sprintf("/instruments/%s/quote", request.Instrument)

	response := &GetInstrumentQuoteResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.InstrumentQuote); err != nil {
		return nil, err
	}

	return response, nil
}
