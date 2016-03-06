package pql_parser

var OR_OC = token{
    FrameType: nil,
    FrameTerminator: nil,
    New: func () Opcoder {
        return new(OR_T)
    }
}

type OR_T struct{}