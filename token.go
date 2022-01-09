package scoutred

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
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
