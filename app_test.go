package clib

import "testing"

func TestAppParseWithNoArgCommand(t *testing.T) {

    app := NewApp("testapp", "1.0.0")
    app.AddCommand("install", "i", "Install command", 0)
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
}

func TestAppParseWithOneArgCommand(t *testing.T) {

    app := NewApp("testapp", "1.0.0")
    app.AddCommand("install", "i", "Install command", 1)
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
    actualArg := app.GetCommandArgs("install")[0]
    if actualArg != expectedArg {
        t.Errorf("got: %v\nwant: %v", actualArg, expectedArg)
    }
}

func TestAppParseOneCommandOneOption(t *testing.T) {

    app := NewApp("testapp", "1.0.0")
    app.AddCommand("install", "i", "Install command", 2)
    
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

    expectedCommandArgs := []string{"hoge", "fuga"}
    actualCommandArgs := app.GetCommandArgs("install")

    for i, _ := range expectedCommandArgs {
        if actualCommandArgs[i] != expectedCommandArgs[i] {
            t.Errorf("got: %v\nwant: %v", actualCommandArgs[i], expectedCommandArgs[i])
        }
    }
    
    expectedArg := "huge"
    if app.Options[2].Args[0] != expectedArg {
        t.Errorf("got: %v\nwant: %v", app.Options[2].Args[0], expectedArg)
    }
}

