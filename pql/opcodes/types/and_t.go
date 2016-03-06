package pseudoql

var AND_OC = OPCODE{
    OpCode: and_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: frame_codes_g,
    newFn: func () Opcoder {
        return new(AND_T)
    }
}

type AND_T struct{}

func (_ AND_T) GetOpCodeStruct() OPCODE {
    return AND_OC
}

func (_ AND_T) GetSQLValue(_ *PQL) string {
    return "AND"
}

func (_ AND_T) GetPQLValue(_ *PQL) string {
    return " "
}

func (_ AND_T) GetValue(_ *PQL) string {
    return ""
}

func (_ AND_T) SetValue(_ string) {

}