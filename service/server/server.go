package main

import (
	"fmt"
	"hpcfusion/service/bash"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	val, _ := bash.Run("cat /Users/zhijie.fang/go/src/hpcfusion/mock/example.json")
	num := jsoniter.Get(val.Bytes(), "Jobs").Keys()
	fmt.Println(jsoniter.Get(val.Bytes(), "Jobs", num[0], "resources_used", "cpupercent").ToString())
	fmt.Println(jsoniter.Get(val.Bytes(), "Jobs", num[0], "resources_used", "cput").ToString())
	fmt.Println(jsoniter.Get(val.Bytes(), "Jobs", num[0], "resources_used", "mem").ToString())
	fmt.Println(jsoniter.Get(val.Bytes(), "Jobs", num[0], "resources_used", "ncpus").ToString())
	fmt.Println(jsoniter.Get(val.Bytes(), "Jobs", num[0], "resources_used", "vmem").ToString())
	fmt.Println(jsoniter.Get(val.Bytes(), "Jobs", num[0], "resources_used", "walltime").ToString())
}
