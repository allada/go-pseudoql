package pseudoql

// Operation Object
type opInterface interface {
    parse(interface{}, string)
}

var t_general_ops []opInterface

func init(){
    t_general_ops = []opInterface{
        t_table_ref,
        t_function,
        t_constant,
        t_field,
        t_group_opener,
    }
}