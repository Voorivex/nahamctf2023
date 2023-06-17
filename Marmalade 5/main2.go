package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func jwtcreator(secret string) string {
	headerBase64 := "eyJhbGciOiJNRDVfSE1BQyJ9"
	payloadBase64 := "eyJ1c2VybmFtZSI6ImFkbWluIn0" // {"username":"admin"}

	// Concatenate the header and payload with a period separator.
	unsignedToken := headerBase64 + "." + payloadBase64
	secretKey := []byte(secret)

	// Sign the token with the secret key using MD5_HMAC.
	mac := hmac.New(md5.New, secretKey)
	mac.Write([]byte(unsignedToken))
	signature := mac.Sum(nil)

	// Encode the signature as Base64Url.
	signatureBase64 := base64.RawURLEncoding.EncodeToString(signature)

	// Concatenate the unsigned token and signature with a period separator to form the final JWT.
	token := unsignedToken + "." + signatureBase64

	return token
}

func main() {

	adminJWT := jwtcreator("fsrwjcfszegvsyfa")
	fmt.Println(adminJWT)
}