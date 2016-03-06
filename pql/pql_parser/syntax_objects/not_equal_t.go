package pql_parser

var NOT_EQUAL_OC = token{
    FrameType: value_frame,
    FrameTerminator: nil,
    New: func () Opcoder {
        return new(NOT_EQUAL_T)
    }
}

type NOT_EQUAL_T struct{}