package pseudoql

import "strings"

var NO_COMP_OC = OPCODE{
    OpCode: no_comp_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: frame_seperator_g,
    newFn: func () Opcoder {
        return new(NO_COMP_T)
    }
}

type NO_COMP_T struct {

}

func (_ NO_COMP_T) GetOpCodeStruct() OPCODE {
    return NO_COMP_OC
}

func (c NO_COMP_T) GetSQLValue(pql *PQL) string {
    return ""
}

func (c NO_COMP_T) GetPQLValue(pql *PQL) string {
    return ""
}

func (c NO_COMP_T) GetValue(_ *PQL) string {
    return ""
}

func (c *NO_COMP_T) SetValue(v string) {

}