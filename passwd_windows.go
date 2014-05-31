// +build windows

package main

// #include <windows.h>
import "C"

import (
	"fmt"
	"os"
)

const inputLength = 256

func stdin_read_password() string {
	var in = C.GetStdHandle(C.STD_INPUT_HANDLE)
	var old_flags C.DWORD
	var isTty = (C.GetFileType(in) == C.FILE_TYPE_CHAR)
	if isTty {
		if C.GetConsoleMode(in, &old_flags) != C.FALSE {
			C.SetConsoleMode(in, C.ENABLE_LINE_INPUT|C.ENABLE_PROCESSED_INPUT)
		} else {
			isTty = false
		}
	}

	var input [inputLength]byte
	var length = inputLength
	var count C.DWORD
	C.ReadFile(in, C.PVOID(&input[0]), inputLength, &count, nil)
	var countInt = int(count)
	if countInt >= 2 && input[countInt-2] == '\r' {
		length = countInt - 2
	} else {
		var buf [inputLength]byte
		for C.ReadFile(in, C.PVOID(&buf[0]), inputLength, &count, nil) > 0 {
			if count >= 2 && buf[count-2] == '\r' {
				break
			}
		}
	}

	fmt.Fprintln(os.Stderr)

	if isTty {
		C.SetConsoleMode(in, old_flags)
	}
	return string(input[:length])
}
