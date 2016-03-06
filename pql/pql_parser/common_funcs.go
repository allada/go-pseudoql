package pql_parser

import (
    "unicode/utf8"
)

func getRuneAt(d []byte, i uint32) (rune, int) {
    return utf8.DecodeRune(d[i:])
}