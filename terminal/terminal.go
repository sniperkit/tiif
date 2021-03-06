package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

func Prompt(q string) (input string, err error) {
	var line []byte

	fmt.Printf(q)
	line, _, err = bufio.NewReader(os.Stdin).ReadLine()
	input = string(line)
	return
}

func IntPrompt(q string) (input int, err error) {
	var line string
	line, err = Prompt(q)
	if err != nil {
		return
	}

	input, err = strconv.Atoi(line)
	return
}

// Props to nsf/termbox-go
type winsize struct {
	rows    uint16
	cols    uint16
	xpixels uint16
	ypixels uint16
}

func Dimensions() (int, int) {
	var sz winsize
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&sz)),
	)

	return int(sz.cols), int(sz.rows)
}
