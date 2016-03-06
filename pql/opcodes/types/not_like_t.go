package pseudoql

import "strings"

var NOT_LIKE_OC = OPCODE{
    OpCode: not_like_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: value_g,
    newFn: func () Opcoder {
        return new(NOT_LIKE_T)
    }
}

type NOT_LIKE_T struct {

}

func (_ NOT_LIKE_T) GetOpCodeStruct() OPCODE {
    return NOT_LIKE_OC
}

func (c NOT_LIKE_T) GetSQLValue(pql *PQL) string {
    return "NOT LIKE"
}

func (c NOT_LIKE_T) GetPQLValue(pql *PQL) string {
    return "!~"
}

func (c NOT_LIKE_T) GetValue(_ *PQL) string {
    return ""
}

func (c *NOT_LIKE_T) SetValue(v string) {

}