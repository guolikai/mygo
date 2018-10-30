package main


import (
	"fmt"
	//"time"
	"runtime"
)

func A() {
	//没缓存，是同步阻塞的
	c := make(chan bool)  //make(chan bool) 是双向通道,能存能取
	go func(){
		fmt.Println("Go GO GO!!!")
		c <- true  //存的操作
	}()
	<-c   //取的操作
}

func A1() {
	//有缓存
	c := make(chan bool,10)  //make(chan bool) 是双向通道,能存能取
	go func(){
		fmt.Println("Go GO GO!!!")
		c <- true  //存的操作
	}()
	<-c   //取的操作
}

func B() {
	c := make(chan bool)
	go func(){
		fmt.Println("Go GO GO!!!")
		c <- true  	//存的操作
		close(c)  	//通知chan关闭
	}()
	for v := range c {
		fmt.Println(v)	
	}
}

func Go(c chan bool,index int){
	a := 1
	for i :=0;i<100000000;i++ {
		a += 1
	}
	fmt.Println(index,a)
	c <- true
}

func Go1(c chan bool,index int){
	a := 1
	for i :=0;i<100000000;i++ {
		a += 1
	}
	fmt.Println(index,a)
	if index == 9 {  //单线程执行
		c <- true
	}
}

func main() {
	//time.Sleep(2 * time.Second)
	//A()
	//B()
	
	runtime.GOMAXPROCS(runtime.NumCPU()) //多核执行
	c := make(chan bool,10)
	for i := 0; i < 10; i++ {
		go Go(c,i)
	}
	
	for i := 0; i < 10; i++ {
		<-c
	}
}
