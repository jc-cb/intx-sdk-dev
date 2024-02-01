package intx

import (
	"context"
	"fmt"
)

type GetSupportedNetworksRequest struct {
	Asset string `json:"asset"`
}

type GetSupportedNetworksResponse struct {
	NetworkDetail []Network
	Request       *GetSupportedNetworksRequest
}

func (c Client) GetSupportedNetworks(ctx context.Context,
	request *GetSupportedNetworksRequest) (*GetSupportedNetworksResponse, error) {

	path := fmt.Sprintf("/assets/%s/networks", request.Asset)

	var networkDetail []Network

	if err := get(ctx, c, path, emptyQueryParams, nil, &networkDetail); err != nil {
		return nil, err
	}

	response := &GetSupportedNetworksResponse{
		NetworkDetail: networkDetail,
		Request:       request,
	}

	return response, nil
}
