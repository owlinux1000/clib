package clib

import "errors"

// Command represents a subcommand
type Command struct {
    // Command name
    Name string
    // Command short name
    ShortName string
    // To be used Usage function
    Synopsis string
    // To be used argument name
    ArgName string
    // Expected argument count
    ArgCount int
    // setFlag represents whether this command was set
    setFlag bool
    // args represents argument of myself
    args []string
}

// NewCommand is a constructor of Command
func NewCommand(name, shortName, synopsis string, argCount int) (*Command, error) {

    return &Command{
        Name: name,
        ShortName: shortName,
        Synopsis: synopsis,
        ArgCount: argCount,
    }, nil

}

// Args returns args of myself
func (c Command) Args() []string {
    return c.args
}

// SetFlag returns setFlag of myself
func (c Command) SetFlag() bool {
    return c.setFlag
}

// Parse is a function to parse the argument
func (c *Command) Parse(args []string, i *uint) (int, error) {
    
    c.setFlag = true
    
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
