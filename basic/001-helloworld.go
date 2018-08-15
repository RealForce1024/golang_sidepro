package basic

import (
	"fmt"
	"runtime"
)

func main() {
	println("hello go")
	fmt.Printf("version: %s\n", runtime.Version())
}
