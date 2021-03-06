package main

import(
    "fmt"
)

func main(){
    sum := 1
    for ; sum <= 100 ;{
    sum += sum   
    }
    fmt.Println(sum)
    
    fmt.Println("start excute forIswhile")
    forIswhile()
}

func forIswhile(){
    sum := 1
    for sum <=1000 {
        sum += sum
    }
    fmt.Println(sum)
}
