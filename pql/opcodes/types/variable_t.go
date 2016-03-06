package pseudoql

import "bytes"

var VARIABLE_OC = OPCODE{
    OpCode: variable_code,
    SizeSz: 4,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: frame_seperator_g,
    newFn: func () Opcoder {
        return new(VARIABLE_T)
    }
}

type VARIABLE_T struct {
    name string
}

func (_ VARIABLE_T) GetOpCodeStruct() OPCODE {
    return VARIABLE_OC
}

func (c VARIABLE_T) GetSQLValue(pql *PQL) string {
    var buff bytes.Buffer

    buff.WriteByte('\'')
    buff.WriteString(EscapeSQL(pql.GetVariableValue(c.name)))
    buff.WriteByte('\'')

    return buff.String()
}

func (c VARIABLE_T) GetPQLValue(pql *PQL) string {
    var buff bytes.Buffer

    buff.WriteByte('\'')
    buff.WriteString(EscapePQL(pql.GetVariableValue(c.name)))
    buff.WriteByte('\'')

    return buff.String()
}

func (c VARIABLE_T) GetValue(_ *PQL) string {
    return c.name
}

func (c *VARIABLE_T) SetValue(v string) {
    c.name = v
}