package basic

import(
	"fmt"
"time"
)

func main(){
	
	type Sender chan <- int
	type Receiver <- chan int
	
	var myChannel = make(chan int,0)
	var number = 6
	go func(){
		var sender Sender = myChannel
		sender <- number		
		fmt.Printf("sent!->%d\n",number)	
	}()

	go func(){
		var receiver Receiver = myChannel
		fmt.Printf("receiver!<-%d\n",<- receiver)
	}()

time.Sleep(time.Second)
}










