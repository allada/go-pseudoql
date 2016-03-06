package pql_parser

type Opcoder interface {
    Parse(string, uint32) (Opcoder, uint32, bool)
    GetOPCodes() []byte
}

type frame struct {
    AllowedObjs         []token
    AllowedSeperators   []token
}

type token struct {
    FrameType       frame
    FrameTerminator frame
    Parse           func([]rune, uint32) chan pqlParserReturnValue
}