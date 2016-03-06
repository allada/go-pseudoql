package pseudoql

import "strings"

var GROUP_BEGIN_OC = OPCODE{
    OpCode: group_begin_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: group_end_code,
    AllowedGroups: general_g,
    newFn: func () Opcoder {
        return new(GROUP_BEGIN_T)
    }
}

type GROUP_BEGIN_T struct {

}

func (_ GROUP_BEGIN_T) GetOpCodeStruct() OPCODE {
    return GROUP_BEGIN_OC
}

func (c GROUP_BEGIN_T) GetSQLValue(pql *PQL) string {
    return '('
}

func (c GROUP_BEGIN_T) GetPQLValue(pql *PQL) string {
    return '('
}

func (c GROUP_BEGIN_T) GetValue(_ *PQL) string {
    return ""
}

func (c *GROUP_BEGIN_T) SetValue(v string) {

}