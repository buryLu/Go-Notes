package main

import(
    "fmt"

)

func slice(){
    s := []int{1, 2, 3, 4, 5}
    fmt.Println(s, len(s), cap(s))

    a := s[:0]
    b := s[:4]
    c := s[2:]

    fmt.Printf("a:%v\nb:%v\nc:%v\n", a, b, c)
}

func nil_slices(){
    var s []int
    fmt.Println(s, len(s), cap(s))
    if s == nil{
        fmt.Println("nil!")
    }
}

func main(){
    slice()
    nil_slices()

}
