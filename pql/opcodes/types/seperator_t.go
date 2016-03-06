package pseudoql

var SEPERATOR_OC = OPCODE{
    OpCode: seperator_code,
    SizeSz: 0,
    ValueSz: 0,
    StackTerminator: no_new_stack,
    AllowedGroups: frame_codes_g,
    newFn: func () Opcoder {
        return new(SEPERATOR_T)
    }
}

type SEPERATOR_T struct{}

func (_ SEPERATOR_T) GetOpCodeStruct() OPCODE {
    return SEPERATOR_OC
}

func (_ SEPERATOR_T) GetSQLValue(_ *PQL) string {
    return ","
}

func (_ SEPERATOR_T) GetPQLValue(_ *PQL) string {
    return ","
}

func (_ SEPERATOR_T) GetValue(_ *PQL) string {
    return ""
}

func (_ SEPERATOR_T) SetValue(_ string) {

}