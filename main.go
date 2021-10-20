package main

import (
    "fmt"
    // "time"
)

var i5 int = 500

func outer() {
    var s4 string = "outer"
    fmt.Println(s4)
}

// func main() {
//     fmt.Println("Hello World")
//     fmt.Println(time.Now())

//     var i int = 100
//     fmt.Println(i)

//     var s string = "Hello Go"
//     fmt.Println(s)

//     var t,f bool = true, false
//     fmt.Println(t, f)

//     var (
//         i2 int = 200
//         s2 string = "Golang"
//     )
//     fmt.Println(i2, s2)

//     var i3 int 
//     var s3 string
//     fmt.Println(i3, s3) //0 //空文字

//     i3 = 300
//     s3 = "Go"
//     fmt.Println(i3, s3)

//     i = 150 //値の更新
//     fmt.Println(i)

//     //暗黙的な定義
//     i4 := 400
//     fmt.Println(i4)

//     i4 = 450
//     fmt.Println(i4)


//     fmt.Println(i5)

//     outer()

// }


func main () {
    var i int = 100

    var i64 int64 = 200

    fmt.Println(i + 100)
    
    fmt.Printf("%T\n", i64) //int64

    fmt.Printf("%T\n", int32(i64)) //変換
    //型が異なると計算ができない
}
