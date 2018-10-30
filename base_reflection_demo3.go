//当前程序的包名
package main

//导入其他的包 
import (
	"fmt"
	"reflect"
)
//如何通过反射来进行方法的调用

func A(){
	x := 123
	v := reflect.ValueOf(&x)
	//Elem()方法取得实际的对象
	v.Elem().SetInt(999)
	fmt.Println(x)
}

//申明User结构
type User struct {
	Id int
	Name string
	Age int
}



func main() {
	//A()
	u := User{1,"OK",12}
	//Info(u)   //值拷贝
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}){
	//v.Kind()做类型判断
	v := reflect.ValueOf(&o)
	if v.Kind()==reflect.Ptr && !v.Elem().CanSet(){
		fmt.Println("XX")
		return
	}else{
		v = v.Elem()

	}
//	if f := v.FieldByName("Name");v.Kind()==reflect.String{
//		f.SetString("XiuGai")
//	}
	f := v.FieldByName("Name")
	if !f.IsValid(){
		f.SetString("No")
		return
	}
	if f.Kind()==reflect.String{
		f.SetString("XiuGai")
	}
}
