package pql_parser

var NOT_LIKE_OC = token{
    FrameType: value_frame,
    FrameTerminator: nil,
    New: func () Opcoder {
        return new(NOT_LIKE_T)
    }
}

type NOT_LIKE_T struct{}