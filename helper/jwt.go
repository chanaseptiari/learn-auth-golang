package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var loginExpirationDuration = time.Duration(6) * time.Minute
var jwtSigningMethod = jwt.SigningMethodHS512
var jwtSignatureKey = []byte(os.Getenv("JWTKEY"))

func GenerateToken(Username string) interface{} {
	issuedAt := time.Now().Unix()
	expiresAt := time.Unix(issuedAt, 0).Add(loginExpirationDuration)
	notBefore := time.Unix(issuedAt, 0).Add(time.Nanosecond * time.Duration(-1))

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expiresAt.Unix(), 0)),
			IssuedAt:  jwt.NewNumericDate(time.Unix(issuedAt, 0)),
			NotBefore: jwt.NewNumericDate(time.Unix(notBefore.Unix(), 0)),
			// Issuer:    "test",
			// Subject:   "Authenticate",
			// ID:        "1",
			Audience: []string{Username},
		},
	)
	res, err := token.SignedString(jwtSignatureKey)
	if err != err {
		return err
	}

	return map[string]string{"X-Api-Key": string(res)}
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validasi algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Kembalikan kunci rahasia
		return []byte(jwtSignatureKey), nil
	})
	return token, err

}

// Fungsi untuk verifikasi token
func VerifyToken(tokenString string) (bool, error) {
	// Parse token
	token, err := ParseToken(tokenString)

	if err != nil {
		return false, err
	}

	// Verifikasi token
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	} else {
		return false, fmt.Errorf("Invalid token")
	}
}
