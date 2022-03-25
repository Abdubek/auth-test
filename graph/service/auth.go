package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Abdubek/auth-test/graph/model"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type contextKey struct {
	name string
}

type AccessDetails struct {
	UserId int
}

var userCtxKey = &contextKey{"user"}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := ExtractToken(authHeader)
			ad, err := ExtractTokenMetadata(tokenString)
			if err != nil {
				if err != errors.New("INVALID_TOKEN") {
					log.Println(err.Error())
				}
				next.ServeHTTP(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), userCtxKey, ad)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ExtractToken(bearToken string) string {
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ForContext(ctx context.Context) *AccessDetails {
	raw, _ := ctx.Value(userCtxKey).(*AccessDetails)
	return raw
}

func ExtractTokenMetadata(tokenString string) (*AccessDetails, error) {
	secret := "secret"

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if token == nil {
		return nil, errors.New("INVALID_TOKEN")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userId, ok := claims["id"].(float64)
		if !ok {
			return nil, errors.New("INVALID_TOKEN")
		}

		ad := &AccessDetails{
			UserId: int(userId),
		}

		return ad, nil
	}
	return nil, errors.New("INVALID_TOKEN")
}

func CreateToken(userId int) (*model.Token, error) {
	token := &model.Token{}
	var err error

	accessExp, err := strconv.Atoi("60")
	if err != nil {
		return nil, err
	}
	refreshExp, err := strconv.Atoi("1440")
	if err != nil {
		return nil, err
	}
	accessTokenExp := time.Now().Add(time.Minute * time.Duration(accessExp)).Unix()
	refreshTokenExp := time.Now().Add(time.Minute * time.Duration(refreshExp)).Unix()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["id"] = userId
	accessTokenClaims["iat"] = time.Now().Unix()
	accessTokenClaims["exp"] = accessTokenExp
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	at, err := accessToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	token.AccessToken = &at

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["id"] = userId
	refreshTokenClaims["iat"] = time.Now().Unix()
	refreshTokenClaims["exp"] = refreshTokenExp
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	token.RefreshToken = &rt

	return token, nil
}
