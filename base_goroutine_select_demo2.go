package main


import (
	"fmt"
)


func main() {
	c := make(chan int)
	go func(){
		for v := range c {
			fmt.Println(v)
		}
		
	}()
	
	for {
		select{   //select作为发送者
		case c <- 0:
		case c <- 1:
		}
	}
} 