/*
1.在一个Golang的项目中通常会有三个目录：src、pkg、bin；src存放源码，pkg存放包，bin存放程序编译后的二进制文件
2.go程序都是以一个包开头，同一个包下面属于同一个工程文件可以直接使用，不用import导入
3.import关键字表示我们所需要导入的包，导入的包必须使用，否则编译器会报错
4.main()为程序的入口
5.go程序运行有两种方法：
	(1)go run codename.go
	(2)go build命令:在任意路径下执行 go install codename,Go编译器会去GOROOT或GOPATH下的src目录寻找源码文件，
	在程序包中找main包的main函数作为程序入口然后编译。然后会在GOPAT目录下创建一个bin目录，并把编译好的二进制文件存放进去
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

/*
go run背后的原理也是生成一份临时的go二进制文件然后执行。
GOROOT=/usr/local/go #gosetup
GOPATH=/Users/burylu/code/Go-Notes:/Users/burylu/goProject #gosetup
/usr/local/go/bin/go build -o /private/var/folders/lw/p__200g50rx2rfj33qcwwzvr0000gn/T/___go_build_01_hello_go__2_ /Users/burylu/code/Go-Notes/01_Go核心编程/01-基础语法/01-hello.go #gosetup
/private/var/folders/lw/p__200g50rx2rfj33qcwwzvr0000gn/T/___go_build_01_hello_go__2_

输出结果：
Hello Golang and I'm coming

*/
