package main

import(
	"fmt" //notice _ not used
	"runtime"
)

func main(){
	
	println("hello world!")
	fmt.Printf("version:%s\n",runtime.Version())
}


