package pseudoql

type seperator interface {
    parse(interface{}, string)
}

type opcode interface {
    getSeperators() *[]seperator
}