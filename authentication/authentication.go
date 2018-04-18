package authentication

import (
	"errors"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Owner is an authorized user
type Owner struct {
	ID           int
	EmailAddress string
	Role         string
}

type owner struct {
	Owner
	jwt.StandardClaims
}

// Authenticate login
func Authenticate(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

// Authorize is used to request token
func Authorize(o Owner, issuer string) (string, error) {
	signingKey := []byte(os.Getenv("APP_KEY"))
	claims := owner{
		o,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

// Token authentication
func Token(tokenString string, roles []string) (Owner, error) {
	match, _ := regexp.Match("^Bearer ", []byte(tokenString))
	if !match {
		return Owner{}, errors.New("Invalid token")
	}
	tokenString = tokenString[7:]

	signingKey := []byte(os.Getenv("APP_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &owner{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if len(roles) != 0 && !contains(roles, token.Claims.(*owner).Role) {
		return Owner{}, errors.New("Forbidden")
	}

	if claims, ok := token.Claims.(*owner); ok && token.Valid {
		return Owner{
			ID:           claims.ID,
			EmailAddress: claims.EmailAddress,
			Role:         claims.Role,
		}, nil
	}

	return Owner{}, err
}

func contains(s []string, e string) bool {
	for _, c := range s {
		if c == e {
			return true
		}
	}
	return false
}
