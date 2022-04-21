package main

import (
	"fmt"

	"github.com/yangtizi/go/sysutils"
)

func main() {
	fmt.Println(sysutils.FloatToStr(1.20, 2))   // 1.2
	fmt.Println(sysutils.FloatToStr(1.2222, 2)) // 1.22
	fmt.Println(sysutils.FloatToStr(1.2022, 2)) // 1.2
	fmt.Println(sysutils.FloatToStr(1.000, 2))  // 1
}
