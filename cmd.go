package clib

type Command struct {
    Name string
    ShortName string
    Synopsis string
    ArgName string
    ArgCount int
    SetFlag bool
    args []string
}

func (c Command) GetArgs() []string {
    return c.args
}

func NewCommand(name, shortName, synopsis string, argCount int) *Command {
    return &Command{
        Name: name,
        ShortName: shortName,
        Synopsis: synopsis,
        ArgCount: argCount,
    }
}

func (c *Command) Parse(args []string, i *uint) int {
    
    c.SetFlag = true
    
    if c.ArgCount == 0 {
        return 0
    }
    
    if c.ArgCount > len(args) - 1 {
        return 1
    }
    
    for _i := 1; _i < c.ArgCount + 1; _i++ {
        c.args = append(c.args, args[_i])
        *i++
    }

    if len(c.args) != c.ArgCount {
        return 1
    } else {
        return 0
    }
}
