/*
1.常量是一个简单值得标识符，在程序运行时不会被修改值，常量只可以是布尔类型、数字（整数型、浮点型、复数）和字符串类型。
2.常量声明
	显式类型定义：const a string = "abc"   显式定义时常量左右类型需一致
	隐式类型定义：const b = "big"
3.iota，特殊常量，一个可以被编译器修改的常量,第一个iota的值等于0，往后每行被使用的时候都会加1，当iota被中断后，必须再次显式使用iota才可以按照自增恢复
*/

package main

import (
	"fmt"
)

const a string = "abc"
const b = "big"

const (
	name string = "bob"
	age  int    = 10
	time        //常量组中如果没有指定常量的类型和值，则常量的类型和值与常量组里上一个常量保持一致
)

const (
	q int = iota
	w
	e
	r string = "haha"
	t
	y int = iota
	u
)

func main() {
	fmt.Println(a, b, name, age, time)
	fmt.Println(q, w, e, r, t, y, u)
}

/*
输出结果：
abc big bob 10 10
0 1 2 haha haha 5 6
*/
