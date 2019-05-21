package main

import (
	"fmt"
	"hpcfusion/service/bash"
)

func main() {
	r, _ := bash.Run("cat /Users/zhijie.fang/go/src/hpcfusion/mock/example")
	fmt.Print(r)
}
