package intx

import (
	"context"
)

type ValidateCounterpartyIdRequest struct {
	CounterpartyId string `json:"counterparty_id"`
}

type ValidateCounterpartyIdResponse struct {
	Validation *Validation                    `json:"validation"`
	Request    *ValidateCounterpartyIdRequest `json:"request"`
}

func (c Client) ValidateCounterpartyId(ctx context.Context,
	request *ValidateCounterpartyIdRequest) (*ValidateCounterpartyIdResponse, error) {

	path := "/transfers/validate-counterparty-id"

	response := &ValidateCounterpartyIdResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
