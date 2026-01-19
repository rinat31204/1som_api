package firebase

import (
	"1som/internal/application/auth"
	"context"
	"errors"
	"strings"

	fbAuth "firebase.google.com/go/v4/auth"
)

type AuthProvider struct {
	client *fbAuth.Client
}

func NewAuthProvider(client *fbAuth.Client) *AuthProvider {
	return &AuthProvider{client: client}
}

func (p *AuthProvider) VerifyToken(
	ctx context.Context,
	token string,
) (*auth.User, error) {

	if token == "" {
		return nil, errors.New("empty token")
	}

	token = strings.TrimSpace(token)
	token = strings.TrimPrefix(token, "Bearer ")

	idToken, err := p.client.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return &auth.User{
		UID:   idToken.UID,
		Email: idToken.Claims["email"].(string),
	}, nil
}
