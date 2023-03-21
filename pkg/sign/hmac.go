package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

type HMACSign struct {
	SecretKey []byte
}

func (s HMACSign) Sign(data string, expire int64) string {
	log.Println(data, expire, "data,express")
	log.Println(s.SecretKey, "key")
	h := hmac.New(sha256.New, s.SecretKey)
	expireTimeStamp := strconv.FormatInt(expire, 10)
	_, err := io.WriteString(h, data+":"+expireTimeStamp)
	if err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(h.Sum(nil)) + ":" + expireTimeStamp
}

func (s HMACSign) Verify(data, sign string) error {
	signSlice := strings.Split(sign, ":")
	log.Println(signSlice[0], "signSlice")
	// check whether contains expire time
	if signSlice[len(signSlice)-1] == "" {
		return ErrExpireMissing
	}
	// check whether expire time is expired
	expires, err := strconv.ParseInt(signSlice[len(signSlice)-1], 10, 64)
	if err != nil {
		return ErrExpireInvalid
	}
	log.Println(expires, "expires")
	// if expire time is expired, return error
	if expires < time.Now().Unix() && expires != 0 {
		return ErrSignExpired
	}
	// verify sign
	log.Println(s.Sign(data, expires), "sign")
	if s.Sign(data, expires) != sign {
		return ErrSignInvalid
	}
	return nil
}

func NewHMACSign(secret []byte) Sign {
	return HMACSign{SecretKey: secret}
}
