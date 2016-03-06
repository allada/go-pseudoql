package pql_parser

var function_param_frame = frame{
    allowedObjs: []OPCODE{
        FIELD_OC,
        FUNCTION_BEGIN_OC,
        GROUP_BEGIN_OC,
    },
    allowedSeperators: []OPCODE{
        // This may seem strange, but it's because the comma is consumed by the function_param terminator.
        // This will just jump back into the allowedObjs as a new param, if it cannot use any of them it jumps out of the frame.
        BLANK_OC,
    },
    checkValue: func (data []byte, ret pqlParserReturnValue) (pqlParserReturnValue, bool) {
        
    }
}