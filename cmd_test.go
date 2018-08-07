package clib

import (
    "testing"
)

func TestParseNoArgCommand(t *testing.T) {
    
    cmd, _ := NewCommand("install", "i", "Install command", 0)
    
    var i uint = 0
    args := []string{
        "install",
        "hoge",
    }

    expected := 0
    actual, _ := cmd.Parse(args, &i)
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }

    if len(cmd.Args()) != 0 {
        t.Errorf("got: %v\nwant: %v", len(cmd.Args()), 0)
    }

    if i != 0 {
        t.Errorf("got: %v\nwant: %v", len(cmd.Args()), 0)
    }
    
}

func TestParseOneArgCommand(t *testing.T) {

    cmd, _ := NewCommand("install", "i", "Install command", 1)

    var i uint = 0
    args := []string{
        "install",
        "hoge",
        "fuga",
    }

    expected := 0
    actual, _ := cmd.Parse(args, &i)
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
    
    if len(cmd.Args()) != 1 {
        t.Errorf("got: %v\nwant: %v", len(cmd.Args()), 1)
    }
    
    if cmd.Args()[0] != "hoge" {
        t.Errorf("got: %v\nwant: %v", cmd.Args(), []string{"hoge"})
    }
    
    if i != 1 {
        t.Errorf("got: %v\nwant: %v", i, 1)
    }
    
}

func TestParseOneArgCommand2(t *testing.T) {

    cmd, _ := NewCommand("install", "i", "Install command", 1)
    
    var i uint = 0
    args := []string{
        "install",
    }

    // args doesn't have argument of install command. So, cmd.Parse returns 1
    expected := 1
    actual, _ := cmd.Parse(args, &i)
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
    
}

func TestParseTwoArgCommand(t *testing.T) {
    
    cmd, _ := NewCommand("install", "i", "Install command", 2)
    
    var i uint = 0
    args := []string{
        "install", 
        "hoge",
        "fuga",
    }

    expected := 0
    actual, _ := cmd.Parse(args, &i)
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
    
    if cmd.Args()[0] != "hoge" || cmd.Args()[1] != "fuga" {
        t.Errorf("got: %v\nwant: %v", cmd.Args(), []string{"hoge", "fuga"})
    }
    
    if len(cmd.Args()) != 2 {
        t.Errorf("got: %v\nwant: %v", len(cmd.Args()), 2)
    }

    if i != 2 {
        t.Errorf("got: %v\nwant: %v", i, 2)
    }
}
