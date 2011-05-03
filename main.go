package main

import (
    "fmt"
    "os"
    "tcattr"
)

func catch() {
    if err := recover(); err != nil {
        fmt.Fprintf(os.Stderr, "%s failed: %s\n", os.Args[0], err)
        os.Exit(1)
    }
}

func usage() {
    fmt.Printf("Usage: %s <htfile> <add|del> <realm> <user>\n", os.Args[0])
    os.Exit(0)
}

func main() {
    defer catch()

    if 5 != len(os.Args) {
        usage()
    }

    htfile := os.Args[1]
    action := os.Args[2]
    realm  := os.Args[3]
    user   := os.Args[4]

    tc := tcattr.BindFd(os.Stdin.Fd())
    tc.LocalUnset(tcattr.ECHO)
    tc.LocalSet(tcattr.ECHONL)
    tc.Apply(tcattr.TCSANOW)

    switch action {
        case "add":
            load_htfile(htfile)
            add_change_user(realm, user)

        case "del":
            load_htfile(htfile)
            delete_user(realm, user)

        default:
            usage()
    }

    tc.LocalSet(tcattr.ECHO)
    tc.Apply(tcattr.TCSANOW)

    save_htfile(htfile)
}
