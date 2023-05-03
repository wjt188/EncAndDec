package main

import (
	"EncAndDec/utils"
	"os"
	"strconv"
)

func main() {
	sercetNum, _ := strconv.ParseInt(os.Args[1], 10, 64)
	fileName := os.Args[2]

	utils.EncryptByte(fileName, "./", "encrypt.enp", "./", sercetNum)
}

