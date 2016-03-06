package pql_parser

var compare_frame = frame{
    allowedObjs: []OPCODE{
        EQUAL_OC,
        GREATER_THAN_OC,
        LESS_THAN_OC,
        LIKE_OC,
        NOT_EQUAL_OC,
        NOT_LIKE_OC,
        
        NO_COMP_OC,
    },
    allowedSeperators: nil,
}
