/*
1.在一个Golang的项目中通常会有三个目录：src、pkg、bin；src存放源码，pkg用于存放安装库源码后的归档文件，bin存放程序编译后的二进制文件
2.go程序都是以一个包开头，同一个包下面属于同一个工程文件可以直接使用，不用import导入
3.import字表示我们所需要导入的包，导入的包必须使用，否则编译器会报错
4.main()为程序关键的入口
5.go程序运行有两种方法：
	(1)go run codename.go
	(2)go build命令:在任意路径下执行 go install codename,Go编译器会去GOROOT或GOPATH下的src目录寻找源码文件，
	在程序包中找main包的main函数作为程序入口然后编译。然后会在GOPAT目录下创建一个bin目录，并把编译好的二进制文件存放进去
*/

/*
6.GO源码文件分类
	(1)命令源码文件:命令源码文件被安装后，如果只有一个GOPATH工作区，那么可执行文件会被安装到当前工作区的bin目录下；如果有多个工作区则会被安装到GOBIN的目录下。
	   同一个代码包中不要存放多个命令源码，go run 可以跑起来，但是go build 和go install会失败。
	(2)库源码文件:库源码文件不存在命令源码文件的两个特征，库源码文件被安装后相应的归档文件(.a)会被存放到当前工作区的pkg目录下
	(3)测试源码文件:名称以_test.go为后缀的代码文件，并且必须包含Test或者Benchmark名称为前缀的函数，这些函数的功能就是测试函数
*/

/*
7.Go常用命令有:build、install、get、run
Go常用的选项有
名称	说明
-a		用于强制重新编译所有涉及的 Go 语言代码包（包括 Go 语言标准库中的代码包），即使它们已经是最新的了。该标记可以让我们有机会通过改动底层的代码包做一些实验。
-n		使命令仅打印其执行过程中用到的所有命令，而不去真正执行它们。如果不只想查看或者验证命令的执行过程，而不想改变任何东西，使用它正好合适。
-race	用于检测并报告指定 Go 语言程序中存在的数据竞争问题。当用 Go 语言编写并发程序的时候，这是很重要的检测手段之一。
-v		用于打印命令执行过程中涉及的代码包。这一定包括我们指定的目标代码包，并且有时还会包括该代码包直接或间接依赖的那些代码包。这会让你知道哪些代码包被执行过了。
-work	用于打印命令执行时生成和使用的临时工作目录的名字，且命令执行完成后不删除它。这个目录下的文件可能会对你有用，也可以从侧面了解命令的执行过程。如果不添加此标记，那么临时工作目录会在命令执行完毕前删除。
-x		使命令打印其执行过程中用到的所有命令，并同时执行它们。
*/

/*
8.Go常用命令解析
	（1）go run
	（2）go build
	（3）go install
	（4）go get
	（5）go clean
	（6）go fmt
	（7）go test
	（8）go doc
*/

package main

import (
	//. "fmt"    可以省略包名称，直接调用包中的函数
	//f "fmt"    别名导入
	//_ "fmt"    导入这个包调用init函数，而不使用用里面别的函数
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {

}

func BenchmarkMethod(t *testing.B) {

}

func main() {
	fmt.Println("Hello Golang and I'm coming")

}

/*
GOROOT=/usr/local/go #gosetup
GOPATH=/Users/burylu/code/Go-Notes:/Users/burylu/goProject #gosetup
/usr/local/go/bin/go build -o /private/var/folders/lw/p__200g50rx2rfj33qcwwzvr0000gn/T/___go_build_01_hello_go__2_ /Users/burylu/code/Go-Notes/01_Go核心编程/01-基础语法/01-hello.go #gosetup
/private/var/folders/lw/p__200g50rx2rfj33qcwwzvr0000gn/T/___go_build_01_hello_go__2_

输出结果：
Hello Golang and I'm coming

*/
