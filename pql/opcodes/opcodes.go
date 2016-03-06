package pseudoql

type OPCODE struct {
    OpCode uint8 // Hex value representing the code type
    SizeSz uint8 // Number of bytes the size value is. 1 = 1 byte (maxValue: 255), 2 = 2 bytes (maxValue: 65535), exc...
    ValueSz uint32 // If SizeSz is 0, this value represents a constant value of number of bytes for the value
    StackTerminator uint8 // Hex of opcode representing the terminator to jump out of stack. If there is a value here will enter a new stack frame until this opcode is reached.
    TailSeperatorRequired bool // If true the frame/stack's allowed seperators will be required next.
    AllowedGroups opcode_group
    newFn func() Opcoder
}

type Opcoder interface {
    GetOpCodeStruct() OPCODE

    GetSQLValue(*PQL) string
    GetPQLValue(*PQL) string

    GetValue(*PQL) string
    SetValue(string)
}
const (
    no_new_stack            uint8 = 0x0 // This is a special value stating that the opcode does not create a new stack
    // SEPERATORS
    and_code                uint8 = 0x1
    or_code                 uint8 = 0x2
    seperator_code          uint8 = 0x3

    // VALUES
    constant_code           uint8 = 0x4
    null_code               uint8 = 0x5
    variable_code           uint8 = 0x6

    table_ref_code          uint8 = 0x7
    field_code              uint8 = 0x8

    function_begin_code     uint8 = 0x9
    function_end_code       uint8 = 0xA
    function_param_begin_code uint8 = 0xB
    function_param_end_code uint8 = 0xC

    group_begin_code        uint8 = 0xD
    group_end_code          uint8 = 0xE

    // COMPARITORS
    equal_code              uint8 = 0xF
    greater_than_code       uint8 = 0x10
    less_than_code          uint8 = 0x11
    like_code               uint8 = 0x12
    no_comp_code            uint8 = 0x13
    not_equal_code          uint8 = 0x14
    not_like_code           uint8 = 0x15

    // If you add custom values here they must be >= 0x80 (128)
)
func NewOpType(code uint8) (Opcoder, bool){
    if s, ok := FindStructFromCode(code); ok {
        return s.newFn(), true
    }
    return nil, false
}

var codeToStructMap map[uint8]OPCODE

func init() {
    codeToStructMap = map[uint8]OPCODE{
        and_code:           AND_OC,
        or_code:            OR_OC,
        seperator_code:     SEPERATOR_OC,
        constant_code:      CONSTANT_OC,
        null_code:          NULL_OC,
        variable_code:      VARIABLE_OC,
        table_ref_code:     TABLE_REF_OC,
        field_code:         FIELD_OC,
        function_begin_code:        FUNCTION_BEGIN_OC,
        function_end_code:          FUNCTION_END_OC,
        function_param_begin_code:    FUNCTION_PARAM_BEGIN_OC,
        function_param_end_code:    FUNCTION_PARAM_END_OC,

        group_begin_code:   GROUP_BEGIN_OC,
        group_end_code:     GROUP_END_OC,
        equal_code:         EQUAL_OC,
        greater_than_code:  GREATER_THAN_OC,
        less_than_code:     LESS_THAN_OC,
        like_code:          LIKE_OC,
        no_comp_code:       NO_COMP_OC,
        not_equal_code:     NOT_EQUAL_OC,
        not_like_code:      NOT_LIKE_OC,
    }
}

func FindStructFromCode (code byte) (OPCODE, bool) {
    return codeToStructMap[code.(uint8)]
}