package main

import (
    "fmt"
    // "time"
    // "strconv"
    // "os"
    // "time"
)

// var i5 int = 500

// func outer() {
//     var s4 string = "outer"
//     fmt.Println(s4)
// }

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

// func main() {
//     var s string = "hello golang"
//     fmt.Println(s)

//     fmt.Printf("%T\n", s)

//     var si string = "300"
//     fmt.Println(si)

//     fmt.Println(`test
//     test
//         test
//     `)

//     fmt.Println("\"")
//     fmt.Println(`""`)

//     fmt.Println(s[0]) //72 byte型
//     fmt.Println(string(s[0])) //h
// }

// func main() {
//     byteA := []byte{72, 73}
//     fmt.Println(byteA)

//     fmt.Println(string(byteA))

//     c := []byte("HI")
//     fmt.Println(c)

//     fmt.Println(string(c))
// }

// func main() {
//     var arr1 [3]int
//     fmt.Println(arr1)

//     fmt.Printf("%T\n", arr1)

//     var arr2 [3]string = [3]string{"A", "B"}
//     fmt.Println(arr2)

//     arr3 := [3]int{1, 2, 3}
//     fmt.Println(arr3)

//     arr4 := [...]string{"C"}
//     fmt.Println(arr4)
//     fmt.Printf("%T\n", arr4)

//     fmt.Println(arr1[0])
//     fmt.Println(arr2[0])
//     fmt.Println(arr2[1])
//     fmt.Println(arr2[2])
//     fmt.Println(arr2[2-1]) //要素数-1が最後の値

//     arr2[2] = "C"
//     fmt.Println(arr2)

//     // var arr5 [4]int
//     // arr5 = arr1
//     // fmt.Println(arr5)

//     fmt.Println(len(arr1))

// }

// func main() {
//     var x interface{}
//     fmt.Println(x)

//     x = 1
//     fmt.Println(x)


//     x = 3.14
//     fmt.Println(x)

//     x = [3]int{1, 2, 3}
//     fmt.Println(x)

//     // x = 2
//     // fmt.Println(x + 3)

// }

// func main() {
    // var i int = 1
    // fl64 := float64(i)
    // fmt.Println(fl64)

    // fmt.Printf("i = %T\n", i)
    // fmt.Printf("fl64 = %T\n", fl64)

    // i2 := int(fl64)
    // fmt.Printf("i2 = %T\n", i2)

    // fl := 10.5
    // i3 := int(fl)
    // fmt.Printf("i3 = %T\n", i3)
    // fmt.Println(i3)

    // var s string = "100"
    // fmt.Printf("s = %T\n", s)

    // i, _ := strconv.Atoi(s)
    // fmt.Println(i)

    // fmt.Printf("i = %T\n", i)


//     var i2 int = 200
//     s2 := strconv.Itoa(i2)
//     fmt.Printf("s2 = %T\n", s2)

//     var h string = "hello world"
//     b := []byte(h)
//     fmt.Println(b)

//     h2 := string(b)
//     fmt.Println(h2)

// }

// const pi = 3.14

// // const (
// //     URL = "http://xxx.com",
// //     SiteName = "test"
// // )

// // const (
// //     A = 1
// //     B
// //     C
// //     D = "A"
// //     E
// //     F
// // )

// const (
//     c0 = iota
//     c1
//     c2
// )

// func main() {
//     fmt.Println(pi)
    
//     // fmt.Println(URL)
//     // fmt.Println(SiteName)

//     // fmt.Println(A, B, C, D, E, F)

//     fmt.Println(c0, c1, c2)

// }

//関数を返す関数
// func ReturnFunc() func() {
//     return func() {
//         fmt.Println("I'm function")
//     }
// }

// func main() {

//     //無名関数
//     // f := func(x, y int) int {
//     //     return x + y
//     // }
//     // i := f(1, 2)
//     // fmt.Println(i)

//     // i2 := func(x, y int) int {
//     //     return x + y
//     // }(1, 2)

//     // fmt.Println(i2)

//     f := ReturnFunc()
//     f()

// }


//関数を引数にとる関数
// func Callfunction(f func()) {
//     f()
// }

// func main() {
//     Callfunction(func() {
//         fmt.Println("I'm a function")
//     })
// }

// func Later() func(string) string {
//     var store string
//     return func(next string) string {
//         s := store
//         store = next
//         return s
//     }
// }

// func main() {
//     f := Later()
//     fmt.Println(f("Hello"))
//     fmt.Println(f("My"))
//     fmt.Println(f("name"))
//     fmt.Println(f("is"))
//     fmt.Println(f("Golang"))
// }


// func integers() func() int {
//     //iは実行され続ける限り、iは保持され続ける
//     i := 0
//     return func() int {
//         i++
//         return i
//     }
// }

