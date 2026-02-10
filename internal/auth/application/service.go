package application

import (
	"context"

	"1som/internal/auth/domain"
)

type AuthService struct {
	provider domain.Provider
}

func NewAuthService(provider domain.Provider) *AuthService {
	return &AuthService{provider: provider}
}

func (s *AuthService) VerifyAndGetUser(
	ctx context.Context,
	token string,
) (*domain.User, error) {
	return s.provider.VerifyToken(ctx, token)
}
