package tool

import (
	"fmt"
	"github.com/FPNL/admin/src/core/entity"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var hmacSampleSecret = []byte("my_secret_key")

func GenerateToken(user *entity.AMI) (string, error) {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	// For HMAC signing method, the key can be any []byte. It is recommended to generate
	// a key using crypto/rand or something equivalent. You need the same key for signing
	// and validating.

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Nickname":  user.Nickname,
		"Organize":  user.Organize,
		"_id":       user.Id,
		"ExpiresAt": tomorrow.Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(hmacSampleSecret)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("%v", err)
	}
}
