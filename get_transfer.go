package intx

import (
	"context"
	"fmt"
)

type GetTransferRequest struct {
	TransferUuid string `json:"transfer_uuid"'`
}

type GetTransferResponse struct {
	Transfers []Transfer          `json:"results"`
	Request   *GetTransferRequest `json:"request"`
}

func (c Client) GetTransfer(ctx context.Context, request *GetTransferRequest) (*GetTransferResponse, error) {
	path := fmt.Sprintf("/transfers/%s", request.TransferUuid)

	response := &GetTransferResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Transfers); err != nil {
		return nil, err
	}

	return response, nil
}
