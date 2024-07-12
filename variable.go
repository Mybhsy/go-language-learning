/*

 */

package main

import "fmt"

var q1 = 100
var q2 = 200
var name2 = "jack"

//一次性声明
var (
	w1 = 100
	w2 = 200
	name3 = "mary"
)

func main(){
	//定义变量/声明变量
	var i int
	var a int

	//根据值自行判断变量类型（类型推导）
	var num = "20"

	//省略var 等价于 var name string 
	// name = "tom"
	name := "tom"

	//一次性声明多个变量
	var n1, n2, n3 int
	var s1, s2, s3 = 100, "tom", 888
	var a1, a2, a3 = 100, "tom", 888

	//给i赋值
	i = 10
	
	//使用变量
	fmt.Println("i =", i)

	//未赋值则使用默认值，默认值为 0
	fmt.Println("a =", a)
	fmt.Println("num =", num)
	fmt.Println("name =", name)
	fmt.Println("n1=", n1, "n2=", n2, "n3", n3)
	fmt.Println("s1=", s1, "s2=", s2, "s3", s3)
	fmt.Println("a1=", a1, "a2=", a2, "a3", a3)

	//输出全局变量
	fmt.Println("q1=", q1, "q2=", q2, "name2", name2)
	fmt.Println("w1=", w1, "w2=", w2, "name3", name3)

}