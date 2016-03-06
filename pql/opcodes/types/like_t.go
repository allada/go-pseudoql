package pseudoql

import "strings"

var LIKE_OC = OPCODE{
    OpCode: like_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    TailSeperatorRequired: false,
    AllowedGroups: value_g,
    newFn: func () Opcoder {
        return new(LIKE_T)
    }
}

type LIKE_T struct {

}

func (_ LIKE_T) GetOpCodeStruct() OPCODE {
    return LIKE_OC
}

func (c LIKE_T) GetSQLValue(pql *PQL) string {
    return "LIKE"
}

func (c LIKE_T) GetPQLValue(pql *PQL) string {
    return '~'
}

func (c LIKE_T) GetValue(_ *PQL) string {
    return ""
}

func (c *LIKE_T) SetValue(v string) {

}