package firebase

import (
	"1som/internal/auth/domain"
	"context"
	"errors"
	"strings"

	fbAuth "firebase.google.com/go/v4/auth"
)

type Provider struct {
	client *fbAuth.Client
}

func NewFirebaseProvider(client *fbAuth.Client) domain.Provider {
	return &Provider{client: client}
}

func (p *Provider) VerifyToken(
	ctx context.Context,
	token string,
) (*domain.User, error) {

	if token == "" {
		return nil, errors.New("empty token")
	}

	token = strings.TrimSpace(token)
	token = strings.TrimPrefix(token, "Bearer ")

	authToken, err := p.client.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		UID:   authToken.UID,
		Email: authToken.Claims["email"].(string),
	}, nil
}
