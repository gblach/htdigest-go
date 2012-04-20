package main

// #include <termios.h>
// #include <unistd.h>
// struct termios T_orig, T;
import "C"

import (
    "fmt"
    "os"
)

func read_password() string {
    print("Password: ")
    passwd := stdin_read_password()

    print("Again: ")
    pwd_again := stdin_read_password()

    if(passwd != pwd_again) {
        panic("Passwords don't match")
    }

    return passwd
}

func stdin_read_password() string {
    stdin := C.int(os.Stdin.Fd())

    C.tcgetattr(stdin, &C.T)
    C.T_orig = C.T

    C.T.c_lflag &^= C.ECHO
    C.T.c_lflag |= C.ECHONL

    C.tcsetattr(stdin, C.TCSANOW, &C.T)

    var passwd string
    fmt.Scanf("%s", &passwd)

    C.tcsetattr(stdin, C.TCSANOW, &C.T_orig)

    return passwd
}
