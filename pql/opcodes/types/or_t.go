package pseudoql

var OR_OC = OPCODE{
    OpCode: or_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    TailSeperatorRequired: false,
    AllowedGroups: frame_codes_g,
    newFn: func () Opcoder {
        return new(OR_T)
    }
}

type OR_T struct{

}

func (_ OR_T) GetOpCodeStruct() OPCODE {
    return OR_OC
}

func (_ OR_T) GetSQLValue(_ *PQL) string {
    return "OR"
}

func (_ OR_T) GetPQLValue(_ *PQL) string {
    return "|"
}

func (_ OR_T) GetValue(_ *PQL) string {
    return ""
}

func (_ OR_T) SetValue(_ string) {

}