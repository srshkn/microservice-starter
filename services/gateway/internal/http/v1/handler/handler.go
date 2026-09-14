package handler

import (
	v1GenAPI "gateway/internal/generated/v1"
)

type handler struct {
	metaHandler
}

func New(meta *metaHandler) *handler {
	return &handler{
		metaHandler: *meta,
	}
}

var _ v1GenAPI.StrictServerInterface = (*handler)(nil)
