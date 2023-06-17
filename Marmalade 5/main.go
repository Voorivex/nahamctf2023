package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
	"sync"
)

func jwtcreator(secret string) string {
	headerBase64 := "eyJhbGciOiJNRDVfSE1BQyJ9"
	payloadBase64 := "eyJ1c2VybmFtZSI6ImFsaV9raGFsa2hhbGkwIn0"

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
	// Open the 5-characters.txt file for reading.
	file, err := os.Open("5-characters.txt")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup
	found := false // boolean variable to track if the desired JWT is found
	for scanner.Scan() {
		if found { // if the desired JWT is found, break out of the loop
			break
		}
		wg.Add(1)
		go func(line string) {
			tmpJWT := jwtcreator("fsrwjcfszeg" + line)
			// If the tmpJWT is equal to the existing JWT, print it to stdin and set found to true
			if tmpJWT == "eyJhbGciOiJNRDVfSE1BQyJ9.eyJ1c2VybmFtZSI6ImFsaV9raGFsa2hhbGkwIn0.81w2OhAsiyAswIr4q6J3VA" {
				fmt.Println("fsrwjcfszeg" + line)
				found = true
			}
			wg.Done()
		}(scanner.Text())
	}
	wg.Wait()

	// Check if there were any errors while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
}