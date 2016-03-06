package main

import "fmt"

type OR_T struct{}
func (_ OR_T) GetOpCodeStruct() OPCODE {
    return OR_OC
}
type Opcoder interface {
    GetOpCodeStruct() OPCODE
}

type OPCODE struct {
    OpCode uint8
    SizeSz uint8
    ValueSz uint32 // If SizeByteSize is 0, this value is a second check for size of value
    newFn func() Opcoder
}

var OR_OC OPCODE

type A struct{}

func (_ A) D(_ string) {
    fmt.Println(OR_OC.newFn())
}

func main () {
    OR_OC = OPCODE{
        OpCode: 0x1,
        SizeSz: 0,
        ValueSz: 0,
        newFn: func () Opcoder {
            return new(OR_T)
        },
    }
}