package pseudoql

import "strings"

var LESS_THAN_OC = OPCODE{
    OpCode: less_than_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    TailSeperatorRequired: false,
    AllowedGroups: value_g,
    newFn: func () Opcoder {
        return new(LESS_THAN_T)
    }
}

type LESS_THAN_T struct {

}

func (_ LESS_THAN_T) GetOpCodeStruct() OPCODE {
    return LESS_THAN_OC
}

func (c LESS_THAN_T) GetSQLValue(pql *PQL) string {
    return '<'
}

func (c LESS_THAN_T) GetPQLValue(pql *PQL) string {
    return '<'
}

func (c LESS_THAN_T) GetValue(_ *PQL) string {
    return ""
}

func (c *LESS_THAN_T) SetValue(v string) {

}