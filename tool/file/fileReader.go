package fileReader

import (
	"io/ioutil"
	"log"
	"bytes"
	"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func ReadTxtFile(file string) []string {
	newFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("读取txt文件失败[%v]\n", err.Error())
		return nil
	}

	reader := transform.NewReader(bytes.NewReader(newFile), simplifiedchinese.GBK.NewDecoder())

	newBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil
	}

	charString := string(newBytes)
	charString = strings.ReplaceAll(charString, "\r", "")

	newString := []string{}
	newString = strings.Split(charString, "\n")

	return newString

}