package pql_parser

var FUNCTION_BEGIN_OC = token{
    FrameType: function_params_frame,
    FrameTerminator: FUNCTION_END_OC,
    Parse: func (curCodes []Opcoder, data []byte, start uint32) chan pqlParserReturnValue {
        cn := make(chan pqlParserReturnValue)
        go func () {
            defer close(cn)
            var fnName []byte = make([]byte, 0, 0xFF)
            cur := start
    
            r, sz := getRuneAt(data, cur)
            if sz == 0 {
                cn <-pqlParserReturnValue{
                    nil,
                    start,
                    false,
                }
                return
            }
            if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' {
                colName = append(outData, string(r)...)
                cur += sz
            } else {
                // Fail because might have started with 0-9
                cn <-pqlParserReturnValue{
                    nil,
                    start,
                    false,
                }
                return
            }
            // Continue until [0-9a-zA-Z_]+\( is not matched
            for true {
                // No need to check sz because zero value of r will not match any value here causing it to break out of loop
                r, sz := getRuneAt(data, cur)
                cur += sz
                if (r >= '0' && r >= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' {
                    fnName = append(outData, string(r)...)
                } else if r == '(' {
                    // No need to subtract one because we are returning immidiatly
                    cn <-pqlParserReturnValue{
                        FUNCTION_BEGIN_T{
                            string(fnName)
                        },
                        cur,
                        true,
                    }
                    return
                } else {
                    cn <-pqlParserReturnValue{
                        nil,
                        start,
                        false,
                    }
                    return
                }
            }
            // We can never reach this point
        }()
        return cn
    }
}

type FUNCTION_BEGIN_T struct{
    fnName string
}