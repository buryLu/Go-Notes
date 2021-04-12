/*
1.变量：存储特定类型值的内存位置的名称。本质是一小块内存
2.变量声明：使用关键字var，名称由字母、数字、下划线组成；在函数内部可以省略var关键字使用 := 来声明变量，:= 左侧的变量不应该是被声明过的变量，:= 不可以用来申明全局变量和赋值
3.注意事项：
	（1）变量先定义后使用，函数内定义了就必须使用，变量的零值也叫默认值
	（2）变量类型和赋值类型必须一致
	（3）同一个作用域内变量名不能冲突
	（4）简短定义方式，左边变量名必须是一个新名称，且不能定义全局变量
*/

package main

import (
	"fmt"
)

var c, java bool            //声明变量名称和类型不赋值
var a, b = 10, 15           //声明变量名称和赋值，变量类型由变量值自动推导得出
var python, C = true, "yes" //var声明并赋值多个变量，多个变量使用逗号隔开

var ( //集合类型
	name string
	age  int
)

var num int = 1

// request, result := "GET", true

func main() {
	var i int
	request, result := "GET", true
	fmt.Println(i, c, java)
	fmt.Println(a, b)
	fmt.Println(python, C)
	fmt.Println(request, result)
}

/*
输出结果：
0 false false
10 15
true yes
GET true
*/
