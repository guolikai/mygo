//当前程序的包名
package main

//导入其他的包 
import (
	"fmt"
	"reflect"
)
//反射会将匿名字段当成独立的字段来处理
//申明User结构
type User struct {
	Id int
	Name string
	Age int
}
//嵌入字段
type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User:User{1,"OK",12},title:"124"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n",t.Field(0))
	fmt.Printf("%#v\n",t.Field(1))
	fmt.Printf("%#v\n",t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v\n",t.FieldByIndex([]int{0, 1}))
	fmt.Printf("%#v\n",t.FieldByIndex([]int{0, 2}))

}
