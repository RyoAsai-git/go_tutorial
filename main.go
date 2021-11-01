// package main

// import (
//     "database/sql"
//     "log"
//     _ "github.com/lib/pq"

// )

// var Db *sql.DB

// var err error

// type Person struct {
//     Name string
//     Age int
// }

// func main() {
//     //userとdbをここで設定
//     Db, err := sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
//     if err != nil {
//         log.Println(err)
//     }
//     defer Db.Close()

//     cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
//     _, err = Db.Exec(cmd, "Nancy", 20)

//     if err != nil {
//         log.Fatalln(err)
//     }
// }

// package main

// import (
//     "fmt"

//     "gopkg.in/go-ini/ini.v1"
// )

// type ConfigLite struct {
//     Port      int
//     DbName    string
//     SQLDriver string
// }

// var Config ConfigLite

// func init() {
//     cfg, _ := ini.Load("config.ini")

//     Config = ConfigLite{
//         //MustIntは数値型を読み込む 引数は初期値
//         Port: cfg.Section("web").Key("port").MustInt(8080),

//         DbName: cfg.Section("db").Key("name").MustString("example.sql"),

//         SQLDriver: cfg.Section("db").Key("driver").String(),
//     }
// }

// func main() {
//     fmt.Println("Port = %v\n", Config.Port)
//     fmt.Println("DbName = %v\n", Config.DbName, Config.DbName)
//     fmt.Println("SQLDriver = %v\n", Config.SQLDriver)
// }

package main

import (
    "fmt"

    "github.com/google/uuid"
)

func main() {
    uuidObj, _ := uuid.NewUUID()
    fmt.Println(" ", uuidObj.String())

    uuidObj2, _ := uuid.NewRandom()
    fmt.Println(" ", uuidObj2.String())
}