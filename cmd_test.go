package clib

import (
    "testing"
)

func TestParseNoArgCommand(t *testing.T) {
    
    cmd := NewCommand("install", "i", "Install command", 0)
    
    var i uint = 0
    args := []string{
        "install",
        "hoge",
    }
    
    actual := cmd.Parse(args, &i)
    
    if actual != 0 {
        t.Errorf("got: %v\nwant: %v", actual, 0)
    }

    if len(cmd.GetArgs()) != 0 {
        t.Errorf("got: %v\nwant: %v", len(cmd.GetArgs()), 0)
    }

    if i != 0 {
        t.Errorf("got: %v\nwant: %v", len(cmd.GetArgs()), 0)
    }
    
}

func TestParseOneArgCommand(t *testing.T) {
    
    cmd := Command{
        Name: "install",
        ShortName: "i",
        ArgCount: 1,
        Synopsis: "Install Command",
    }
    
    var i uint = 0
    args := []string{
        "install",
        "hoge",
        "fuga",
    }
    
    actual := cmd.Parse(args, &i)
    if actual != 0 {
        t.Errorf("got: %v\nwant: %v", actual, 0)
    }
    
    if len(cmd.GetArgs()) != 1 {
        t.Errorf("got: %v\nwant: %v", len(cmd.GetArgs()), 1)
    }
    
    if cmd.GetArgs()[0] != "hoge" {
        t.Errorf("got: %v\nwant: %v", cmd.GetArgs(), []string{"hoge"})
    }
    
    if i != 1 {
        t.Errorf("got: %v\nwant: %v", i, 1)
    }
    
}

func TestParseOneArgCommand2(t *testing.T) {
    
    cmd := Command{
        Name: "install",
        ShortName: "i",
        ArgCount: 1,
        Synopsis: "Install Command",
    }
    
    var i uint = 0
    args := []string{
        "install",
    }

    // args doesn't have argument of install command. So, cmd.Parse returns 1
    actual := cmd.Parse(args, &i)
    if actual != 1 {
        t.Errorf("got: %v\nwant: %v", actual, 1)
    }
    
}

func TestParseTwoArgCommand(t *testing.T) {
    
    cmd := Command{
        Name: "install",
        ShortName: "i",
        ArgCount: 2,
        Synopsis: "Install Command",
    }
    
    var i uint = 0
    args := []string{
        "install", 
        "hoge",
        "fuga",
    }
    
    actual := cmd.Parse(args, &i)
    
    if actual != 0 {
        t.Errorf("got: %v\nwant: %v", actual, 0)
    }
    
    if cmd.GetArgs()[0] != "hoge" || cmd.GetArgs()[1] != "fuga" {
        t.Errorf("got: %v\nwant: %v", cmd.GetArgs(), []string{"hoge", "fuga"})
    }
    
    if len(cmd.GetArgs()) != 2 {
        t.Errorf("got: %v\nwant: %v", len(cmd.GetArgs()), 2)
    }

    if i != 2 {
        t.Errorf("got: %v\nwant: %v", i, 2)
    }
}
