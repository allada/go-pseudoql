package pseudoql

var NULL_OC = OPCODE{
    OpCode: null_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: frame_seperator_g,
    newFn: func () Opcoder {
        return new(NULL_T)
    }
}

type NULL_T struct{}

func (_ NULL_T) GetOpCodeStruct() OPCODE {
    return NULL_OC
}

func (_ NULL_T) GetSQLValue(_ *PQL) string {
    return "NULL"
}

func (_ NULL_T) GetPQLValue(_ *PQL) string {
    return "-"
}

func (_ NULL_T) GetValue(_ *PQL) string {
    return ""
}

func (_ NULL_T) SetValue(_ string) {

}