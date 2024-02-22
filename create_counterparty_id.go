package intx

import (
	"context"
)

type CreateCounterpartyIdRequest struct {
	Portfolio string `json:"portfolio"`
}

type CreateCounterpartyIdResponse struct {
	Counterparty *Counterparty                `json:"counterparty"`
	Request      *CreateCounterpartyIdRequest `json:"request"`
}

func (c Client) CreateCounterpartyId(ctx context.Context,
	request *CreateCounterpartyIdRequest) (*CreateCounterpartyIdResponse, error) {

	path := "/transfers/create-counterparty-id"

	queryParams := appendQueryParam("", "portfolio", request.Portfolio)

	response := &CreateCounterpartyIdResponse{Request: request}

	if err := post(ctx, c, path, queryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
