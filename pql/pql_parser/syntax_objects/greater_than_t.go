package pql_parser

var GREATER_THAN_OC = token{
    FrameType: value_frame,
    FrameTerminator: nil,
    New: func () Opcoder {
        return new(GREATER_THAN_T)
    }
}

type GREATER_THAN_T struct{}