# clib

[![Build Status](https://travis-ci.org/owlinux1000/clib.svg?branch=master)](https://travis-ci.org/owlinux1000/clib)
[![GoDoc](https://godoc.org/github.com/owlinux1000/clib?status.svg)](https://godoc.org/github.com/owlinux1000/clib)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE.txt)


clib is yet another golang command line parser.

## Install

```
$ go get github.com/owlinux1000/clib
```

## Example

Here is a sample implemented like ```git add -u``` .

```
package main

import (
    "os"
    "fmt"
    "github.com/owlinux1000/clib"
)

func main() {

    app, err := clib.NewApp("mygit", "1.0.0", "A toy git client")
    if err != nil {
        panic(err)
    }
    
    if err := app.AddCommand("add", "a", "Add file contents to the index", 0); err != nil {
        panic(err)
    }
    
    if err := app.AddOption("-u", "update tracked files", 0); err != nil {
        panic(err)
    }

    if len(os.Args) == 0 {
        os.Exit(0)
    }
    
    exitStatus, err := app.Parse(os.Args[1:])
    if err != nil {
        panic(err)
    }

    if ok, _ := app.OptionFlag("-h"); ok {
        fmt.Println(app.Usage())
        os.Exit(exitStatus)
    }

    if ok, _ := app.CommandFlag("add"); ok {
        fmt.Println("You executed `mygit add`")
        if ok, _ := app.OptionFlag("-u"); ok {
            fmt.Println("You executed `mygit add -u`")
        }
    }
    
}
```
