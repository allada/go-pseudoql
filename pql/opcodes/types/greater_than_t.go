package pseudoql

import "strings"

var GREATER_THAN_OC = OPCODE{
    OpCode: greater_than_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: value_g,
    newFn: func () Opcoder {
        return new(GREATER_THAN_T)
    }
}

type GREATER_THAN_T struct {

}

func (_ GREATER_THAN_T) GetOpCodeStruct() OPCODE {
    return GREATER_THAN_OC
}

func (c GREATER_THAN_T) GetSQLValue(pql *PQL) string {
    return '>'
}

func (c GREATER_THAN_T) GetPQLValue(pql *PQL) string {
    return '>'
}

func (c GREATER_THAN_T) GetValue(_ *PQL) string {
    return ""
}

func (c *GREATER_THAN_T) SetValue(v string) {

}