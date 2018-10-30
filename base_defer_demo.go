package main
import (
	"fmt"
)
//defer 在函数体执行结束后按照调用顺序的相反顺序逐个执行

func A() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func B(){
	for i := 0; i < 5; i++ {
		defer func(){
			//i作为局部变量引用，在退出循环体时已经是5；defer之后都是5
			fmt.Println(i)
		}()
	}
}

func main() {
	//fmt.Println("a")
	//defer fmt.Println("b")
	//defer fmt.Println("c")
	//A()
	//B()
	//C()
	//D()
	//E()  //E不执行
//-----------------
	C()
	D1()
	E()  
}

//Go 没有异常机制，但有 panic/recover 模式来处理错误
//Panic 可以在任何地方引发，但recover只有在defer调用的函数中有效
func C() {
	fmt.Println("Func C")
}

func D() {
	panic ("Panic In D")

}

func  D1() {
	defer func(){
		if err :=recover();err !=nil {
			fmt.Println("Recover D")
		}
	}()
	panic ("Panic In D")

}

func E() {
	fmt.Println("Func E")
}