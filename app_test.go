package clib

import "testing"

func TestAppParseWithNoArgCommand(t *testing.T) {

    app, _ := NewApp("testapp", "1.0.0")
    app.AddCommand("install", "i", "Install command", 0)
    args := []string{
        "testapp",
        "install",
        "hoge",
    }
    
    actual, _ := app.Parse(args[1:])
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

    app, _ := NewApp("testapp", "1.0.0")
    app.AddCommand("install", "i", "Install command", 1)
    args := []string{    
        "testapp",
        "install",
        "hoge",
    }
    
    actual, _ := app.Parse(args[1:])
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

    app, _ := NewApp("testapp", "1.0.0")
    app.AddCommand("install", "i", "Install command", 2)
    app.AddOption("a", "a option", 1)
    
    args := []string{
        "testapp",
        "install",
        "hoge",
        "fuga",
        "-a",
        "huge",
    }
    
    actual, _ := app.Parse(args[1:])
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
    actualArgs := app.GetOptionArgs("a")
    if actualArgs[0] != expectedArg {
        t.Errorf("got: %v\nwant: %v", actualArgs[0], expectedArg)
    }
    
}

