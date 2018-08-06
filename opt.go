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
        return 0
    }
    
    if o.ArgCount > len(args) - 1 {
        return 1
    }

    for _i := 1; _i < o.ArgCount + 1; _i++ {
        o.Args = append(o.Args, args[_i])
        *i++
    }

    if len(o.Args) != o.ArgCount {
        return 1
    } else {
        return 0        
    }
    
}
