package pseudoql

import "strings"

var FIELD_OC = OPCODE{
    OpCode: field_code,
    SizeSz: 0,
    ValueSz: 2,
    StackTerminator: no_new_stack,
    AllowedGroups: comparator_g,
    newFn: func () Opcoder {
        return new(FIELD_T)
    }
}

type FIELD_T struct {
    field_ref uint16
}

func (_ FIELD_T) GetOpCodeStruct() OPCODE {
    return FIELD_OC
}

func (c FIELD_T) GetSQLValue(pql *PQL) string {
    return MakeFieldRef(c.field_ref, pql)
}

func (c FIELD_T) GetPQLValue(pql *PQL) string {
    table := pql.GetTableOfFieldRef(c.field_ref)
    table_refs := pql.GetTableRefs(table)
    field_name := pql.GetFieldName(c.field_ref)

    append(table_refs, field_name)
    return strings.Join(refs, '.')
}

func (c FIELD_T) GetValue(_ *PQL) string {
    return c.field_ref
}

func (c *FIELD_T) SetValue(v string) {
    c.field_ref = v
}