package pql_parser

var GROUP_END_OC = token{
    FrameType: nil,
    FrameTerminator: nil,
    New: func () Opcoder {
        return new(GROUP_END_T)
    }
}

type GROUP_END_T struct{}