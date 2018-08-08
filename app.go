package clib

import (
    "fmt"
    "sort"
    _ "strings"
)

// App represents a command line application.
type App struct {
    // Application name
    Name string
    // Application version
    Version string
    // Synopsis
    Synopsis string
    // all Commands
    Commands map[string]*Command
    // all Options
    Options map[string]*Option
    // NoArgUsageFlag represents whether or not to call Usage when executed without argument
    NoArgUsageFlag bool
    // Argument
    args []string
}

// NewApp is a constructor of App
func NewApp(name, version, synopsis string) (*App, error) {
    
    app := &App {
        Name: name,
        Version: version,
        Synopsis: synopsis,
        Commands: map[string]*Command{},
        Options: map[string]*Option{},
    }
    
    if err := app.AddOption("-h", "Display this message", 0); err != nil {
        return nil, err
    }
    
    if err := app.AddOption("-v", "Print version info and exit", 0); err != nil {
        return nil, err
    }
    
    return app, nil
    
}

// Parse is a function to parse the argument
func (a *App) Parse(args []string) (int, error){

    if len(args) == 1 {
        if a.NoArgUsageFlag {
            fmt.Printf(a.Usage())
            return 0, nil
        } else {
            return 0, nil
        }
    }
    args = args[1:]
    argsLen := uint(len(args))
    for i := uint(0); i < argsLen; i++ {
        if a.Options[args[i]] != nil {
            if s, err := a.Options[args[i]].Parse(args[i:], &i); err != nil {
                return s, err
            }
        } else if a.Commands[args[i]] != nil {
            if s, err := a.Commands[args[i]].Parse(args[i:], &i); err != nil {
                return s, err
            }
        } else {
            a.args = append(a.args, args[i])
        }
        
    }
    
    return 0, nil

}

// Args returns argument of myself
func (a App) Args() []string {
    return a.args
}

// Command returns *Command
func (a App) Command(name string) *Command {
    if a.Commands[name] != nil {
        return a.Commands[name]
    }
    return nil
}

// Command returns *Option
func (a App) Option(name string) *Option {
    if a.Options[name] != nil {
        return a.Options[name]
    }
    return nil
}

// CommandFlag returns whether the command was set
func (a App) CommandFlag(name string) (bool, error) {
    if cmd := a.Command(name); cmd != nil {
        return cmd.SetFlag(), nil
    }
    return false, fmt.Errorf("No command: %s", name)
}

// CommandFlag returns whether the option was set
func (a App) OptionFlag(name string) (bool, error) {
    if opt := a.Option(name); opt != nil {
        return opt.SetFlag(), nil
    }
    return false, fmt.Errorf("No option: %s", name)
}


// CommandArgs is a function to get arguments of given command name
func (a App) CommandArgs(name string) ([]string, error) {
    if cmd := a.Command(name); cmd != nil {
        return cmd.Args(), nil
    }
    return []string{}, fmt.Errorf("No command: %s", name)
}

// OptionArgs is a function to get arguments of given option name
func (a App) OptionArgs(name string) ([]string, error) {
    if opt := a.Option(name); opt != nil {
        return opt.Args(), nil
    }
    return []string{}, fmt.Errorf("No option: %s", name)
}

// AddOption is a function to add given option to App
func (a *App) AddOption(name, synopsis string, argCount int) error {
    
    if a.Options[name] != nil {
        return fmt.Errorf("%s option duplicated.\n", name)
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

// Usage returns the usage message created automatically
func (a *App) Usage() (s string) {
    
    s = "Usage: \n\t" + a.Name
    if a.Synopsis != "" {
        s = a.Synopsis + "\n\n" + s
    }

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
        s += "\t" + k
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
