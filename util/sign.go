package util

import (
	"crypto/md5"
	"fmt"
)

func GenSig(baseUrl string) (sig string) {
	srcCode := md5.Sum([]byte(baseUrl))
	sig = fmt.Sprintf("%x", srcCode)
	return
}
