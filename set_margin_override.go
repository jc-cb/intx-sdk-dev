package intx

import (
	"context"
	"errors"
)

type SetMarginOverrideRequest struct {
	Portfolio      string `json:"portfolio"`
	MarginOverride string `json:"margin_override"`
}

type SetMarginOverrideResponse struct {
	MarginOverride *MarginOverride           `json:"margin_override"`
	Request        *SetMarginOverrideRequest `json:"request"`
}

func (c Client) SetMarginOverride(ctx context.Context,
	request *SetMarginOverrideRequest) (*SetMarginOverrideResponse, error) {
	if request == nil {
		return nil, errors.New("set margin override is nil")
	}

	path := "/portfolios/margin"

	response := &SetMarginOverrideResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
