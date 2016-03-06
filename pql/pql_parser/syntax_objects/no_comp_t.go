package pql_parser

var NO_COMP_OC = token{
    FrameType: nil,
    FrameTerminator: nil,
    New: func () Opcoder {
        return new(NO_COMP_T)
    }
}

type NO_COMP_T struct{}