package main

import (
    "fmt"
    "io"
    "os"
    "syscall"
    "strings"
    "crypto/md5"
)


var htdata = map[string] string{}


func add_change_user(realm string, user string) {
    var passwd, pwd_again string
    print("Password: ")
    fmt.Scan(&passwd)
    print("Again: ")
    fmt.Scan(&pwd_again)

    if(passwd != pwd_again) {
        panic("Passwords don't match")
    }

    hash := md5.New()
    io.WriteString(hash, fmt.Sprintf("%s:%s:%s", user, realm, passwd))
    htdata[ fmt.Sprintf("%s:%s", user, realm) ] = fmt.Sprintf("%x", hash.Sum())
}

func delete_user(realm string, user string) {
    htdata[ fmt.Sprintf("%s:%s", user, realm) ] = "", false;
}


func load_htfile(htfile string) {
    fh := fopen_read(htfile)
    defer fh.Close()

    var line string

    for {
        _, err := fmt.Fscanln(fh, &line)

        if len(line) == 0 {
            break
        } else if err != nil {
            panic( err.String() )
        }

        elm := strings.Split(line, ":", 3)
        htdata[ fmt.Sprintf("%s:%s", elm[0], elm[1]) ] = elm[2]

        line = "";
    }
}

func save_htfile(htfile string) {
    fh := fopen_write(htfile)
    defer fh.Close()

    for key, value := range htdata {
        fmt.Fprintf(fh, "%s:%s\n", key, value)
    }
}


func fopen_read(htfile string) (*os.File) {
    fh, err := os.Open(htfile, syscall.O_RDONLY, 0)

    if err != nil {
        switch PathError_to_Errno(err) {
            case syscall.ENOENT:
                return nil

            default:
                panic( err.String() )
        }
    }

    return fh
}

func fopen_write(htfile string) *os.File {
    fh, err := os.Open(htfile, syscall.O_WRONLY | syscall.O_TRUNC | syscall.O_CREAT, 0644)

    if err != nil {
        panic( err.String() )
    }

    return fh
}

func PathError_to_Errno(err os.Error) os.Errno {
    return err.(*os.PathError).Error.(os.Errno)
}
