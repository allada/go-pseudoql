package pql_parser

var ARRAY_BEGIN_OC = token{
    FrameType: array_frame,
    FrameTerminator: ARRAY_END_OC,
    Parse: func (curCodes []Opcoder, data []rune, start uint32) chan pqlParserReturnValue {
        cn := make(chan pqlParserReturnValue)
        go func () {
            defer close(cn)
            c, sz := getRuneAt(data, start)
            if c == '[' {
                frameChan := ARRAY_BEGIN_OC.FrameType.Parse(data, start + sz)
                
                cn <-pqlParserReturnValue{
                    ARRAY_BEGIN_T{},
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

type ARRAY_BEGIN_T struct{}