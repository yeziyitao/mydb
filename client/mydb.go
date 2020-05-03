package client

import (
    "fmt"
    "net"
    "os"

    "bufio"
    "strings"
    "errors"
    "strconv"
    //"os/exec"
)

var (
    //Name program name
    Name = "mydb v0.0.2"

    //port defalut listen port
    host = "127.0.0.1"

    //port defalut listen port
    port = "12358"
)

func firstPrint(){
        fmt.Println("Welcome to the "+Name)
        fmt.Println(`Type 'help;' or '\h' for help.`)
}

func help(){
        fmt.Println("List of all "+Name+" commands:")
        fmt.Println(`
exit      () Exit mydb.
help      (\h) Display this help.
status    () Display server status.
keys      () List 10 keys.
set       () set key value.
get       () get key.
`)
}

func Test() {
        firstPrint()
        reader := bufio.NewReader(os.Stdin)
        for {
                fmt.Print(Name+"> ")
                cmdString, err := reader.ReadString('\n')
                if err != nil {
                        fmt.Fprintln(os.Stderr, err)
                }
                err = runCommand(cmdString)
                if err != nil {
                        fmt.Fprintln(os.Stderr, err)
                }
        }
}

func runCommand(commandStr string) error {
        commandStr = strings.TrimSuffix(commandStr, "\n")
        arrCommandStr := strings.Fields(commandStr)
        switch arrCommandStr[0] {
        case "help":
                help()
                return nil
        case "\\h":
                help()
                return nil
        case "exit":
                os.Exit(0)
        case "get":
                fmt.Println(commandStr,"success")
                return nil
        case "set":
                fmt.Println(commandStr,"success")
                return nil
        case "del":
                fmt.Println(commandStr,"success")
                return nil
        case "plus":
                // Not using `sum` because it's a registered command in unix
                if len(arrCommandStr) < 3 {
                        return errors.New("Required for 2 arguments")
                }
                arrNum := []int64{}
                for i, arg := range arrCommandStr {
                        if i == 0 {
                                continue
                        }
                        n, _ := strconv.ParseInt(arg, 10, 64)
                        arrNum = append(arrNum, n)
                }
                fmt.Fprintln(os.Stdout, sum(arrNum...))
                return nil
                // add another case here for custom commands.
        }
        //cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
        //cmd.Stderr = os.Stderr
        //cmd.Stdout = os.Stdout
        //return cmd.Run()
	return nil
}

func sum(numbers ...int64) int64 {
        res := int64(0)
        for _, num := range numbers {
                res += num
        }
        return res
}

func Conn() {

    firstPrint()

    conn, err := net.Dial("tcp", host+":"+port)
    if err != nil {
        fmt.Println("Error dialing", err.Error())
        return
    }
    fmt.Println("connected")

    defer conn.Close()
    inputReader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print(Name+"> ")
        cmdString, err := inputReader.ReadString('\n')
        if err != nil {
                fmt.Fprintln(os.Stderr, err)
        }
        err = runCommand(cmdString)
        if err != nil {
                fmt.Fprintln(os.Stderr, err)
        }

        //send cmd to server
        trimmedInput := strings.Trim(cmdString, "\r\n")
        _, err = conn.Write([]byte(trimmedInput))
        if err != nil {
            return
        }
	//reveive msg from server
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
	    fmt.Println("read err:", err,conn)
	    return
	}
	
	fmt.Println(string(buf[0:n]))
    }
}

func ExeConnOnce(cmdString string) {
    conn, err := net.Dial("tcp", host+":"+port)
    if err != nil {
        fmt.Println("Error dialing", err.Error())
        return
    }
    defer conn.Close()

    //send cmd to server
    trimmedInput := strings.Trim(cmdString, "\r\n")
    _, err = conn.Write([]byte(trimmedInput))
    if err != nil {
        return
    }
    //reveive msg from server
    buf := make([]byte, 512)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("read err:", err,conn)
        return
    }
    
    fmt.Println(string(buf[0:n]))
}
