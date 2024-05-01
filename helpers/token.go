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
	Id       uuid.UUID
	Email    string
	Fullname string
	jwt.StandardClaims
}

func init() {
	viper.AutomaticEnv()
}

func GenerateAccessToken(user *entities.User) (string, error) {
	accessTokenSecret := []byte(viper.GetString("ACCESS_TOKEN_SECRET"))
	claims := JWTClaims{
		Id:       user.Id,
		Email:    user.Email,
		Fullname: user.Fullname,
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

func GenerateRefreshToken(user *entities.User) (string, error) {
	accessTokenSecret := []byte(viper.GetString("REFRESH_TOKEN_SECRET"))
	claims := JWTClaims{
		Id:       user.Id,
		Email:    user.Email,
		Fullname: user.Fullname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
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

func ParseJWT(tokenStr string) (*uuid.UUID, error) {
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

	return &claims.Id, nil
}
