package clib

type Command struct {
    Name string
    ShortName string
    ArgCount int
    Help string
    Flag bool
    ArgName string
    Args []string
}

func (c *Command) Parse(args []string, i *uint) int {
    
    c.Flag = true
    
    if c.ArgCount == 0 {
        *i -= 1
        return 0
    }

    if c.ArgCount > len(args) {
        return 1
    }

    for _i := 0; _i < c.ArgCount; _i++ {
        
        c.Args = append(c.Args, args[_i])
        
        if c.ArgCount != 1 {
            *i += 1
        }
    }

    return 0
    
}
