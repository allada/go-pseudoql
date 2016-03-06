package pql_parser

var general_frame = frame{
    allowedObjs: []OPCODE{
        FIELD_OC,
        FUNCTION_BEGIN_OC,
        GROUP_BEGIN_OC,
    },
    allowedSeperators: []OPCODE{
        AND_OC,
        OR_OC,
    },
}
