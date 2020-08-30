package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

// Win Type
type wchar uint16
type word uint16
type dword uint32
type short int16

type keyEventRecord struct {
	keyDown         bool
	repeatCount     word
	virtualScanCode word
	virtualKeyCode  word
	unicodeChar     wchar
	controlKeyState dword
}

type coord struct {
	x, y short
}

type mouseEventRecord struct {
	mousePosition   coord
	buttonState     dword
	controlKeyState dword
	eventFlags      dword
}

type windowBufferSizeRecord struct {
	size coord
}

type menuEventRecord struct {
	commandId uint
}

type forcusEventRecord struct {
	setForcus bool
}

type inputRecord struct {
	eventType              word
	keyEventRecord         keyEventRecord
	windowBufferSizeRecord windowBufferSizeRecord
	mouseEventRecord       mouseEventRecord
	menuEventRecord        menuEventRecord
	forcusEventRecord      forcusEventRecord
}

var kernel32 = syscall.NewLazyDLL("kernel32.dll")

// BOOL WINAPI ReadConsoleInput(
// _In_  HANDLE        hConsoleInput,
// _Out_ PINPUT_RECORD lpBuffer,
// _In_  DWORD         nLength,
// _Out_ LPDWORD       lpNumberOfEventsRead
// );
var procReadConsoleInput = kernel32.NewProc("ReadConsoleInputW")

func readConsoleInput(fd uintptr, inputrec *inputRecord, length uint32) *keyEventRecord {
	var numOfEvents uint32
	procReadConsoleInput.Call(fd, uintptr(unsafe.Pointer(inputrec)), uintptr(length), uintptr(unsafe.Pointer(&numOfEvents)))
	return &(*inputrec).keyEventRecord
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	inputRec := new(inputRecord)
	fd := os.Stdin.Fd()
	r.Scan()
	kr := readConsoleInput(fd, inputRec, 1)
	if kr.virtualKeyCode == 40 {
		fmt.Println("bye!!")
	}
}
