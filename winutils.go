// winutils project winutils.go
package winutils

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32, _        = syscall.LoadLibrary("kernel32.dll")
	getModuleHandle, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")
	moveFileEx, _      = syscall.GetProcAddress(kernel32, "MoveFileExW")

	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")

	shell32, _      = syscall.LoadLibrary("shell32.dll")
	shellExecute, _ = syscall.GetProcAddress(shell32, "ShellExecuteW")
)

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND

	MB_DEFBUTTON1 = 0x00000000
	MB_DEFBUTTON2 = 0x00000100
	MB_DEFBUTTON3 = 0x00000200
	MB_DEFBUTTON4 = 0x00000300

	MOVEFILE_DELAY_UNTIL_REBOOT = 4
)

func abort(funcname string, err error) {
	panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

func MessageBox(caption, text string, style uintptr) (result int) {
	var nargs uintptr = 4
	ret, _, callErr := syscall.Syscall9(uintptr(messageBox),
		nargs,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		style,
		0,
		0,
		0,
		0,
		0)
	if callErr != 0 {
		abort("Call MessageBox", callErr)
	}
	result = int(ret)
	return
}

func GetModuleHandle() (handle uintptr) {
	var nargs uintptr = 0
	if ret, _, callErr := syscall.Syscall(uintptr(getModuleHandle), nargs, 0, 0, 0); callErr != 0 {
		abort("Call GetModuleHandle", callErr)
	} else {
		handle = ret
	}
	return
}

func OpenUrl(url string) (result int) {
	var nargs uintptr = 6
	ret, _, callErr := syscall.Syscall6(uintptr(shellExecute),
		nargs,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("open"))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(url))),
		0,
		0,
		0)
	if callErr != 0 {
		abort("Call ShellExecute", callErr)
	}
	result = int(ret)
	return
}

func RemoveFileOnReboot(file string) (result int) {
	var nargs uintptr = 3
	ret, _, callErr := syscall.Syscall(uintptr(moveFileEx), nargs, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(file))), uintptr(0), MOVEFILE_DELAY_UNTIL_REBOOT)
	if callErr != 0 {
		abort("Call MoveFileEx", callErr)
	}
	result = int(ret)
	return
}
