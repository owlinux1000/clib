package clib

import (
    "fmt"
    "sort"
    "strings"
)

// App is a main struct of clib package.
type App struct {
    // Application name
    Name string
    // Application version
    Version string
    // Synopsis
    Synopsis string
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

func (a *App) addSynopsis(i interface{}) {
    switch i.(type) {
    case Option:

    case Command:

    }
}

// Help is a function to display help message
func (a *App) Help() (s string) {

    s = "Usage: \n\t" + a.Name
    if len(a.Options) > 0 {
        s += " [option]"
        if have, count := a.haveOptionArg(); have {
            if count > 1 {
                s += " [<args>...]\n"
            } else if count == 1 {
                s += " [<args>]\n"
            }
        } else {
            s += "\n"
        }
    }

    if len(a.Commands) != 0 {
        s += "\t" + a.Name
        s += " <command>"
        if have, count := a.haveCommandArg(); have {
            f := a.haveCommandNoArg()
            if count > 1 && f {
                s += " [<args>...]\n"
            } else if count == 1 && f {
                s += " [<args>]\n"
            } else if count == 1 {
                s += " <args>\n"
            }
        } else {
            s += "\n"
        }
    }
    
    keys := make([]string, 0, len(a.Options))
    for k := range a.Options {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    s += "\nOptions:\n"
    for _, k := range keys {
        s += "\t-" + k
        if a.Options[k].ArgCount > 1 {
            s += a.Options[k].ArgName + " ...\t"
        } else if a.Options[k].ArgCount == 1 {
            s += a.Options[k].ArgName + "\t\t"
        } else {
            s += "\t\t"
        }

        s += a.Options[k].Synopsis + "\n"

    }

    if len(a.Commands) == 0 {
        return s
    }

    s += "\nCommands:\n"
    keys = make([]string, 0, len(a.Commands))
    for k := range a.Commands {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    
    for _, k := range keys {
        s += "\t" + k + "\t"
        if a.Commands[k].ArgCount > 1 {
            s += a.Commands[k].ArgName + " ...\t"
        } else if a.Commands[k].ArgCount == 1 {
            s += a.Commands[k].ArgName + "\t"
        } else {
            s += "\t"
        }
        s += a.Commands[k].Synopsis + "\n"
    }

    return s

}

func (a App) haveOptionArg() (bool, int) {
    f := false
    max_argc := 0
    for _, o := range a.Options {
        if o.ArgCount != 0 {
            f = true
            if o.ArgCount > max_argc {
                max_argc = o.ArgCount
            }
        }
    }
    return f, max_argc
}

func (a App) haveCommandArg() (bool, int) {
    f := false
    max_argc := 0
    for _, c := range a.Commands {
        if c.ArgCount != 0 {
            f = true
            if c.ArgCount > max_argc {
                max_argc = c.ArgCount
            }
        }
    }
    return f, max_argc
}

func (a App) haveCommandNoArg() bool {
    for _, c := range a.Commands {
        if c.ArgCount == 0 {
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
