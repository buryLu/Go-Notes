package main

import(
    "fmt"
)

func array(){
    var a [2]string
    a[0] = "hello"
    a[1] = "Golang"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    age := [5]int{20, 21, 22, 23, 24}
    fmt.Println(age)
}

func array_slice(){
    b := [6]int{1, 2, 3, 4, 5, 6} 
    var c []int = b[1:4]
    c[0] = 100
    fmt.Println(c)
    fmt.Println(b)
    
    s := []struct{
        a int
        b string
    } {{10, "zhangsan"}, {11, "lisi"}, {12, "wang5"}}

    fmt.Println(s)
}

func main(){
    fmt.Println("start func array")
    array()
    fmt.Println("start func array_slice")
    array_slice()
}
