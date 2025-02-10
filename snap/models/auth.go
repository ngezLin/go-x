package models

import "time"

type (
	AuthB2BRequest[T any] struct {
		GrantType      string `json:"grantType"`
		AdditionalInfo T      `json:"additionalInfo"`
	}
	AuthB2BResponse[T any] struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		AccessToken     string `json:"accessToken"`
		TokenType       string `json:"tokenType"`
		ExpiresIn       int    `json:"expiresIn"`
		AdditionalInfo  T      `json:"additionalInfo"`
	}
)

type (
	AuthB2B2CRequest[T any] struct {
		GrantType      string `json:"grantType"`
		AuthCode       string `json:"authCode"`
		RefreshToken   string `json:"refreshToken"`
		AdditionalInfo T      `json:"additionalInfo"`
	}
	AuthB2B2CResponse[T any] struct {
		ResponseCode           string    `json:"responseCode"`
		ResponseMessage        string    `json:"responseMessage"`
		TokenType              string    `json:"tokenType"`
		AccessToken            string    `json:"accessToken"`
		AccessTokenExpiryTime  time.Time `json:"accessTokenExpiryTime"`
		RefreshToken           string    `json:"refreshToken"`
		RefreshTokenExpiryTime time.Time `json:"refreshTokenExpiryTime"`
		AdditionalInfo         T         `json:"additionalInfo"`
	}
)
