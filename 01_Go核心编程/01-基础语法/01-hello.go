/*
1.在一个Golang的项目中通常会有三个目录：src、pkg、bin；src存放源码，pkg存放包，bin存放程序编译后的二进制文件
2.go程序都是以一个包开头，同一个包下面属于同一个工程文件可以直接使用，不用import导入
3.import关键字表示我们所需要导入的包，导入的包必须使用，否则编译器会报错
4.main()为程序的入口
*/

package main

import (
	//. "fmt"    可以省略包名称，直接调用包中的函数
	//f "fmt"    别名导入
	//_ "fmt"    导入这个包调用init函数，而不使用用里面别的函数
	"fmt"
)

func main() {
	fmt.Println("Hello Golang and I'm coming")

}
