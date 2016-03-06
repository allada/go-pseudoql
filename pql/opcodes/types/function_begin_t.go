package pseudoql

import "strings"

var FUNCTION_BEGIN_OC = OPCODE{
    OpCode: function_begin_code,
    SizeSz: 0,
    ValueSz: 2,
    StackTerminator: function_end_code,
    AllowedGroups: function_param_g,
    newFn: func () Opcoder {
        return new(FUNCTION_BEGIN_T)
    }
}

type FUNCTION_BEGIN_T struct {
    function_ref uint16
}

func (_ FUNCTION_BEGIN_T) GetOpCodeStruct() OPCODE {
    return FUNCTION_BEGIN_OC
}

func (c FUNCTION_BEGIN_T) GetSQLValue(pql *PQL) string {
    return MakeFunctionRef(c.function_ref, pql)
}

func (c FUNCTION_BEGIN_T) GetPQLValue(pql *PQL) string {
    var buff bytes.Buffer

    buff.WriteString(pql.GetFunctionName(c.function_ref))
    buff.WriteString('(')

    return buff.String()
}

func (c FUNCTION_BEGIN_T) GetValue(_ *PQL) string {
    return c.function_ref
}

func (c *FUNCTION_BEGIN_T) SetValue(v string) {
    c.function_ref = v
}