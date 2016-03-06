package pql_parser

var EQUAL_OC = token{
    FrameType: value_frame,
    //FrameTerminator: nil,
    Parse: func (curCodes []Opcoder, data []byte, start uint32) chan pqlParserReturnValue {
        cn := make(chan pqlParserReturnValue)
        go func () {
            defer close(cn)
            c, sz := getRuneAt(data, start)
            if c == ':' || c == '=' {
                cn <-pqlParserReturnValue{
                    EQUAL_T{},
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

type EQUAL_T struct{}