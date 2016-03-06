package pql_parser

var QUERY_OC = token{
    FrameType: general_frame,
    FrameTerminator: nil,
    New: func () Opcoder {
        return new(QUERY_T)
    }
}

type QUERY_T struct{}
