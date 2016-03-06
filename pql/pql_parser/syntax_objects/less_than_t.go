package pql_parser

var LESS_THAN_OC = token{
    FrameType: value_frame,
    FrameTerminator: nil,
    New: func () Opcoder {
        return new(LESS_THAN_T)
    }
}

type LESS_THAN_T struct{}