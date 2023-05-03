package utils

import (
	crand "crypto/rand"
	"log"
	"math/big"
)

//生成随机字节
func GenRand(maxInt int64) (b byte, err error) {
	r, err := crand.Int(crand.Reader, big.NewInt(maxInt))
	b = byte(r.Int64())
	return
}

//生成随机字节切片，用于进行文件内容的混淆
func GenConfuseBytes(n uint) (cb []byte, err error) {
	cb = make([]byte, n, n)
	for i, _ := range cb {
		b, errByte := GenRand(256)
		if errByte != nil {
			log.Println("utils/rand.go/GenConfuseBytes:generate random bytes error:", errByte)
			err = errByte
			return
		}
		cb[i] = b
	}
	return
}
