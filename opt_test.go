package clib

import "testing"

func TestParseNoArgOption(t *testing.T) {
    
    opt := Option{
        Name: "a",
        ArgCount: 0,
        Help: "a option",
    }
    
    var i uint = 0
    args := []string{
        "-a",
        "hoge",
    }
    
    actual := opt.Parse(args, &i)

    if actual != 0 {
        t.Errorf("got: %v\nwant: %v", actual, 0)
    }

    if len(opt.Args) != 0 {
        t.Errorf("got: %v\nwant: %v", len(opt.Args), 0)
    }

    if i != 0 {
        t.Errorf("got: %v\nwant: %v", i, 0)
    }
    
}

func TestParseOneArgOption(t *testing.T) {
    
    opt := Option{
        Name: "a",
        ArgCount: 1,
        Help: "a option",
    }
    
    var i uint = 0
    args := []string{
        "-a",
        "hoge",
    }
    
    actual := opt.Parse(args, &i)

    if actual != 0 {
        t.Errorf("got: %v\nwant: %v", actual, 0)
    }

    if len(opt.Args) != 1 {
        t.Errorf("got: %v\nwant: %v", len(opt.Args), 1)
    }

    if i != 1 {
        t.Errorf("got: %v\nwant: %v", i, 1)
    }
    
}

func TestParseOneArgOption2(t *testing.T) {
    
    opt := Option{
        Name: "a",
        ArgCount: 1,
        Help: "a option",
    }
    
    var i uint = 0

    args := []string{
        "-a", 
    }
    
    // args doesn't have argument of a option. So, opt.Parse returns 1
    actual := opt.Parse(args, &i)
    if actual != 1 {
        t.Errorf("got: %v\nwant: %v", actual, 1)
    }
    
}

func TestParseTwoArgOption(t *testing.T) {
    
    opt := Option{
        Name: "a",
        ArgCount: 2,
        Help: "a option",
    }
    
    var i uint = 0
    args := []string{
        "-a",
        "hoge",
        "fuga",
    }
    
    actual := opt.Parse(args, &i)
    
    if actual != 0 {
        t.Errorf("got: %v\nwant: %v", actual, 0)
    }

    if len(opt.Args) != 2 {
        t.Errorf("got: %v\nwant: %v", len(opt.Args), 2)
    }

    if i != 2 {
        t.Errorf("got: %v\nwant: %v", i, 2)
    }
    
}
