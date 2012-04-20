package main

import (
    "bufio"
    "crypto/md5"
    "fmt"
    "io"
    "os"
    "strings"
    "syscall"
)

var htdata = map[string] string{}

func add_or_change_user(realm string, user string) {
    passwd := read_password()

    hash := md5.New()
    io.WriteString(hash, fmt.Sprintf("%s:%s:%s", user, realm, passwd))
    htdata[ fmt.Sprintf("%s:%s", user, realm) ] = fmt.Sprintf("%x", hash.Sum(nil))
}

func delete_user(realm string, user string) {
    delete(htdata, fmt.Sprintf("%s:%s", user, realm))
}

func load_htfile(htfile string) {
    fh, err := os.Open(htfile)
    if err != nil && err.(*os.PathError).Err == syscall.ENOENT {
        return
    } else if err != nil {
        panic(err)
    }
    defer fh.Close()

    reader := bufio.NewReader(fh)

    for {
        line, _, err := reader.ReadLine()

        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }

        elm := strings.Split(string(line), ":")
        htdata[ fmt.Sprintf("%s:%s", elm[0], elm[1]) ] = elm[2]
    }
}

func save_htfile(htfile string) {
    fh, err := os.Create(htfile)
    if err != nil {
        panic(err)
    }
    defer fh.Close()

    for key, value := range htdata {
        fmt.Fprintf(fh, "%s:%s\n", key, value)
    }
}
