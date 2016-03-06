package pql_parser

var GROUP_BEGIN_OC = token{
    FrameType: GENERAL_OC,
    FrameTerminator: GROUP_END_OC,
    New: func () Opcoder {
        return new(GROUP_BEGIN_T)
    }
}

type GROUP_BEGIN_T struct{}