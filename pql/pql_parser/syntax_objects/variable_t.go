package pql_parser

var VARIABLE_OC = token{
    FrameType: nil,
    FrameTerminator: nil,
    New: func () Opcoder {
        return new(VARIABLE_T)
    }
}

type VARIABLE_T struct{}