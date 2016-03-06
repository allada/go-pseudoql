package pseudoql

import "strings"

var EQUAL_OC = OPCODE{
    OpCode: equal_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: value_g,
    newFn: func () Opcoder {
        return new(EQUAL_T)
    }
}

type EQUAL_T struct {

}

func (_ EQUAL_T) GetOpCodeStruct() OPCODE {
    return EQUAL_OC
}

func (c EQUAL_T) GetSQLValue(pql *PQL) string {
    return '='
}

func (c EQUAL_T) GetPQLValue(pql *PQL) string {
    return ':'
}

func (c EQUAL_T) GetValue(_ *PQL) string {
    return ""
}

func (c *EQUAL_T) SetValue(v string) {

}