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
        return 0
    }
    
    if c.ArgCount > len(args) - 1 {
        return 1
    }
    
    for _i := 1; _i < c.ArgCount + 1; _i++ {
        c.Args = append(c.Args, args[_i])
        *i++
    }
    
    return 0
    
}
