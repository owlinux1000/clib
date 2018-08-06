package clib

type Option struct {
    Name string
    Help string
    Flag bool
    Args []string
    ArgCount int
    ArgName string
}

func (o *Option) Parse(args []string, i *uint) int {
    
    o.Flag = true
    
    if o.ArgCount == 0 {
        *i -= 1
        return 0
    }
    
    if o.ArgCount > len(args) {
        return 1
    }
    
    for _i := 0; _i < o.ArgCount; _i++ {
        o.Args = append(o.Args, args[_i])
        if o.ArgCount != 1 {
            *i += 1
        }
    }
    
    return 0
    
}
