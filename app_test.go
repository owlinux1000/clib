package clib

import "testing"

func TestAppParseWithNoArgCommand(t *testing.T) {

    app := NewApp("testapp", "1.0.0")
    app.AddCommand(Command{
        Name: "install",
        ShortName: "i",
        ArgCount: 0,
        Help: "Install Command",
    })

    args := []string{
        "testapp",
        "install",
        "hoge",
    }
    
    actual := app.Parse(args[1:])
    expected := 0
    
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }

    expectedArg := "hoge"
    if app.Args[0] != expectedArg {
        t.Errorf("got: %v\nwant: %v", app.Args[0], expectedArg)
    }

    flagCmds := app.FlagCommands()
    if !flagCmds["install"] || !flagCmds["i"]{
        t.Errorf("got: true\nwant: %v", flagCmds["install"])
    }
    
}

func TestAppParseWithOneArgCommand(t *testing.T) {

    app := NewApp("testapp", "1.0.0")
    app.AddCommand(Command{
        Name: "install",
        ShortName: "i",
        ArgCount: 1,
        Help: "Install Command",
    })

    args := []string{
        "testapp",
        "install",
        "hoge",
    }
    
    actual := app.Parse(args[1:])
    expected := 0
    
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }

    expectedArg := "hoge"
    if app.Commands[0].Args[0] != expectedArg {
        t.Errorf("got: %v\nwant: %v", app.Commands[0].Args[0], expectedArg)
    }

    flagCmds := app.FlagCommands()
    if !flagCmds["install"] || !flagCmds["i"]{
        t.Errorf("got: true\nwant: %v", flagCmds["install"])
    }
    
}

func TestAppParseOneCommandOneOption(t *testing.T) {

    app := NewApp("testapp", "1.0.0")
    
    app.AddCommand(Command{
        Name: "install",
        ShortName: "i",
        ArgCount: 2,
        Help: "Install Command",
    })
    
    app.AddOption(Option{
        Name: "a",
        ArgCount: 1,
        Help: "a option",
    })

    args := []string{
        "testapp",
        "install",
        "hoge",
        "fuga",
        "-a",
        "huge",
    }
    
    actual := app.Parse(args[1:])
    expected := 0
    
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }

    expectedArg := "hoge"
    if app.Commands[0].Args[0] != expectedArg {
        t.Errorf("got: %v\nwant: %v", app.Commands[0].Args[0], expectedArg)
    }

    expectedArg = "fuga"
    if app.Commands[0].Args[1] != expectedArg {
        t.Errorf("got: %v\nwant: %v", app.Commands[0].Args[1], expectedArg)
    }
    
    expectedArg = "huge"
    if app.Options[2].Args[0] != expectedArg {
        t.Errorf("got: %v\nwant: %v", app.Options[2].Args[0], expectedArg)
    }

    flagCmds := app.FlagCommands()
    if !flagCmds["install"] || !flagCmds["i"]{
        t.Errorf("got: true\nwant: %v", flagCmds["install"])
    }

    optsCmds := app.FlagOptions()
    if !optsCmds["a"] {
        t.Errorf("got: true\nwant: %v", optsCmds["a"])
    }
    
}

