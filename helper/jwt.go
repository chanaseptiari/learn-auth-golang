package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var loginExpirationDuration = time.Duration(1) * time.Hour
var jwtSigningMethod = jwt.SigningMethodHS256
var jwtSignatureKey = []byte(os.Getenv("JWTKEY"))

type customClaims struct {
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(Username string) interface{} {
	issuedAt := time.Now().Unix()
	expiresAt := time.Unix(issuedAt, 0).Add(loginExpirationDuration)
	notBefore := time.Unix(expiresAt.Unix(), 0).Add(time.Minute * time.Duration(-5))
	token := jwt.NewWithClaims(
		jwtSigningMethod,
		jwt.MapClaims{
			"Username":  Username,
			"ExpiresAt": expiresAt.Unix(),
			"IssuedAt":  issuedAt,
			"NotBefore": notBefore.Unix(),
		},
	)
	res, err := token.SignedString(jwtSignatureKey)
	if err != err {
		return err
	}

	return []string{res}
}
