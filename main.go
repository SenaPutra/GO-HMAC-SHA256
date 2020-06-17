/**
 * @project HMAC-SHA256
 * @author sena on 17/06/20
 */

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	secret := "8b125cc228adcc57"


	// TODO the post request method. body data
	requestBody := "{\"orderId\":\"12345678\"}"
	rawValue := "1528097319441" + ""
	rawValue = rawValue + "$" + requestBody
	//sign: xdcITyQR92qO1kY0q1rhkBvLmmH/HKNdBlXOehbo43o=
	// TODO Encoded to url parameter
	fmt.Println(ComputeHmac256(rawValue, secret))
}


