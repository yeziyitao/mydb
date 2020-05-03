package main

import (
        "encoding/json"
        "flag"
        "fmt"
        "os"
        "time"

	"github.com/yezi/mydb/client"
)

var (
        //VERSION tool version
        VERSION = "1.1.0"

        //h help
        h bool
        //v,V version
        v, V bool

        //e execute
        e string
)

/*
struct MSG
markdown msg
*/
type MSG struct {
        // msg type
        Msgtype string `json:"msgtype"`

        // markdown msg content
        Info Markdown `json:"markdown"`
}

/*
struct Markdown
markdown content
*/
type Markdown struct {
        // content detail info ,markdown type
        Content string `json:"content"`
}

/*
@ConvertJson
Convert msg to json format with markdown
send msg to wx
*/
func ConvertJson(content string) (jsonStr []byte) {
        mymsg := MSG{
                Msgtype: "json",
                Info: Markdown{
                        Content: content,
                },
        }
        // convert json to []byte
        myjsonStr, _ := json.Marshal(mymsg)
        return myjsonStr
}

/*
@usage
edit help msg
*/
func usage() {
        fmt.Println(os.Args[0], `version:`, VERSION, `
Usage: `, os.Args[0], `[OPTIONS]

Options:
        `)
        flag.PrintDefaults()
}

/*
@init
init flag
*/
func init() {
        // set help and version
        flag.BoolVar(&h, "h", false, "this help")
        flag.BoolVar(&v, "v", false, "show version and exit")
        flag.BoolVar(&V, "V", false, "show version and configure options then exit")

        // wx msg webhook addr
        flag.StringVar(&e, "e", "", "")

        //set defalut Usage to usage
        flag.Usage = usage

}

/*
@main
*/
func main() {
        flag.Parse()
        if h || v || V {
                flag.Usage()
                os.Exit(0)
        }
        if len(e) > 0 {

                var cstZone = time.FixedZone("CST", 8*3600) // 8 timezone
                timestamp := time.Now().In(cstZone).Format("2006-01-02 15:04:05")

                info := timestamp + "=" + e
                fmt.Println(e, string(ConvertJson(info)))
                client.ExeConnOnce(e)
        } else {
                client.Conn()
        }

}
