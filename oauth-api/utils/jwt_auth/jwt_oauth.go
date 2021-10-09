package jwt_auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/bookstores-go-microservices/oauth-api/dtos"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func CreateToken(userId int64, tokenDetail *dtos.TokenDetailsDto) error {
	accessSecret := os.Getenv("ACCESS_SECRET")
	var err error

	tokenDetail.AtExpires = time.Now().Add(time.Hour * 24).Unix()
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	tokenDetail.AccessUuid = newUUID.String()

	atClaim := jwt.MapClaims{}
	atClaim["authorized"] = true
	atClaim["access_uuid"] = tokenDetail.AccessUuid
	atClaim["user_id"] = userId
	atClaim["exp"] = tokenDetail.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)

	tokenDetail.AccessToken, err = at.SignedString([]byte(accessSecret))
	if err != nil {
		return err
	}

	return nil
}

func CreateRefreshToken(userId int64, tokenDetail *dtos.TokenDetailsDto) error {
	refreshSecret := os.Getenv("REFRESH_SECRET")
	var err error

	tokenDetail.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	tokenDetail.RefreshUuid = newUUID.String()

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tokenDetail.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = tokenDetail.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	tokenDetail.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return err
	}

	return nil
}

func VerifyToken(tokenString string, isRefeshToken bool) (*jwt.Token, error) {
	accessSecret := ""
	if isRefeshToken {
		accessSecret = os.Getenv("REFRESH_SECRET")
	} else {
		accessSecret = os.Getenv("ACCESS_SECRET")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(accessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(tokenString string, isRefeshToken bool) (*dtos.TokenDecodedDto, error) {
	token, err := VerifyToken(tokenString, isRefeshToken)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if isRefeshToken {
			accessUUID, ok := claims["access_uuid"].(string)
			if !ok {
				return nil, err
			}

			userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
			if err != nil {
				return nil, err
			}

			return &dtos.TokenDecodedDto{
				AccessTokenId: accessUUID,
				UserId:        userId,
			}, nil
		} else {
			accessUUID, ok := claims["refresh_uuid"].(string)
			if !ok {
				return nil, err
			}

			userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
			if err != nil {
				return nil, err
			}

			return &dtos.TokenDecodedDto{
				RefeshTokenId: accessUUID,
				UserId:        userId,
			}, nil
		}
	}

	return nil, err
}
