package scoutred_test

import (
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/scoutred/scoutred-go"
)

func TestTokenHasExpired(t *testing.T) {
	type tcase struct {
		token      string
		hasExpired bool
	}

	fn := func(tc tcase) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()

			expired, err := scoutred.TokenHasExpired(tc.token)
			if err != nil {
				t.Errorf("unexepected err: %v", err)
				return
			}

			if expired != tc.hasExpired {
				t.Errorf("expected %v got %v", tc.hasExpired, expired)
				return
			}
		}
	}

	// genTokenExpires is a helper to generate tokens and set the ExpiresAt value
	// using a provided daysFromNow offset
	genTokenExpires := func(daysFromNow int) string {
		tokenPast := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, daysFromNow).Unix(), // subtract one day
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := tokenPast.SignedString([]byte("scoutred-secret"))
		if err != nil {
			t.Fatalf("unexepected err: %v", err)
		}

		return tokenString
	}

	// Auth handler will make sure that the route is never allowed, so
	// no need to test for anonymous user.
	tests := map[string]tcase{
		"expired": {
			token:      genTokenExpires(-1),
			hasExpired: true,
		},
		"not expired": {
			token:      genTokenExpires(1),
			hasExpired: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, fn(tc))
	}
}
