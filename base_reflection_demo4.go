//当前程序的包名
package main

//导入其他的包 
import (
	"fmt"
	"reflect"
)
//如何通过反射来进行方法的调用

//申明User结构
type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello(name string){
	fmt.Println("Hello",name,",my name is",u.Name)
}

func main() {
	u := User{1,"OK",12}
	//u.Hello("Joe")  //方法调用
	//下面的是通过反射"动态"调用方法
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args :=[]reflect.Value{reflect.ValueOf("Joe")}
	mv.Call(args)
}
