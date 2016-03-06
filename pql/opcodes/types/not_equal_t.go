package pseudoql

import "strings"

var NOT_EQUAL_OC = OPCODE{
    OpCode: not_equal_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: value_g,
    newFn: func () Opcoder {
        return new(NOT_EQUAL_T)
    }
}

type NOT_EQUAL_T struct {

}

func (_ NOT_EQUAL_T) GetOpCodeStruct() OPCODE {
    return NOT_EQUAL_OC
}

func (c NOT_EQUAL_T) GetSQLValue(pql *PQL) string {
    return "!="
}

func (c NOT_EQUAL_T) GetPQLValue(pql *PQL) string {
    return "!:"
}

func (c NOT_EQUAL_T) GetValue(_ *PQL) string {
    return ""
}

func (c *NOT_EQUAL_T) SetValue(v string) {

}