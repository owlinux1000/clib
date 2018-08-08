package clib

import "testing"

func TestAppParseWithNoArgCommand(t *testing.T) {

    app, _ := NewApp("testapp", "1.0.0", "")
    app.AddCommand("install", "i", "Install command", 0)
    args := []string{
        "testapp",
        "install",
        "hoge",
    }
    
    actual, _ := app.Parse(args)
    expected := 0
    
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }

    expectedArg := "hoge"
    if app.Args()[0] != expectedArg {
        t.Errorf("got: %v\nwant: %v", app.Args()[0], expectedArg)
    }
    
}

func TestAppParseWithOneArgCommand(t *testing.T) {

    app, _ := NewApp("testapp", "1.0.0", "")
    app.AddCommand("install", "i", "Install command", 1)
    args := []string{    
        "testapp",
        "install",
        "hoge",
    }
    
    actual, _ := app.Parse(args)
    expected := 0
    
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }

    expectedArg := "hoge"
    actualArgs, _ := app.CommandArgs("install")
    if actualArgs[0] != expectedArg {
        t.Errorf("got: %v\nwant: %v", actualArgs[0], expectedArg)
    }
    
}

func TestAppParseOneCommandOneOption(t *testing.T) {

    app, _ := NewApp("testapp", "1.0.0", "")
    app.AddCommand("install", "i", "Install command", 2)
    app.AddOption("-a", "a option", 1)
    
    args := []string{
        "testapp",
        "install",
        "hoge",
        "fuga",
        "-a",
        "huge",
    }
    
    actual, _ := app.Parse(args)
    expected := 0
    
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }

    expectedCommandArgs := []string{"hoge", "fuga"}
    actualCommandArgs, _ := app.CommandArgs("install")
    
    for i, _ := range expectedCommandArgs {
        if actualCommandArgs[i] != expectedCommandArgs[i] {
            t.Errorf("got: %v\nwant: %v", actualCommandArgs[i], expectedCommandArgs[i])
        }
    }

    expectedArg := "huge"
    actualArgs, _ := app.OptionArgs("-a")
    if actualArgs[0] != expectedArg {
        t.Errorf("got: %v\nwant: %v", actualArgs[0], expectedArg)
    }
    
}

func TestHelp(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0", "")
    expected := "Usage: \n\ttestapp [option]\n"
    expected += "\nOptions:\n\t-h\t\tDisplay this message\n\t-v\t\tPrint version info and exit\n"
    actual := app.Usage()
    if actual != expected {
        t.Errorf("got: \n%v\nwant: \n%v", actual, expected)
    }
}

func TestHelp2(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0", "")
    expected := "Usage: \n\ttestapp [option] [<args>]\n"
    expected += "\nOptions:\n"
    expected += "\t-a\t\ta option\n"
    expected += "\t-h\t\tDisplay this message\n\t-v\t\tPrint version info and exit\n"
    app.AddOption("-a", "a option", 1)
    actual := app.Usage()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestHelp3(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0", "")
    expected := "Usage: \n\ttestapp [option] [<args>...]\n"
    expected += "\nOptions:\n"
    expected += "\t-a ...\ta option\n"
    expected += "\t-h\t\tDisplay this message\n\t-v\t\tPrint version info and exit\n"
    app.AddOption("-a", "a option", 2)
    actual := app.Usage()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestHelp4(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0", "")
    expected := "Usage: \n\ttestapp [option]\n\ttestapp <command>\n"
    expected += "\nOptions:\n"
    expected += "\t-h\t\tDisplay this message\n\t-v\t\tPrint version info and exit\n"
    expected += "\nCommands:\n"
    expected += "\tinstall\t\tInstall command\n"
    app.AddCommand("install", "i", "Install command", 0)
    actual := app.Usage()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestHelp5(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0", "")
    expected := "Usage: \n\ttestapp [option]\n\ttestapp <command> <args>\n"
    expected += "\nOptions:\n"
    expected += "\t-h\t\tDisplay this message\n\t-v\t\tPrint version info and exit\n"
    expected += "\nCommands:\n"
    expected += "\tinstall\t\tInstall command\n"
    app.AddCommand("install", "i", "Install command", 1)
    actual := app.Usage()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestHelp6(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0", "")
    expected := "Usage: \n\ttestapp [option]\n\ttestapp <command> [<args>]\n"
    expected += "\nOptions:\n"
    expected += "\t-h\t\tDisplay this message\n\t-v\t\tPrint version info and exit\n"
    expected += "\nCommands:\n"
    expected += "\tinstall\tFILE\tInstall command\n"
    expected += "\tinstall2\t\tInstall2 command\n"
    app.AddCommand("install", "i", "Install command", 1)
    app.Commands["install"].ArgName = "FILE"
    app.AddCommand("install2", "i2", "Install2 command", 0)
    actual := app.Usage()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

func TestHelp7(t *testing.T) {
    app, _ := NewApp("testapp", "1.0.0", "")
    expected := "Usage: \n\ttestapp [option]\n\ttestapp <command> [<args>...]\n"
    expected += "\nOptions:\n"
    expected += "\t-h\t\tDisplay this message\n\t-v\t\tPrint version info and exit\n"
    expected += "\nCommands:\n"
    expected += "\tinstall\tFILE ...\tInstall command\n"
    expected += "\tinstall2\t\tInstall2 command\n"
    app.AddCommand("install", "i", "Install command", 2)
    app.AddCommand("install2", "i2", "Install2 command", 0)
    app.Commands["install"].ArgName = "FILE"
    actual := app.Usage()
    if actual != expected {
        t.Errorf("got: %v\nwant: %v", actual, expected)
    }
}

