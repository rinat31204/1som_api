package application

import (
	"1som/internal/application/service"
	"context"

	"1som/internal/auth/domain"
)

type AuthService struct {
	provider domain.Provider
	eventBus service.EventBus
}

func NewAuthService(provider domain.Provider, eventBus service.EventBus) *AuthService {
	return &AuthService{
		provider: provider,
		eventBus: eventBus,
	}
}

func (s *AuthService) VerifyUser(
	ctx context.Context,
	token string,
) (*domain.User, error) {
	//user, err := s.provider.VerifyToken(ctx, token)
	//if err != nil {
	//	return nil, err
	//}

	user := &domain.User{
		UID:   "7fcfa8a0-d950-4d85-bd29-7c9427476099",
		Email: "test@gmail.com",
	}

	event := service.UserVerifiedEvent{
		UserID: user.UID,
		Email:  user.Email,
	}

	go s.eventBus.Publish(service.EventUserVerified, event)

	return user, nil
}
