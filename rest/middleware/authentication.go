package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (middlewares *Middlewares) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 || headerArr[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		accessToken := headerArr[1]
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]
		if jwtHeader == "" || jwtPayload == "" || jwtSignature == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		message := jwtHeader + "." + jwtPayload
		byteArrMessage := []byte(message)
		expectedSignature := hmac.New(sha256.New, []byte(middlewares.config.SecretKey))
		expectedSignature.Write(byteArrMessage)
		hash := expectedSignature.Sum(nil)
		encodedHash := base64UrlEncode(hash)
		if encodedHash != jwtSignature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return

		}

		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
