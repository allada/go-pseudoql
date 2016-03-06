package pql_parser

var FUNCTION_END_OC = token{
    //FrameType: nil,
    //FrameTerminator: nil,
    Parse: func (curCodes []Opcoder, data []byte, start uint32) chan pqlParserReturnValue {
        cn := make(chan pqlParserReturnValue)
        go func () {
            defer close(cn)
            c, sz := getRuneAt(data, start)
            if c == ')' {
                cn <-pqlParserReturnValue{
                    FUNCTION_END_T{},
                    start + sz,
                    true,
                }
            } else {
                cn <-pqlParserReturnValue{
                    nil,
                    start,
                    false,
                }
            }
        }()
        return cn
    }
}

type FUNCTION_END_T struct{}