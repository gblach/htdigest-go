package main

// #include <termios.h>
// #include <unistd.h>
// struct termios T_orig, T;
import "C"

import (
    "fmt"
    "os"
)

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
