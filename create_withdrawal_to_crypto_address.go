package intx

import (
	"context"
)

type CreateWithdrawalToCryptoAddressRequest struct {
	Portfolio            string `json:"portfolio"`
	Asset                string `json:"asset"`
	Amount               string `json:"amount"`
	AddNetworkFeeToTotal bool   `json:"add_network_fee_to_total"`
	NetworkArnId         string `json:"network_arn_id"`
	Address              string `json:"address"`
	Nonce                string `json:"nonce"`
}

type CreateWithdrawalToCryptoAddressResponse struct {
	Idem    string                                  `json:"idem"`
	Request *CreateWithdrawalToCryptoAddressRequest `json:"request"`
}

func (c Client) CreateWithdrawalToCryptoAddress(ctx context.Context,
	request *CreateWithdrawalToCryptoAddressRequest) (*CreateWithdrawalToCryptoAddressResponse, error) {

	path := "/transfers/withdraw"

	response := &CreateWithdrawalToCryptoAddressResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
