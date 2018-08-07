package clib

import (
    "fmt"
    "strings"
    "errors"
)

// App is a main struct of clib package.
type App struct {
    // Application name
    Name string
    // Application version
    Version string
    // Argument
    Args []string
    // all Commands
    Commands map[string]*Command
    // all Options
    Options map[string]*Option
}

// NewApp is a constructor of App struct
func NewApp(name, version string) (*App, error) {
    
    app := &App {
        Name: name,
        Version: version,
        Commands: map[string]*Command{},
        Options: map[string]*Option{},
    }
    
    if err := app.AddOption("h", "Display the usage message", 0); err != nil {
        return nil, err
    }
    
    if err := app.AddOption("v", "Display the my version", 0); err != nil {
        return nil, err
    }
    
    return app, nil
    
}

// GetCommandArgs is a function to get arguments of given command name
func (a App) GetCommandArgs(name string) []string {
    if a.Commands[name] != nil {
        return a.Commands[name].GetArgs()
    }
    return []string{}
}

// GetOptionArgs is a function to get arguments of given option name
func (a App) GetOptionArgs(name string) []string {
    if a.Options[name] != nil {
        return a.Options[name].GetArgs()
    }
    return []string{}
}

// AddOption is a function to add given option to App
func (a *App) AddOption(name, synopsis string, argCount int) error {
    
    if a.Options[name] != nil {
        return fmt.Errorf("-%s option duplicated.\n", name)
    }
    
    opt, err := NewOption(name, synopsis, argCount)
    if err != nil {
        return err
    }
    
    a.Options[name] = opt
    return nil
    
}

// AddCommand is a function to add given command to App
func (a *App) AddCommand(name, shortName, synopsis string, argCount int) error {

    if a.Commands[name] != nil {
        return fmt.Errorf("%s command is duplicated", name)
    }

    cmd, err := NewCommand(name, shortName, synopsis, argCount)
    if err != nil {
        return err
    }

    a.Commands[name] = cmd
    return nil

}

// Help is a function to display help message
func (a App) Help() {
    fmt.Printf("Usage: \n\t")
    fmt.Printf("%s", a.Name)
    if len(a.Options) != 0 {
        fmt.Printf(" [option]")
        if a.hasOptionArg() {
            fmt.Println(" [<args>]")
        } else {
            fmt.Println()
        }
    }
    
    if len(a.Commands) != 0 {
        fmt.Printf("\t%s <command>", a.Name)
        if a.hasCommandArg() {
            fmt.Println(" [<args>]")
        } else {
            fmt.Println()
        }
    }
    fmt.Printf("\nOptions:\n\n")
    for _, o := range a.Options {
        fmt.Printf("\t-%s ", o.Name)
        if o.ArgCount > 1 {
            fmt.Printf("%s ...\t", o.ArgName)
        } else if o.ArgCount == 1 {
            fmt.Printf("%s\t\t", o.ArgName)
        } else {
            fmt.Printf("\t\t")
        }
        fmt.Printf("%s\n", o.Synopsis)
    }

    if len(a.Commands) != 0 {
        fmt.Printf("\nCommands:\n\n")
    }
    
    for _, c := range a.Commands {
        fmt.Printf("\t%s\t", c.Name)
        if c.ArgCount > 1 {
            fmt.Printf("%s ...\t", c.ArgName)            
        } else if c.ArgCount == 1 {
            fmt.Printf("%s\t", c.ArgName)            
        } else {
            fmt.Printf("\t")
        }
        fmt.Printf("%s\n", c.Synopsis)
    }
}

func (a App) hasOptionArg() bool {
    for _, o := range a.Options {
        if o.ArgCount != 0 {
            return true
        }
    }
    return false
}


func (a App) hasCommandArg() bool {
    for _, c := range a.Commands {
        if c.ArgCount != 0 {
            return true
        }
    }
    return false
}

// Parse is a function to parse the argument
func (a *App) Parse(args []string) (int, error){
    
    args_len := uint(len(args))
    
    if args_len == 0 {
        a.Help()
        return 0, nil
    }

    if args[0] == "-h" {
        a.Help()
        return 0, nil
    }

    if args[0] == "-v" {
        fmt.Printf("%s %s\n", a.Name, a.Version)
        return 0, nil
    }

    var i uint
    for i = 0; i < args_len; i++ {
        
        if strings.HasPrefix(args[i], "-") {
            
            if len(args[i]) != 2 {
                return 1, fmt.Errorf("Invalid option format: %v", args[i])
            }
            
            o := string(args[i][1])
            if a.Options[o] != nil {
                if s, err := a.Options[o].Parse(args[i:], &i); err != nil {
                    return s, err
                }
            }
        } else if a.Commands[args[i]] != nil {
            if s, err := a.Commands[args[i]].Parse(args[i:], &i); err != nil {
                return s, err
            }
        } else {
            a.Args = append(a.Args, args[i])
        }
        
    }
    
    return 0, nil

}

// hasComand is a function to exist the Option
func (a App) hasOption(s string) bool {
    for _, v := range a.Options {
        if v.Name == s {
            return true
        }
    }
    return false
}

// hasComand is a function to exist the Command
func (a App) hasCommand(s string) bool {
    for _, v := range a.Commands {
        if v.Name == s || v.ShortName == s {
            return true
        }
    }
    return false
}

// indexOfOption is a function to get the index of Option
func (a App) indexOfOption(s string) (uint) {
    
    var i uint
    for _, o := range a.Options {
        if o.Name == s {
            break
        }
        i += 1
    }
    return i
}

// indexOfOption is a function to get the index of Command
func (a App) indexOfComand(s string) (uint) {
    var i uint
    for _, c := range a.Commands {
        if c.Name == s || c.ShortName == s {
            break
        }
        i += 1
    }
    return i
}
