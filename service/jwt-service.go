package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(UserID string, Email string, Profile string, Jk string, telephone string, pin string, name string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID  string `json:"userid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Profile string `json:"profile"`
	Telp    string `json:"telp"`
	Pin     string `json:"pin"`
	Jk      string `json:"jk"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "aminivan",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "aminivan"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string, Email string, Profile string, Jk string, Telephone string, Pin string, Name string) string {
	claims := &jwtCustomClaim{
		UserID,
		Name,
		Email,
		Profile,
		Telephone,
		Pin,
		Jk,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 3, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
