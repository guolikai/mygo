package main

import (
	"fmt"
	"time"
)
//如果希望修改slice，最好在函数中设置返回值
func Pingpong(s []int) []int {
	s = append(s,3)
	return s
}


func main() {
	s := make([]int,0)
	fmt.Println(s)
	s = Pingpong(s)
	fmt.Println(s)

//time.Format时间不能修改	
	t := time.Now()
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("Mon Jan _2 15:04:06 2006"))

//for range闭包的坑: slice应用的是地址，for已经是最后的值
	s1 := []string{"a","b","c"}
	for _,v := range s1 {
		go func(){
			fmt.Println(v)
		}()
	}

	s2 := []string{"a","b","c"}
	for _,v := range s2 {
		go func(v string){
			fmt.Println(v)
		}(v)
	}
}
