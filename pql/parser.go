package pseudoql

import "reflect"

type PQL struct {
    Query string
    Table string
    Group string
    OrderBy SortObj
    Select SelectObj
    Variables VariablesObj
}

type SortObj map[string]string
type SelectObj map[string]string
type VariablesObj map[string]string

func (p *PQL) ToSQL() string {
    
}

func (p *PQL) ToOpCodes() *Opcode {
    
}

func (p *PQL) t_general(pos uint) (newPos uint, opcode *Opcode) {
    
}

type frame struct {
    allowedObjs []OPCODE
    allowedSeperators []OPCODE
    checkValue func ([]byte, pqlParserReturnValue) (pqlParserReturnValue, bool)
}
func (f frame) Parse (data []byte, start uint32) pqlParserReturnValue {
    mainCn := make(chan pqlParserReturnValue)
    cnt := len(f.allowedObjs)
    counter := 0
    
    for i, code := range f.allowedObjs {
        cn := code.Parse(data, start)
        channels[i] = cn
        go func () {
            c := <-cn
            counter++
            if c.success {
                mainCn <-c
            } else if counter >= cnt {
                close(mainC)
            }
        }()
    }

    // Each Parse() function will return a channel which will return a pqlParserReturnValue. Each function
    // will be ran as a thread. The first one to be returned will be the one it will use. In the event
    // that more than one of these functions returns a success, the one that returns first will be used
    // and the others will be ignored. It is very important (to keep unexpected issues from happening) that
    // only one function returns success per allowedObjs array. After the return value is received it's
    // return value is allowed to be modified or fail by passing the info the the frame's checkValue()
    // function.
    ok := false
    for !ok {
        select {
            case v = <-mainCn:
                // Only success pqlParserReturnValues should even be seen here
                if f.checkValue {
                    v, ok = f.checkValue(data, v)
                    if ok {
                        close(mainC)
                        return v
                    }
                } else {
                    close(mainC)
                    return v
                }
            default:
                close(mainC)
                return pqlParserReturnValue{
                    success: false,
                }
        }
    }

    //f.allowedObjs: []OPCODE
    //f.allowedSeperators: []OPCODE
}
type pqlParserReturnValue struct {
    code Opcoder
    size uint32
    success bool
}