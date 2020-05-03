package main

import (
    "fmt"
    "net"

    "github.com/yezi/mydb/server/engine"
)

var (
    //port defalut listen port
    host = "127.0.0.1"

    //port defalut listen port
    port = "12358"
)

func main() {
    fmt.Println("start server...listen:",port)
    listen, err := net.Listen("tcp", host+":"+port)
    if err != nil {
        fmt.Println("listen failed, err:", err)
        return
    }
    for {
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println("accept failed, err:", err)
            continue
        }
        go engine.Process(conn)
    }
}
