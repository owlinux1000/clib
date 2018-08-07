package clib

import "errors"

type Command struct {
    Name string
    ShortName string
    Synopsis string
    ArgName string
    ArgCount int
    SetFlag bool
    args []string
}

// NewCommand is a constructor of Command struct
func NewCommand(name, shortName, synopsis string, argCount int) (*Command, error) {

    return &Command{
        Name: name,
        ShortName: shortName,
        Synopsis: synopsis,
        ArgCount: argCount,
    }, nil

}

func (c Command) GetArgs() []string {
    return c.args
}

func (c *Command) Parse(args []string, i *uint) (int, error) {
    
    c.SetFlag = true
    
    if c.ArgCount == 0 {
        return 0, nil
    }
    
    if c.ArgCount > len(args) - 1 {
        return 1, errors.New("Argument count is invalid")
    }
    
    for _i := 1; _i < c.ArgCount + 1; _i++ {
        c.args = append(c.args, args[_i])
        *i++
    }

    if len(c.args) != c.ArgCount {
        return 1, errors.New("Argument count is invalid")
    } else {
        return 0, nil
    }

}
