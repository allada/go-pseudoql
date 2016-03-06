package pql_parser

var COMMA_OC = token{
    //FrameType: nil,
    //FrameTerminator: nil,
    Parse: func (curCodes []Opcoder, data []rune, start uint32) chan pqlParserReturnValue {
        cn := make(chan pqlParserReturnValue)
        go func () {
            defer close(cn)
            c, sz := getRuneAt(data, start)
            if c == ',' {
                cn <-pqlParserReturnValue{
                    COMMA_T{},
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
    },
}

type COMMA_T struct{}