package auth

import (
	"context"

	authv1 "starter.local/gen/auth/v1"
)

type Handler struct {
	authv1.UnimplementedAuthServiceServer
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Register(
	ctx context.Context,
	req *authv1.RegisterRequest,
) (*authv1.RegisterResponse, error) {
	return &authv1.RegisterResponse{
		UserId: "test-user-id",
	}, nil
}

func (h *Handler) Login(
	ctx context.Context,
	req *authv1.LoginRequest,
) (*authv1.LoginResponse, error) {
	return &authv1.LoginResponse{
		AccessToken:  "access",
		RefreshToken: "refresh",
	}, nil
}
