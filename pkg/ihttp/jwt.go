package ihttp

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// TokenService proporciona funciones para generar y validar tokens JWT
type TokenService struct {
	secretKey string
	expiry    time.Duration
}

// NewTokenService crea una nueva instancia de TokenService
func NewTokenService(secretKey string, expiry time.Duration) *TokenService {
	return &TokenService{
		secretKey: secretKey,
		expiry:    expiry,
	}
}

type JwtClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(email, role string) (string, error) {

	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(viper.GetString("token.access_token")))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el algoritmo de firma sea HMAC y la clave coincida
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(viper.GetString("token.access_token")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
