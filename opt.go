package clib

import "errors"

// Option represents a option
type Option struct {
    // Option name such as -h
    Name string
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

// NewOption is a constructor of Option
func NewOption(name, synopsis string, argCount int) (*Option, error) {

    if len(name) != 2  || name[0] != '-' {
        return nil, errors.New("Option name must following format: -v")
    }
    
    return &Option{
        Name: name,
        Synopsis: synopsis,
        ArgCount: argCount,
    }, nil
    
}

// Args returns args of myself
func (o Option) Args() []string {
    return o.args
}

// SetFlag returns setFlag of myself
func (o Option) SetFlag() bool {
    return o.setFlag
}

// Parse is a function to parse the argument
func (o *Option) Parse(args []string, i *uint) (int, error) {
    
    o.setFlag = true
    
    if o.ArgCount == 0 {
        return 0, nil
    }
    
    if o.ArgCount > len(args) - 1 {
        return 1, errors.New("Argument count is invalid")
    }

    for _i := 1; _i < o.ArgCount + 1; _i++ {
        o.args = append(o.args, args[_i])
        *i++
    }

    if len(o.args) != o.ArgCount {
        return 1, errors.New("Argument count is invalid")
    } else {
        return 0, nil
    }
    
}
