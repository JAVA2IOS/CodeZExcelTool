package tool

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func UTF8(s string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewDecoder())

	utf8String, err := ioutil.ReadAll(reader)

	if err != nil {
		return ""
	}

	return string(utf8String)
}

func GBK(s string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewEncoder())

	gbkString, err := ioutil.ReadAll(reader)

	if err != nil {
		return ""
	}

	return string(gbkString)
}