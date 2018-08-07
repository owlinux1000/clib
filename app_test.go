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

func TestHelp(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0")
    expected := "Usage: \n\ttestapp [option]\n"
    //expected += "\nOptions:\n\t-h\t\tDisplay the usage message\n\t-v\t\tDisplay the my version\n"
    actual := app.Help()
    if actual != expected {
        t.Errorf("got: \n%v\nwant: \n%v", actual, expected)
    }
}

func TestHelp2(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0")
    expected := "Usage: \n\ttestapp [option] [<args>]\n"
    app.AddOption("a", "a option", 1)
    actual := app.Help()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestHelp3(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0")
    expected := "Usage: \n\ttestapp [option] [<args>...]\n"
    app.AddOption("a", "a option", 2)
    actual := app.Help()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestHelp4(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0")
    expected := "Usage: \n\ttestapp [option]\n\ttestapp <command>\n"
    app.AddCommand("install", "i", "Install command", 0)
    actual := app.Help()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestHelp5(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0")
    expected := "Usage: \n\ttestapp [option]\n\ttestapp <command> <args>\n"
    app.AddCommand("install", "i", "Install command", 1)
    actual := app.Help()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestHelp6(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0")
    expected := "Usage: \n\ttestapp [option]\n\ttestapp <command> [<args>]\n"
    app.AddCommand("install", "i", "Install command", 1)
    app.AddCommand("install2", "i2", "Install2 command", 0)
    actual := app.Help()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestHelp7(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0")
    expected := "Usage: \n\ttestapp [option]\n\ttestapp <command> [<args>...]\n"
    app.AddCommand("install", "i", "Install command", 2)
    app.AddCommand("install2", "i2", "Install2 command", 0)
    actual := app.Help()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

