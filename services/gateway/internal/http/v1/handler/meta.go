package handler

import (
	"context"

	v1GenAPI "gateway/internal/generated/v1"
)

type metaHandler struct{}

func NewMeta() *metaHandler {
	return &metaHandler{}
}

func (h *metaHandler) GetHealth(
	ctx context.Context,
	request v1GenAPI.GetHealthRequestObject,
) (v1GenAPI.GetHealthResponseObject, error) {
	return v1GenAPI.GetHealth200JSONResponse{
		Status: "OK",
	}, nil
}
