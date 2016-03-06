package pql_parser

var FUNCTION_PARAM_OC = token{
    FrameType: general_frame,
    FrameTerminator: COMMA_OC, // This will consume the comma making it unavailable for the param to know where to end easily.
    Parse: func (curCodes []Opcoder, data []byte, start uint32) chan pqlParserReturnValue {
        // THIS FUNCTION WILL NEVER BE EXECUTED! Instead a FUNCTION_PARAM_T is created on the checkValue() of
        // the function_param_frame.
        cn := make(chan pqlParserReturnValue)
        return cn
    }
}

type FUNCTION_PARAM_T struct{
    param OPCODE
}