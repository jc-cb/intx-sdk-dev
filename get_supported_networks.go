package intx

import (
	"context"
	"fmt"
)

type GetSupportedNetworksRequest struct {
	Asset string `json:"asset"`
}

type GetSupportedNetworksResponse struct {
	NetworkDetail []Network                    `json:"network_detail"`
	Request       *GetSupportedNetworksRequest `json:"request"`
}

func (c Client) GetSupportedNetworks(ctx context.Context,
	request *GetSupportedNetworksRequest) (*GetSupportedNetworksResponse, error) {

	path := fmt.Sprintf("/assets/%s/networks", request.Asset)

	response := &GetSupportedNetworksResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.NetworkDetail); err != nil {
		return nil, err
	}

	return response, nil
}
