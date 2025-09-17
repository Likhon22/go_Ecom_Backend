package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}
type Signature struct{}

func CreateJwt(secret string, data Payload) (string, error) {

	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	byteArrPayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	headerB64 := base64UrlEncode(byteArrHeader)
	payloadB64 := base64UrlEncode(byteArrPayload)

	message := headerB64 + "." + payloadB64
	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)
	hash := hmac.New(sha256.New, byteArrSecret)

	hash.Write(byteArrMessage)

	signature := hash.Sum(nil)

	signatureB64 := base64UrlEncode(signature)
	jwt := headerB64 + "." + payloadB64 + "." + signatureB64
	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
