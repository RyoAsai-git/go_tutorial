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


// func main () {
//     var i int = 100

//     var i64 int64 = 200

//     fmt.Println(i + 100)
    
//     fmt.Printf("%T\n", i64) //int64

//     fmt.Printf("%T\n", int32(i64)) //変換
//     //型が異なると計算ができない
// }


// func main() {
//     var fl64 float64 = 2.4
//     fmt.Println(fl64)

//     fl := 3.2
//     fmt.Println(fl64 + fl)
//     fmt.Printf("%T", "%T\n", fl64, fl)

//     var fl32 float32 = 1.2
//     fmt.Printf("%T\n", fl32)

//     fmt.Printf("%T\n", float64(fl32))

//     zero := 0.0
//     pinf := 1.0 / zero
//     fmt.Println(pinf)

//     ninf := -1.0 / zero
//     fmt.Println(ninf)

//     nan := zero / zero
//     fmt.Println(nan)

// }

// func main() {
//     var t, f bool = true, false
//     fmt.Println(t, f)

    
// }

func main() {
    var s string = "hello golang"
    fmt.Println(s)

    fmt.Printf("%T\n", s)

    var si string = "300"
    fmt.Println(si)

    fmt.Println(`test
    test
        test
    `)

    fmt.Println("\"")
    fmt.Println(`""`)

    fmt.Println(s[0]) //72 byte型
    fmt.Println(string(s[0])) //h
}