// func main() {
//     ints := integers()
//     fmt.Println(ints())
//     fmt.Println(ints())
//     fmt.Println(ints())
//     fmt.Println(ints())

//     otherints := integers()
//     fmt.Println(otherints())
//     fmt.Println(otherints())
//     fmt.Println(otherints())
// }


// func main () {
//     a := 0
//     // a = 2
//     if a == 2 {
//         fmt.Println("two")
//     } else if a == 1 {
//         fmt.Println("one")
//     } else {
//         fmt.Println("i don't know")
//     }

//     //簡易文付きif文
//     if b := 100; b == 100 {
//         fmt.Println("one hundred")
//     }

//     x := 0
//     if x := 2; true {
//         fmt.Println(x) //2
//     }
//     fmt.Println(x) //0
// }

//defer 
// func TestDefer() {
//     defer fmt.Println("end")
//     fmt.Println("start")
// }

// func RunDefer() {
//     defer fmt.Println("1")
//     defer fmt.Println("2")
//     defer fmt.Println("3")
// }

// func main() {
//     TestDefer()

//     // defer func() {
//     //     fmt.Println("1")
//     //     fmt.Println("2")
//     //     fmt.Println("3")
//     // }()

//     RunDefer()
    
//     file, err := os.Create("test.txt")
//     if err != nil {
//         fmt.Println(err)
//     }
//     defer file.Close()

//     file.Write([]byte("Hello"))
// }

// func main() {
//     defer func() {
//         if x := recover(); x != nil {
//             fmt.Println(x)
//         }
//     }()

//     panic("runtime error")
//     //runtimeエラーを強制発生
//     fmt.Println("Start")
// }

//並行処理
// func sub() {
//     for {
//         fmt.Println("Sub Loop")
//         time.Sleep(100 * time.Millisecond)
//     }
// }

// func main() {
//     go sub()
//     go sub()

//     for {
//         fmt.Println("Main Loop")
//         time.Sleep(200 * time.Millisecond)
//     }
// }

// func init() {
//     fmt.Println("init")
// }

// func init() {
//     fmt.Println("init2")
// }

// func main() {
//     fmt.Println("Main")
// }

//スライス

// func main() {
//     var sl []int
//     fmt.Println(sl)

//     var sl2 []int = []int{100, 200}
//     fmt.Println(sl2)

//     sl3 := []string{"A", "B"}
//     fmt.Println(sl3)

//     sl4 := make([]int, 5)
//     fmt.Println(sl4)

//     sl2[0] = 1000
//     fmt.Println(sl2)

//     sl5 := []int{1, 2, 3, 4, 5}
//     fmt.Println(sl5)

//     fmt.Println(sl5[0])

//     fmt.Println(sl5[2:4])

//     fmt.Println(sl5[:2])

//     fmt.Println(sl5[2:])

//     fmt.Println(sl5[:])

//     fmt.Println(sl5[len(sl5) - 1])
    
//     fmt.Println(sl5[1:len(sl5) -1])
// }

// func main() {
//     sl := []int{100, 200}
//     fmt.Println(sl)

//     sl = append(sl, 300)
//     fmt.Println(sl)

//     sl = append(sl, 400, 500, 600)
//     fmt.Println(sl)

//     sl2 := make([]int, 5)
//     fmt.Println(sl2)

//     fmt.Println(len(sl2))

//     fmt.Println(cap(sl2))

//     sl3 := make([]int, 5, 10)

//     fmt.Println(len(sl3))
//     fmt.Println(cap(sl3))

//     sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
//     fmt.Println(len(sl3))
//     fmt.Println(cap(sl3))
// }

// func main() {
//     // sl := []int{100, 200}
//     // sl2 := sl
//     // //参照型特有の性質

//     // sl2[0] = 1000
//     // fmt.Println(sl)
//     // //同じsl[0]も更新されてしまう

//     // var i int = 10
//     // i2 := i
//     // i2 = 100
//     // fmt.Println(i)
//     // fmt.Println(i2)

//     sl := []int{1, 2, 3, 4, 5}
//     sl2 := make([]int, 5, 10)
//     fmt.Println(sl2)
//     n := copy(sl2, sl)
//     //nはコピーに成功した数

//     fmt.Println(n, sl2)

// }

func main() {
    sl := []string{"A", "B", "C"}
    fmt.Println(sl)

    // for i, v := range sl {
    //     fmt.Println(i, v)
    // }

    // for _, v := range sl {
    //     fmt.Println(v)
    // }

    // for i := range sl {
    //     fmt.Println(i)
    // }

    for i := 0; i < len(sl); i++ {
        fmt.Println(sl[i])
    }
}