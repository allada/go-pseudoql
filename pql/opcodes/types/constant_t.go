package pseudoql

import "bytes"

var CONSTANT_OC = OPCODE{
    OpCode: constant_code,
    SizeSz: 4,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: frame_seperator_g,
    newFn: func () Opcoder {
        return new(CONSTANT_T)
    }
}

type CONSTANT_T struct {
    value string
}

func (_ CONSTANT_T) GetOpCodeStruct() OPCODE {
    return CONSTANT_OC
}

func (c CONSTANT_T) GetSQLValue(_ *PQL) string {
    var buff bytes.Buffer

    buff.WriteByte('\'')
    buff.WriteString(EscapeSQL(c.value))
    buff.WriteByte('\'')

    return buff.String()
}

func (c CONSTANT_T) GetPQLValue(_ *PQL) string {
    var buff bytes.Buffer

    buff.WriteByte('"')
    buff.WriteString(EscapePQL(c.value))
    buff.WriteByte('"')

    return buff.String()
}

func (c CONSTANT_T) GetValue(_ *PQL) string {
    return c.value
}

func (c *CONSTANT_T) SetValue(v string) {
    c.value = v
}