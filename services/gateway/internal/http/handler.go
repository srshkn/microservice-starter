package handler

import (
	"context"

	v1GenAPI "gateway/internal/generated/v1"
)

type metaHandler interface {
	GetHealth(ctx context.Context, request v1GenAPI.GetHealthRequestObject) (v1GenAPI.GetHealthResponseObject, error)
}

type handler struct {
	metaHandler
}

func New(meta metaHandler) *handler {
	return &handler{
		metaHandler: meta,
	}
}

var _ v1GenAPI.StrictServerInterface = (*handler)(nil)
