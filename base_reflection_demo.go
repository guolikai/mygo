//当前程序的包名
package main

//导入其他的包 
import (
	"fmt"
	"reflect"
)
//申明User结构
type User struct {
	Id int
	Name string
	Age int
}

//申明Hello方法
func (u User) Hello(){
	fmt.Println("Hello World.")
}

//空接口
//
func Info(o interface{}){
//获取类型信息
	t :=  reflect.TypeOf(o)
	fmt.Println("Type:",t.Name())
    
	if k := t.Kind();k != reflect.Struct{
		fmt.Println("XX")
		return
		}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
//获取字段信息
	for i := 0; i < t.NumField();i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n",f.Name,f.Type,val)
	}
//获取方法信息
	for i :=0;i < t.NumMethod();i++ {
		f := t.Method(i)
		fmt.Printf("%s: %v\n",f.Name,f.Type)
	}

}

func main() {
	u := User{1,"OK",12}
	//Info(u) //值拷贝
	Info(&u)
}

