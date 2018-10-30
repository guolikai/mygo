package main

import (
	"fmt"
	"sync" 		//同步;WaitGroup
	"runtime" 	//多核执行
)

func Go(wg *sync.WaitGroup,index int){
	a := 1
	for i :=0;i<100000000;i++ {
		a += 1
	}
	fmt.Println(index,a)
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //多核执行
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg,i)
	}
	wg.Wait()
} 
