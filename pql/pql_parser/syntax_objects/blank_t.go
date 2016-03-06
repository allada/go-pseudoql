package pql_parser

var BLANK_OC = token{
    //FrameType: nil,
    //FrameTerminator: nil,
    Parse: func (curCodes []Opcoder, data []rune, start uint32) chan pqlParserReturnValue {
        cn := make(chan pqlParserReturnValue)
        go func () {
            defer close(cn)
            cn <- pqlParserReturnValue{
                BLANK_T{},
                start,
                true,
            }
        }
        return cn
    },
}

type BLANK_T struct{}