package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func CalcSha256(inputFile, inputPath, outputPath, hashFileName string) (err error) {
	fp, err := os.Open(filepath.Join(inputPath, inputFile))
	if err != nil {
		log.Println("utils/sha256/CalcSha256,while check hash sha256 read input file error:", err)
		return
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Println("utils/sha256/CalcSha256,while check hash sha256 close input file error:", err)
			return
		}
	}()
	hashSha256 := sha256.New()
	if _, err = io.Copy(hashSha256, fp); err != nil {
		log.Println("utils/sha256/CalcSha256,while check hash sha256 copy input file error:", err)
		return
	}
	//当hashSha256变量读取完字节内容后，使用Sum方法传入（nil）
	//在转换成16进制
	hashSha256String := fmt.Sprintf("%x", hashSha256.Sum(nil))
	fmt.Println("input file hash sha256 code is :", hashSha256String)
	//将计算得到的文件散列值写入一个文件中保存
	hashCodeFileFP, err := os.Create(filepath.Join(outputPath, hashFileName))
	if err != nil {
		log.Println("utils/sha256/CalcSha256,create hash file error:", err)
		return
	}
	fileContent := "file name:" + inputFile + "\r\n" + "sha256 code:" + hashSha256String + "\r\n"
	_, err = hashCodeFileFP.Write([]byte(fileContent))
	if err != nil {
		log.Println("utils/sha256/CalcSha256,write hash file error:", err)
		return
	}
	defer func() {
		if err = hashCodeFileFP.Close(); err != nil {
			log.Println("utils/sha256/CalcSha256,close hash file error:", err)
			return
		}
	}()
	return

}
