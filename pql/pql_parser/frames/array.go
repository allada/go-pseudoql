package pql_parser

var array_frame = frame{
    allowedObjs: []OPCODE{
        CONSTANT_OC,
        VARIABLE_OC,
    },
    allowedSeperators: []OPCODE{
        COMMA_OC,
    },
}