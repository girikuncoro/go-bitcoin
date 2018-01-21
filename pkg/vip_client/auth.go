package vip_client

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strconv"
	"time"
)

type Auther interface {
	Sign(string) string
	Nonce() string
}

type Auth struct {
	ApiKey    string
	SecretKey []byte
}

func NewAuth(apiKey, secretKey string) *Auth {
	newAuth := &Auth{
		ApiKey:    apiKey,
		SecretKey: []byte(secretKey),
	}
	return newAuth
}

func (a *Auth) Sign(message string) string {
	sig := hmac.New(sha512.New, a.SecretKey)
	sig.Write([]byte(message))
	return hex.EncodeToString(sig.Sum(nil))
}

func (a *Auth) Nonce() string {
	t := time.Now()
	return strconv.Itoa(int(t.Unix()))
}
