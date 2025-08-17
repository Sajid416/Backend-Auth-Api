package helper

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/sajid414/Go-react-auth/model"
)

type CustomClaims struct {
	Email  string `json:"email"`
	UserID string `json:"userId"` // ✅ secure string userId
	jwt.RegisteredClaims
}

var secret string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret = os.Getenv("JWT_SECRET") // assign once when package loads
}
func GenerateToken(user model.User) (string, error) {
	claims := CustomClaims{
		Email:  user.Email,
		UserID: user.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(72 * time.Hour)), // 3 days
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    "blog_app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Println("Error in token Signing")
		return "", err
	}

	return t, nil

}

func ValidateToken(clientToken string) (claims *CustomClaims, msg string) {
	token, err := jwt.ParseWithClaims(clientToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}

	if !token.Valid {
		msg = "Invalid or expired token"
		return
	}

	parsedClaims, ok := token.Claims.(*CustomClaims)
	if !ok {
		msg = "Invalid token claims type"
		return
	}

	claims = parsedClaims
	return claims, ""
}
