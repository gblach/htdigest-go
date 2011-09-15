package main

import (
    "fmt"
    "io"
    "os"
    "syscall"
    "bufio"
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
    fh, err := os.Open(htfile)
    if err != nil {
        switch PathError_to_Errno(err) {
            case syscall.ENOENT:
                return

            default:
                panic(err.String())
        }
    }
    defer fh.Close()

    reader := bufio.NewReader(fh)

    for {
        line, _, err := reader.ReadLine()

        if len(line) == 0 {
            break
        } else if err != nil {
            panic(err.String())
        }

        elm := strings.Split(string(line), ":")
        htdata[ fmt.Sprintf("%s:%s", elm[0], elm[1]) ] = elm[2]
    }
}

func save_htfile(htfile string) {
    fh, err := os.Create(htfile)
    if err != nil {
        panic(err.String())
    }
    defer fh.Close()

    for key, value := range htdata {
        fmt.Fprintf(fh, "%s:%s\n", key, value)
    }
}

func PathError_to_Errno(err os.Error) os.Errno {
    return err.(*os.PathError).Error.(os.Errno)
}
