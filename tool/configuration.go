package configure

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

const (
	yamlFilePath = "config.yaml"
)

type Configure struct {
	Version string `yaml:"version"`
}

func (c *Configure) Instance() *Configure {
	// 读取yaml文件
	yamlFile, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		fmt.Printf("yaml文件[%v]获取错误[%v]\n", yamlFilePath, err.Error())
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		fmt.Printf("读取数据错误[%v]", err)
	}

	return c
}