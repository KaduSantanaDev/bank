package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload contém os dados do token JWT.
type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

// NewPayload cria um novo payload com informações do token.
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	payload := &Payload{
		ID:       tokenID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	}

	return payload, nil
}

// Valid verifica se o token ainda é válido.
func (payload *Payload) Valid() error {
	// Verifica se o token está expirado
	if payload.ExpiresAt != nil && payload.ExpiresAt.Time.Before(time.Now()) {
		return ErrExpiredToken
	}

	// Outras validações podem ser adicionadas aqui, se necessário
	return nil
}
