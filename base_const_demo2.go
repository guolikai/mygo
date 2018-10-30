package main
import (
	"fmt"
)
//常量用法: 常量一般是大写字母(或首字母大写)
//MAX_COUNT = 1
//内部使用： _MAX_COUNT,cMAX_COUNT其中c代表const
const a int = 1
const b = 'A'
//表达式、常量表达式区别
const (
	c = "123456" //iota = 0
	d = len(c)	 //iota = 1
	e 			 //iota = 2
	f = iota	 //iota = 3
)

//常量iota的值与它定义的顺序有关系,每个方法中都会0开始计算
const (
	a1 = iota //iota = 0
	b1 = len(c)	 //iota = 1
	c1 			 //iota = 2
	d1 = iota	 //iota = 3
)

/* 四位运算符
 6: 0110
11: 1011
-----------------
&	0010 2   #和运算符，2个都是1就是1，1个不是1结果为0
|	1111 15  #或运算符，1个是1结果为1
^	1011 13  #非运算符，对比过程中，两位如果都是1不成立为0，一个为1成立
&^	0100 4   #如果第2个数的一位为1，就把第一个数对应的数强制改为0；
 */


func main() {
	/*
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(f)
		fmt.Println(d)
		fmt.Println(e)
		fmt.Println(f)
		fmt.Println(a1)
		fmt.Println(d1)
	 */

	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)
// && 左边的不成立，右边不执行
	a2 := 0
	if a2 > 0 && (10/a2) > 1{
		fmt.Println("ok")
	}
} 