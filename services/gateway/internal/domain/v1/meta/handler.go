package meta

import (
	"context"

	v1GenAPI "gateway/internal/generated/v1"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) GetHealth(
	ctx context.Context,
	request v1GenAPI.GetHealthRequestObject,
) (v1GenAPI.GetHealthResponseObject, error) {
	return v1GenAPI.GetHealth200JSONResponse{
		Status: "OK",
	}, nil
}
