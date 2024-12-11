package token

import (
	"time"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)

	verifyToken(token string) (*Payload, error)
}
