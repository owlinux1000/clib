package clib

import "errors"

type Option struct {
    Name string
    Synopsis string
    ArgName string
    ArgCount int
    SetFlag bool
    args []string
}

func NewOption(name, synopsis string, argCount int) (*Option, error) {

    if len(name) != 1 {
        return nil, errors.New("Option name must one letter.")
    }
    
    return &Option{
        Name: name,
        Synopsis: synopsis,
        ArgCount: argCount,
    }, nil
    
}

func (o Option) GetArgs() []string {
    return o.args
}

func (o *Option) Parse(args []string, i *uint) (int, error) {
    
    o.SetFlag = true
    
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
