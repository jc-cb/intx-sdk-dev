package intx

import (
	"context"
)

type CreateCryptoAddressRequest struct {
	Portfolio    string `json:"portfolio"`
	Asset        string `json:"asset"`
	NetworkArnId string `json:"network_arn_id"`
}

type CreateCryptoAddressResponse struct {
	Address *Address                    `json:"address"`
	Request *CreateCryptoAddressRequest `json:"request"`
}

func (c Client) CreateCryptoAddress(ctx context.Context,
	request *CreateCryptoAddressRequest) (*CreateCryptoAddressResponse, error) {

	path := "/transfers/address"

	response := &CreateCryptoAddressResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
