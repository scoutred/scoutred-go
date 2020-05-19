package scoutred

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenHasExpred will check if a token is expired. This function does not
// do any other validation checks.
func TokenHasExpired(token string) (bool, error) {
	var claims jwt.StandardClaims

	_, _, err := new(jwt.Parser).ParseUnverified(token, &claims)
	if err != nil {
		return true, err
	}

	return claims.ExpiresAt < time.Now().Unix(), nil
}
