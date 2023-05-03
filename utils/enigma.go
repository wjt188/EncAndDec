package utils

import (
	"log"
	"os"
	"path/filepath"
)

func EncryptByte(inputFile, inputPath, outputFile, outputPath string, seed int64) (err error) {
	fileBytes, err := os.ReadFile(filepath.Join(inputPath, inputFile))
	if err != nil {
		log.Println("utils/enigma.go/EncryptByte:read input file error:", err)
		return
	}
	outputFileBytes := make([]byte, 0)
	m := GenByteMap(seed)
	for i, _ := range fileBytes {
		outputFileBytes = append(outputFileBytes, m[fileBytes[i]])
	}
	confuseBytesTop, err := GenConfuseBytes(77)
	if err != nil {
		log.Println("utils/enigma.go/EncryptByte:gen confuse file error:", err)
		return
	}
	confuseBytesTail, err := GenConfuseBytes(66)
	if err != nil {
		log.Println("utils/enigma.go/EncryptByte:gen confuse file error:", err)
		return
	}
	outputFileBytes = append(confuseBytesTop, outputFileBytes...)
	outputFileBytes = append(outputFileBytes, confuseBytesTail...)

	outputFP, err := os.Create(filepath.Join(outputPath, outputFile))
	if err != nil {
		log.Println("utils/enigma.go/EncryptByte:create output file error:", err)
		return
	}
	_, err = outputFP.Write(outputFileBytes)
	if err != nil {
		log.Println("utils/enigma.go/EncryptByte:write output file error:", err)
		return
	}
	hashFileName := "sha256_origin.txt"
	CalcSha256(inputFile, inputPath, outputPath, hashFileName)
	return

}

func DecryptByte(inputFile, inputPath, outputFile, outputPath string, seed int64) (err error) {
	fileBytes, err := os.ReadFile(filepath.Join(inputPath, inputFile))
	if err != nil {
		log.Println("utils/enigma.go/DecryptByte:read input file error:", err)
		return
	}
	l := len(fileBytes)
	fileBytesCore := fileBytes[77 : l-66]
	outputFileBytes := make([]byte, l-77-66, l-77-66)
	m := GenByteMap(seed)
	n := ReverseByteMap(m)
	for i, _ := range fileBytesCore {
		outputFileBytes[i] = n[fileBytesCore[i]]
	}
	outputFP, err := os.Create(filepath.Join(outputPath, outputFile))
	if err != nil {
		log.Println("utils/enigma.go/DecryptByte:create output file error:", err)
		return
	}
	_, err = outputFP.Write(outputFileBytes)
	if err != nil {
		log.Println("utils/enigma.go/DecryptByte:write output file error:", err)
		return
	}
	hashFileName := "sha256_origin.txt"
	CalcSha256(inputFile, inputPath, outputPath, hashFileName)
	return

}
