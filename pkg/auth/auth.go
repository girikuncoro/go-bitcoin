package auth

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strconv"
	"time"
)

type Auther interface {
	Sign(string) string
}

type Auth struct {
	ApiKey    string
	SecretKey []byte
	Nonce     string
}

func NewAuth(apiKey, secretKey string) *Auth {
	newAuth := &Auth{
		ApiKey:    apiKey,
		SecretKey: []byte(secretKey),
	}
	newAuth.Nonce = newAuth.currentNonce()
	return newAuth
}

func (a *Auth) Sign(message string) string {
	sig := hmac.New(sha512.New, a.SecretKey)
	sig.Write([]byte(message))
	return hex.EncodeToString(sig.Sum(nil))
}

func (a *Auth) currentNonce() string {
	t := time.Now()
	return strconv.Itoa(int(t.Unix()))
}
