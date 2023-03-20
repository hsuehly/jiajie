package sign

import (
	"sync"
	"time"

	"go_test/pkg/sign"
)

var once sync.Once
var instance sign.Sign

func Sign(data string) string {
	expire := 4
	if expire == 0 {
		return NotExpired(data)
	} else {
		return WithDuration(data, time.Duration(expire)*time.Hour)
	}
}

func WithDuration(data string, d time.Duration) string {
	once.Do(Instance)
	return instance.Sign(data, time.Now().Add(d).Unix())
}

func NotExpired(data string) string {
	once.Do(Instance)
	return instance.Sign(data, 0)
}

func Verify(data string, sign string) error {
	once.Do(Instance)
	return instance.Verify(data, sign)
}

func Instance() {
	instance = sign.NewHMACSign([]byte("alist-6f66dd20-f10c-49ff-bf9d-2013cc953ad5X5Gyew8v95rbhKdqwTn59zJ2I4J2uG2vKiFHNL8PEAg8VWx1FWFxziX6mboc9K9T"))
}
