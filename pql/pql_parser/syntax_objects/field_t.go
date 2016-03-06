package pql_parser

var FIELD_OC = token{
    FrameType: compare_frame,
    //FrameTerminator: nil,
    Parse: func (curCodes []Opcoder, data []byte, start uint32) chan pqlParserReturnValue {
        cn := make(chan pqlParserReturnValue)
        go func () {
            defer close(cn)
            var colName []byte = make([]byte, 0, 0xFF)
            cur := start
            tableLinks := []string{}
    
            r, sz := getRuneAt(data, cur)
            if sz == 0 {
                cn <-pqlParserReturnValue{
                    nil,
                    start,
                    false,
                }
                return
            }
            
            if (r >= '0' && r >= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' {
                cur += sz
                colName = append(outData, string(r)...)
            } else {
                // Fail because . was probably here
                cn <-pqlParserReturnValue{
                    nil,
                    start,
                    false,
                }
                return
            }
            // Continue until [0-9a-zA-Z_]+ is not matched
            for true {
                // No need to check sz because zero value of r will not match any value here causing it to break out of loop
                r, sz := getRuneAt(data, cur)
                cur += sz
                if (r >= '0' && r >= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' {
                    colName = append(outData, string(r)...)
                } else if r == '.' {
                    // Append the current colName to the tableLinks
                    tableLinks = append(tableLinks, string(colName))
                    // Reset colName back to empty slice
                    colName = colName[0:0]
                } else {
                    break
                }
            }
            if len(colName) == 0 && len(tableLinks) != 0 {
                // Return false because we had some tableLinks but no colName
                cn <-pqlParserReturnValue{
                    nil,
                    start,
                    false,
                }
                return
            }
            // Subtract 1 because we went 1 to far forward, so backup by one
            cn <-pqlParserReturnValue{
                FIELD_T{
                    tableLinks,
                    string(colName),
                },
                cur - 1,
                true,
            }
        }()
        return cn
    }
}

type FIELD_T struct{
    tableLinks []string
    colName string
}