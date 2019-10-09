package main

import (
	"fmt"
	"./tool"
)

func main() {
	var cg configure.Configure
	fmt.Printf("版本号:%v", cg.Instance().Version)
}
