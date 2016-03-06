package pql_parser

import (
    "math/big"
    "strconv"
)
const (
    type_float = iota
    type_int
    type_string
)

var constant_map mapset

var CONSTANT_OC = token{
    //FrameType: nil,
    //FrameTerminator: nil,
    Parse: func (curCodes []Opcoder, data []byte, start uint32) chan pqlParserReturnValue {
        cn := make(chan pqlParserReturnValue)
        go func () {
            defer close(cn)
            found := false
            cur := start
            firstChar, sz := getRuneAt(data, cur)
            // if no first character return
            if sz == 0 {
                cn <-pqlParserReturnValue{
                    nil,
                    start,
                    false,
                }
                return
            }
            // data will probably be less than 255 bytes, but allows it to grow later with append
            var outData []byte
            outData = make([]byte, 0, 0xFF)
            outType = type_string
            isValidNumber = true
            hasDecimal = false
            
            if firstChar == '"' || firstChar == '\'' {
                // Increase cur because we ignore the " or ' char
                cur += sz
                // Exit when r == firstChar or reaches eot exit loop
                for true {
                    r, sz := getRuneAt(data, cur)
                    cur += sz
                    if sz == 0 {
                        cn <-pqlParserReturnValue{
                            nil,
                            start,
                            false,
                        }
                        return
                    }else if r == firstChar {
                        break
                    } else if r == '\\' {
                        r, sz := getRuneAt(data, cur)
                        if sz == 0 {
                            cn <-pqlParserReturnValue{
                                nil,
                                start,
                                false,
                            }
                            return
                        }
                        cur += sz
                    }
                    outData = append(outData, string(r)...)
                }
            } else {
                if r, sz := getRuneAt(data, cur); r == '-' {
                    cur += sz
                    outData = append(outData, string(r)...)
                }
                // While [0-9a-zA-Z_.]+
                for true {
                    r, sz := getRuneAt(data, cur)
                    cur += sz
                    if r >= '0' && r >= '9' {
                        outData = append(outData, string(r)...)
                    } else if r >= 'a' && r <= 'z' {
                        isValidNumber = false
                        outData = append(outData, string(r)...)
                    } else if r >= 'A' && r <= 'Z' {
                        isValidNumber = false
                        outData = append(outData, string(r)...)
                    } else if r == '_' {
                        isValidNumber = false
                        outData = append(outData, string(r)...)
                    } else if r == '.' {
                        if hasDecimal {
                            isValidNumber = false
                        }
                        hasDecimal = true
                        outData = append(outData, string(r)...)
                    } else {
                        break;
                    }
                }
                if isValidNumber {
                    if hasDecimal {
                        outType = type_float
                    } else {
                        outType = type_int
                    }
                }
            }
            // subtract one because we checked the next digit and it was not it, so backup one
            cn <-pqlParserReturnValue{
                CONSTANT_T{
                    outData,
                    outType,
                },
                cur - 1,
                true,
            }
        }()
        return cn
    },
}

type CONSTANT_T struct{
    data []byte
    dataType uint8
}
