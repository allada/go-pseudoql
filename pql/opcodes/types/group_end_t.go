package pseudoql

import "strings"

var GROUP_END_OC = OPCODE{
    OpCode: function_end_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: frame_codes_g,
    newFn: func () Opcoder {
        return new(GROUP_BEGIN_T)
    }
}

type GROUP_END_T struct {

}

func (_ GROUP_END_T) GetOpCodeStruct() OPCODE {
    return GROUP_END_OC
}

func (c GROUP_END_T) GetSQLValue(pql *PQL) string {
    return ')'
}

func (c GROUP_END_T) GetPQLValue(pql *PQL) string {
    return ')'
}

func (c GROUP_END_T) GetValue(_ *PQL) string {
    return ""
}

func (c *GROUP_END_T) SetValue(v string) {

}