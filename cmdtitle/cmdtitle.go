/*
* @Author: Try
* @Date:   2021/5/5 22:50
 */
package cmdtitle

import (
	"syscall"
	"unsafe"
)

func SetTitle(title string) {
	kernel32, _ := syscall.LoadLibrary(`kernel32.dll`)
	sct, _ := syscall.GetProcAddress(kernel32, `SetConsoleTitleW`)
	syscall.Syscall(sct, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	syscall.FreeLibrary(kernel32)
}
