package pseudoql

import "strings"

var FUNCTION_PARAM_BEGIN_OC = OPCODE{
    OpCode: function_param_begin_code,
    SizeSz: 0,
    ValueSz: 2,
    StackTerminator: function_param_end_code,
    AllowedGroups: function_param_g,
    newFn: func () Opcoder {
        return new(FUNCTION_PARAM_BEGIN_T)
    }
}

type FUNCTION_PARAM_BEGIN_T struct {

}

func (_ FUNCTION_PARAM_BEGIN_T) GetOpCodeStruct() OPCODE {
    return FUNCTION_PARAM_BEGIN_OC
}

func (c FUNCTION_PARAM_BEGIN_T) GetSQLValue(pql *PQL) string {
    return ""
}

func (c FUNCTION_PARAM_BEGIN_T) GetPQLValue(pql *PQL) string {
    return ""
}

func (c FUNCTION_PARAM_BEGIN_T) GetValue(_ *PQL) string {
    return ""
}

func (c *FUNCTION_PARAM_BEGIN_T) SetValue(v string) {

}