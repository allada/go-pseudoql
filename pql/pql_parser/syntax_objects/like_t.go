package pql_parser

var LIKE_OC = token{
    FrameType: value_frame,
    FrameTerminator: nil,
    New: func () Opcoder {
        return new(LIKE_T)
    }
}

type LIKE_T struct{}