package pql_parser

var AND_OC = token{
    //FrameType: nil,
    //FrameTerminator: nil,
    Parse: func (curCodes []Opcoder, data []byte, start uint32) <-chan pqlParserReturnValue {
        cn := make(chan pqlParserReturnValue)
        go func () {
            defer close(cn)
            found := false
            curPos := start
            for c, sz := getRuneAt(data, curPos);
                c == ' ' || c == '\t' || c == '\r' || c == '\n';
                c, sz := getRuneAt(data, curPos) {
                    found = true
                    curPos += sz
            }
            if found {
                cn <-pqlParserReturnValue{
                    AND_T{},
                    curPos,
                    true
                }
            } else {
                cn <-pqlParserReturnValue{
                    nil,
                    start,
                    false
                }
            }
        }
        return cn
    },
}

type AND_T struct{}