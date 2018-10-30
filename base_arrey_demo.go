package main

import (
	"fmt"
)

//Go中数组冒泡排序
func main() {
	a := [...]int{1,8,3,12,520,1314}
	fmt.Println(a)
	num := len(a)
	for i:=0;i<num;i++ {
		for j:=i+1;j<num;j++ {
			if a[i] < a[j] {
				temp := a[i] 
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
