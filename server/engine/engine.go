package engine

import (
    "fmt"
    "net"
    "strings"
    "strconv"
)

var (
    //port defalut listen port
    host = "127.0.0.1"

    //port defalut listen port
    port = "12358"

    //db defalut database
    db = "mydb0"

    //dbMap defalut map 
    dbMap map[string]string

    //myStat system status
    myStat STAT
)

func Process(conn net.Conn) {
    defer conn.Close()
    for {
        buf := make([]byte, 512)
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read err:", err,conn)
            return
        }

        //fmt.Println(string(buf[0:n]))

	//send msg to client
	//msg := string(buf[0:n]) + " ok from server"
	//_, err = conn.Write([]byte(msg))
	//if err != nil {
	//    return
	//}
	msg := runCommand(string(buf[0:n]))
	_, err = conn.Write([]byte(msg))
	if err != nil {
	    return
	}
    }
}

func runCommand(commandStr string) string{
	msg:="not implement"
        commandStr = strings.TrimSuffix(commandStr, "\n")
        arrCommandStr := strings.Fields(commandStr)
        switch arrCommandStr[0] {
        case "status":
                fmt.Println(commandStr,"success")
                return Status()
        case "keys":
                fmt.Println(commandStr,"success")
                return keys()
        case "count":
                fmt.Println(commandStr,"success")
                return Count()
        case "get":
                //fmt.Println(commandStr,"success")
                return get(arrCommandStr[1])
        case "set":
                //fmt.Println(commandStr,"success")
		return set(arrCommandStr[1],arrCommandStr[2])
        case "del":
                //fmt.Println(commandStr,"success")
                return msg
	}
        return msg
}

func set(key, value string)string{
    	myStat.SetNum+=1
	dbMap[key] = value
	return key+" set ok"
}

func get(key string) string{
    myStat.GetNum+=1
    //check if exist
    v, ok := dbMap[key]
    if (ok) {
        //fmt.Println("ok", v)
	return v
    } else {
        //fmt.Println("not exist")
	return key+" not exist"
    }
}

func keys()string{
    msg:=""
    i:=0
    for key:= range dbMap{
        i++
        fmt.Println(key, "=", dbMap[key])
	msg+=key+"="+dbMap[key]+"\n"
	if i == 10 {
		break
 	}
    }
    return msg
}

func Count()string{
    cnt:=len(dbMap)
    myStat.Count=cnt
    return strconv.Itoa(cnt)
}

func Status()string{
    stat:="[status]\n"
    stat+="Count:"+strconv.Itoa(len(dbMap))+"\n"
    stat+="GetNum:"+strconv.Itoa(myStat.GetNum)+"\n"
    stat+="SetNum:"+strconv.Itoa(myStat.SetNum)+"\n"
    stat+="DelNum:"+strconv.Itoa(myStat.DelNum)+"\n"
    return stat
}

func init(){
	dbMap = make(map[string]string)
	myStat =STAT{
		Count:0,
		GetNum:0,
		SetNum:0,
		DelNum:0,
        }
}

func Test() {
    fmt.Println("start server...")
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
        go Process(conn)
    }
}

