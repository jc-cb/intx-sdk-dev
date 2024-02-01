package intx

import (
	"context"
	"fmt"
)

type GetHistoricalFundingRequest struct {
	Instrument string `json:"instrument"`
}

type GetHistoricalFundingResponse struct {
	HistoricalFundingRates *HistoricalFundingRate       `json:"historical_funding_rates"`
	Request                *GetHistoricalFundingRequest `json:"request"`
}

func (c Client) GetHistoricalFundingRates(ctx context.Context,
	request *GetHistoricalFundingRequest) (*GetHistoricalFundingResponse, error) {
	path := fmt.Sprintf("/instruments/%s/funding", request.Instrument)

	response := &GetHistoricalFundingResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.HistoricalFundingRates); err != nil {
		return nil, err
	}

	return response, nil
}
