package engine

/*
struct STAT
record db status
*/
type STAT struct {
        //Count record num
        Count int `json:"count"`

        //GetNum record read num,by get
        GetNum int `json:"get_num"`

        //SetNum record write num,by set
        SetNum int `json:"set_num"`

        //DelNum record del num,by del
        DelNum int `json:"del_num"`
}
