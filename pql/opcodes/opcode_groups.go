package pseudoql

type opcode_group struct {
    types []OPCODE
    allowedSeperators opcode_group
}
// SPECIAL value representing the frame's allowed seperators.
var frame_seperator_g = opcode_group{
    []OPCODE{},
    nil,
}
// SPECIAL value representing the frame's allowed codes
var frame_codes_g = opcode_group{
    []OPCODE{},
    nil,
}

var gate_no_sep_g = opcode_group{
    []OPCODE{
        AND_OC,
        OR_OC,
    },
    nil,
}
var gate_with_sep_g = opcode_group{
    []OPCODE{
        AND_OC,
        OR_OC,
        SEPERATOR_OC,
    },
    nil,
}
var function_param_sep_g = opcode_group{
    []OPCODE{
        SEPERATOR_OC,
    },
    nil,
}
var function_param_g = opcode_group{
    []OPCODE{
        CONSTANT_OC,
        NULL_OC,
        VARIABLE_OC,
        FIELD_OC,
        FUNCTION_BEGIN_OC,
        GROUP_BEGIN_OC,
    },
    function_param_sep_g,
}

var condition_in_group_g = opcode_group{
    []OPCODE{
        FIELD_OC,
        FUNCTION_BEGIN_OC,
        GROUP_BEGIN_OC,
    },
    nil,
}
var ref_g = opcode_group{
    []OPCODE{
        FIELD_OC,
        FUNCTION_BEGIN_OC,
        GROUP_BEGIN_OC,
    },
    nil,
}
var value_g = opcode_group{
    []OPCODE{
        CONSTANT_OC,
        NULL_OC,
        VARIABLE_OC,
    },
    nil,
}
var comparator_g = opcode_group{
    []OPCODE{
        EQUAL_OC,
        GREATER_THAN_OC,
        LESS_THAN_OC,
        LIKE_OC,
        NOT_EQUAL_OC,
        NOT_LIKE_OC,
        NO_COMP_OC,
    },
    nil,
}

var general_g = opcode_group{
    []OPCODE{
        FIELD_OC,
        FUNCTION_BEGIN_OC,
        GROUP_BEGIN_OC,
    },
    seperator_no_param_g,
}
var general_with_sep_g = opcode_group{
    []OPCODE{
        FIELD_OC,
        FUNCTION_BEGIN_OC,
        GROUP_BEGIN_OC,
    },
    seperator_param_g,
}