package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("Map:%v\n", m)
	info = fmt.Sprintf("Osinfo:%s,archInfo:%s\n", runtime.GOOS, runtime.GOARCH)
}

var m map[int]string = map[int]string{1: "a", 2: "b", 3: "c"}
var info string

func main() {
	fmt.Printf("info:%s", info)
}
