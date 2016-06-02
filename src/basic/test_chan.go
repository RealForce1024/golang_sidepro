package main

import "fmt"

func main(){
	ch1 := make(chan int,5)
	value := <- ch1
     	fmt.Println(value)

}
