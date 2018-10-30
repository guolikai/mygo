package main


import (
	"fmt"
)

func Pingpong(){
	i := 0
	for {
		fmt.Println(<-c)
		//fmt.Spintf返回string类型格式化后的的字符串
		c <-  
		fmt.Spintf("From Pingpong:Hi,#%d",i)
		i++
	}
}


func main() {
	c := make(chan string)
	go Pingpong()
	for i := 0; i < 10; i++ {
		c <- fmt.Spintf("From Main:Hello,#%d",i)
		fmt.Println(<-c)
	}
} 
