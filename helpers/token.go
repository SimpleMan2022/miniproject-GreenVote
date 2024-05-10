package helpers

import (
	"errors"
	"evoting/entities"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"time"
)

type JWTClaims struct {
	Id      uuid.UUID
	Email   string
	IsAdmin bool
	jwt.StandardClaims
}

func init() {
	viper.AutomaticEnv()
}

func GenerateAccessToken(user interface{}) (string, error) {
	accessTokenSecret := []byte(viper.GetString("ACCESS_TOKEN_SECRET"))
	var isAdmin bool

	switch user.(type) {
	case *entities.User:
		isAdmin = false
	case *entities.Admin:
		isAdmin = true
	default:
		return "", errors.New("invalid user type")
	}

	var userID uuid.UUID
	var userEmail string

	switch u := user.(type) {
	case *entities.User:
		userID = u.Id
		userEmail = u.Email
	case *entities.Admin:
		userID = u.Id
		userEmail = u.Email
	default:
		return "", errors.New("invalid user type")
	}

	claims := JWTClaims{
		Id:      userID,
		Email:   userEmail,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(accessTokenSecret)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func GenerateRefreshToken(user interface{}) (string, error) {
	refreshTokenSecret := []byte(viper.GetString("REFRESH_TOKEN_SECRET"))
	var isAdmin bool

	switch user.(type) {
	case *entities.User:
		isAdmin = false
	case *entities.Admin:
		isAdmin = true
	default:
		return "", errors.New("invalid user type")
	}

	var userID uuid.UUID
	var userEmail string

	switch u := user.(type) {
	case *entities.User:
		userID = u.Id
		userEmail = u.Email
	case *entities.Admin:
		userID = u.Id
		userEmail = u.Email
	default:
		return "", errors.New("invalid user type")
	}

	claims := JWTClaims{
		Id:      userID,
		Email:   userEmail,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(refreshTokenSecret)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func ParseJWT(tokenStr string) (*JWTClaims, error) {
	accessTokenSecret := []byte(viper.GetString("ACCESS_TOKEN_SECRET"))
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessTokenSecret, nil
	})

	if err != nil || !token.Valid {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Invalid token signature")
		}
		return nil, errors.New("Your token is expired")
	}

	claims := token.Claims.(*JWTClaims)
	if claims == nil {
		return nil, errors.New("Your token is expired")
	}

	return claims, nil
}
