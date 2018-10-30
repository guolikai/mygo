package main
import (
	"fmt"
)
//常量用法
func A(s ...int){
	s[0] = 3
	s[1] = 4
	fmt.Println(s)
}

//函数中的slice是地址的拷贝。string、int参数常规传递是值的拷贝。
func B(c *int){
	*c = 110
	fmt.Println(*c)
}

// go语言也是一种类型
func D(){
	fmt.Println("Func D")
}





func main() {
	a,b := 1,2
	A(a,b)
	fmt.Println(a,b)
	c := 1
	fmt.Println(c)
	B(&c)

	d := D
	d()
//匿名函数
	e := func() {
		fmt.Println("Func e")
	}
	e()
//闭包,作用返回一个匿名函数
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))

}
//fmt.Printf 传递参数
func closure(x int) func (int) int{
	fmt.Printf("%p\n",&x)
	return func(y int) int {
		fmt.Printf("%p\n",&x)
		return x + y
	}
}
