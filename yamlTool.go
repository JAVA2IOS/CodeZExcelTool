package main

import (
	"./tool"
	"fmt"
)

func main() {
	k := 4
	switch k {
	case 4:
		fmt.Println("这是1")
		fallthrough
	case 5:
		fmt.Println("这是5")
		fallthrough
	default:
		fmt.Println("这是默认")
	}

	var cg configure.Configure
	fmt.Printf("版本号:%v", cg.Instance().Version)

	var i = 5
	var p *int = &i
	*p = 13

	fmt.Printf("结果是p : %v, %v", *p, p)
}
