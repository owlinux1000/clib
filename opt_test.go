package clib

import "testing"

func TestParseNoArgOption(t *testing.T) {
    
    opt, _ := NewOption("a", "a option", 0)
    
    var i uint = 0
    args := []string{
        "-a",
        "hoge",
    }

    expected := 0
    actual, _ := opt.Parse(args, &i)
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
    
    if len(opt.GetArgs()) != 0 {
        t.Errorf("got: %v\nwant: %v", len(opt.GetArgs()), 0)
    }

    if i != 0 {
        t.Errorf("got: %v\nwant: %v", i, 0)
    }
    
}

func TestParseOneArgOption(t *testing.T) {

    opt, _ := NewOption("a", "a option", 1)
    
    var i uint = 0
    args := []string{
        "-a",
        "hoge",
    }

    expected := 0
    actual, _ := opt.Parse(args, &i)
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
    
    if len(opt.GetArgs()) != 1 {
        t.Errorf("got: %v\nwant: %v", len(opt.GetArgs()), 1)
    }

    if i != 1 {
        t.Errorf("got: %v\nwant: %v", i, 1)
    }
    
}

func TestParseOneArgOption2(t *testing.T) {
    
    opt, _ := NewOption("a", "a option", 1)
    
    var i uint = 0
    args := []string{
        "-a", 
    }
    
    // args doesn't have argument of a option. So, opt.Parse returns 1
    expected := 1
    actual, _ := opt.Parse(args, &i)
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestParseTwoArgOption(t *testing.T) {

    opt, _ := NewOption("a", "a option", 2)
    
    var i uint = 0
    args := []string{
        "-a",
        "hoge",
        "fuga",
    }

    expected := 0
    actual, _ := opt.Parse(args, &i)
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
    
    if len(opt.GetArgs()) != 2 {
        t.Errorf("got: %v\nwant: %v", len(opt.GetArgs()), 2)
    }

    if i != 2 {
        t.Errorf("got: %v\nwant: %v", i, 2)
    }
    
}
