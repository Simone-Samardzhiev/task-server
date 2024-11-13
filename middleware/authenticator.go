package middleware

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTAuthenticator is an implementation of Authenticator.
type JWTAuthenticator struct {
	// secret is used to hash the password.
	secret []byte
	// audience is used to set the audience of the token.
	audience []string
	// issuer is used to set the token issuer.
	issuer string
}

// checkClaims will check if the claims are valid.
func (a *JWTAuthenticator) checkClaims(claims *jwt.RegisteredClaims) bool {
	if claims == nil {
		log.Printf("The token claims are nil")
		return false
	}

	if claims.Issuer != a.issuer {
		log.Printf("The token issuer is not %s", a.issuer)
		return false
	}

	if !reflect.DeepEqual([]string(claims.Audience), a.audience) {
		log.Printf("The token audience is not %v", a.audience)
		return false
	}

	return true
}

// CreateRefreshToken will create a new token with an id.
func (a *JWTAuthenticator) CreateRefreshToken(id *uuid.UUID) (*string, error) {
	claims := jwt.RegisteredClaims{
		Audience:  a.audience,
		Issuer:    a.issuer,
		ID:        id.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.secret)
	if err != nil {
		log.Printf("Error in middleware-authenticator-CreateRefreshToken: %v", err)
		return nil, err
	}

	return &tokenString, nil
}

// CheckRefreshToken will check the access token and return the id of the token.
func (a *JWTAuthenticator) CheckRefreshToken(tokenString *string) (*uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(*tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return a.secret, nil
	})
	if err != nil {
		log.Printf("Error in middleware-authenticator-CheckRefreshToken: %v", err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		log.Printf("The token or the claims are not valid in CheckRefreshToken line 73")
		return nil, fmt.Errorf("the token claims are not valid")
	}

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		log.Printf("Error in middleware-authenticator-CheckRefreshToken: %v", err)
	}

	return &id, nil
}

// CreateAccessToken will create a new token with an id of the user.
func (a *JWTAuthenticator) CreateAccessToken(id *uuid.UUID) (*string, error) {
	claims := jwt.RegisteredClaims{
		Audience:  a.audience,
		Issuer:    a.issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(9 * time.Minute)),
		Subject:   id.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(a.secret)
	if err != nil {
		log.Printf("Error in middleware-authenticator-CreateAccessToken: %v", err)
	}
	return &tokenString, nil
}

// CheckAccessToken will check the access token and return the user id.
func (a *JWTAuthenticator) CheckAccessToken(tokenString *string) (*uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(*tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return a.secret, nil
	})
	if err != nil {
		log.Printf("Error in middleware-authenticator-CheckAccessToken: %v", err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		log.Printf("The token or the claims are not valid in middleware-authenticator-CheckAccessToken")
		return nil, fmt.Errorf("the token claims are not valid")
	}

	if !a.checkClaims(claims) {
		log.Printf("The token claims are not valid in middleware-authenticator-CheckAccessToken")
		return nil, fmt.Errorf("the token claims are not valid")
	}

	id, err := uuid.Parse(claims.Subject)
	if err != nil {
		log.Printf("Error in middleware-authenticator-CheckAccessToken: %v", err)
	}

	return &id, nil
}

// NewJWTAuthenticator will create a new authenticator.
func NewJWTAuthenticator(secret []byte, audience []string, issuer string) *JWTAuthenticator {
	return &JWTAuthenticator{secret, audience, issuer}
}
