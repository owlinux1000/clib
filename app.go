package clib

import (
    "fmt"
    "strings"
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
    Commands []Command
    // all Options
    Options []Option
}

func NewApp(name, version string) *App {
    app := &App {
        Name: name,
        Version: version,
    }
    app.AddOption(Option{
        Name: "h",
        Help: "Display the usage message",
    })
    app.AddOption(Option{
        Name: "v",
        Help: "Display the my version",
    })
    return app
}

// AddOption is a function to add given option to App
func (a *App) AddOption(option Option) bool {
    if a.hasOption(option.Name) {
        fmt.Printf("-%s option duplicated.\n", option.Name)
        return false
    }
    a.Options = append(a.Options, option)
    return true
}

// AddCommand is a function to add given command to App
func (a *App) AddCommand(command Command) bool {
    if a.hasCommand(command.Name) {
        fmt.Printf("%s command duplicated.\n", command.Name)
        return false
    }
    a.Commands = append(a.Commands, command)
    return true
}

func (a App) FlagOptions() map[string]bool {
    m := map[string]bool{}
    for _, v := range a.Options {
        m[v.Name] = v.Flag
    }
    return m
}

func (a App) FlagCommands() map[string]bool {
    m := map[string]bool{}
    for _, v := range a.Commands {
        m[v.Name] = v.Flag
        m[v.ShortName] = v.Flag
    }
    return m
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
        fmt.Printf("%s\n", o.Help)
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
        fmt.Printf("%s\n", c.Help)
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
func (a *App) Parse(args []string) int {
    
    args_len := uint(len(args))
    
    if args_len == 0 {
        a.Help()
        return 0
    }

    if args[0] == "-h" {
        a.Help()
        return 0
    }

    if args[0] == "-v" {
        fmt.Printf("%s %s\n", a.Name, a.Version)
        return 0
    }
    
    var i uint
    for i = 0; i < args_len; i++ {
        if strings.HasPrefix(args[i], "-") {
            o := string(args[i][1])
            if a.hasOption(o) {
                o_i := a.indexOfOption(o)
                if exitStatus := a.Options[o_i].Parse(args[i:], &i); exitStatus != 0 {
                    return exitStatus
                }
            }
        } else if a.hasCommand(args[i]) {
            c_i := a.indexOfComand(args[i]);
            if exitStatus := a.Commands[c_i].Parse(args[i:], &i); exitStatus != 0 {
                return exitStatus
            }
        } else {
            a.Args = append(a.Args, args[i])
        }
        
    }
    
    return 0
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

// 
func (a App) OptionArgs(s string) ([]string, bool) {
    for _, o := range a.Options {
        if o.Name == s {
            return o.Args, o.Flag
        }
    }
    return []string{}, false
}

func (a App) CommandArgs(s string) ([]string, bool) {
    for _, c := range a.Commands {
        if c.Name == s || c.ShortName == s {
            return c.Args, c.Flag
        }
    }
    return []string{}, false
}
