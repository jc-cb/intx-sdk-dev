package intx

import (
	"context"
)

type CreateWithdrawalToCounterpartyIdRequest struct {
	Portfolio      string `json:"portfolio"`
	CounterpartyId string `json:"counterparty_id"`
	Asset          string `json:"asset"`
	Amount         string `json:"amount"`
	Nonce          string `json:"nonce"`
}

type CreateWithdrawalToCounterpartyIdResponse struct {
	Withdrawal *CounterpartyWithdrawal                  `json:"withdrawal"`
	Request    *CreateWithdrawalToCounterpartyIdRequest `json:"request"`
}

func (c Client) CreateWithdrawalToCounterpartyId(ctx context.Context,
	request *CreateWithdrawalToCounterpartyIdRequest) (*CreateWithdrawalToCounterpartyIdResponse, error) {

	path := "/transfers/withdraw/counterparty"

	response := &CreateWithdrawalToCounterpartyIdResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
