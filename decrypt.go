package main

import (
	"EncAndDec/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	seed, _ := strconv.ParseInt(os.Args[1], 10, 64)
	encryptFile := "encrypt.enp"
	outputFile := "b.zip"
	if len(os.Args) == 4 {
		encryptFile = os.Args[2]
		outputFile = os.Args[3]
	}
	err := utils.DecryptByte(encryptFile, "./", outputFile, "./", seed)
	if err != nil {
		panic(err)
	}
	fmt.Printf("decrypt the encrypted File \"%s\" to normal file \"%s successfully:\n", encryptFile, outputFile)
}
