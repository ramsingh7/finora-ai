package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims carried in access tokens (HS256).
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// JWT signs and parses access tokens.
type JWT struct {
	secret []byte
	issuer string
	ttl    time.Duration
}

// NewJWT builds a signer/parser. Secret must be non-empty.
func NewJWT(secret, issuer string, ttl time.Duration) (*JWT, error) {
	if len(secret) < 8 {
		return nil, fmt.Errorf("jwt secret must be at least 8 bytes")
	}
	if issuer == "" {
		issuer = "finora-backend"
	}
	if ttl < time.Minute || ttl > 24*time.Hour*7 {
		return nil, fmt.Errorf("jwt ttl must be between 1m and 7d")
	}
	return &JWT{secret: []byte(secret), issuer: issuer, ttl: ttl}, nil
}

// Sign issues an access token for the subject user id and email.
func (j *JWT) Sign(userID, email string) (token string, expiresAt time.Time, err error) {
	now := time.Now()
	expiresAt = now.Add(j.ttl)
	claims := Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString(j.secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return s, expiresAt, nil
}

// Parse validates a bearer token and returns claims.
func (j *JWT) Parse(token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}
	c, ok := t.Claims.(*Claims)
	if !ok || !t.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return c, nil
}
