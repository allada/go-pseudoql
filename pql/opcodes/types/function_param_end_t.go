package pseudoql

import "strings"

var FUNCTION_PARAM_END_OC = OPCODE{
    OpCode: function_param_end_code,
    SizeSz: 0,
    ValueSz: 2,
    StackTerminator: no_new_stack,
    AllowedGroups: frame_codes_g,
    newFn: func () Opcoder {
        return new(FUNCTION_PARAM_END_T)
    }
}

type FUNCTION_PARAM_END_T struct {

}

func (_ FUNCTION_PARAM_END_T) GetOpCodeStruct() OPCODE {
    return FUNCTION_PARAM_END_OC
}

func (c FUNCTION_PARAM_END_T) GetSQLValue(pql *PQL) string {
    return ","
}

func (c FUNCTION_PARAM_END_T) GetPQLValue(pql *PQL) string {
    return ","
}

func (c FUNCTION_PARAM_END_T) GetValue(_ *PQL) string {
    return ""
}

func (c *FUNCTION_PARAM_END_T) SetValue(v string) {

}