package pseudoql

import "strings"

// Can never be parsed from string
var TABLE_REF_OC = OPCODE{
    OpCode: table_ref_code,
    SizeSz: 0,
    ValueSz: 2,
    newFn: func () Opcoder {
        return new(TABLE_REF_T)
    }
}

type TABLE_REF_T struct {
    table_code uint16
}

func (_ TABLE_REF_T) GetOpCodeStruct() OPCODE {
    return TABLE_REF_OC
}

func (c TABLE_REF_T) GetSQLValue(pql *PQL) string {
    return MakeTableRef(c.table_code, pql)
}

func (c TABLE_REF_T) GetPQLValue(pql *PQL) string {
    refs := pql.GetTableRefs(c.table_code)
    
    return strings.Join(refs, '.')
}

func (c TABLE_REF_T) GetValue(_ *PQL) string {
    return c.table_code
}

func (c *TABLE_REF_T) SetValue(v string) {
    c.table_code = v
}