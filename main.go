//go:generate go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=.github/logo.ico
package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	SHCNE_ASSOCCHANGED = 0x08000000
	SHCNF_IDLIST       = 0x0000

	HWND_BROADCAST   = 0xFFFF
	WM_SETTINGCHANGE = 0x001A
	SMTO_ABORTIFHUNG = 0x0002
)

func main() {
	// https://docs.microsoft.com/en-us/windows/desktop/api/shlobj_core/nf-shlobj_core-shchangenotify
	windows.NewLazyDLL("shell32.dll").NewProc("SHChangeNotify").Call(
		uintptr(SHCNE_ASSOCCHANGED),
		uintptr(SHCNF_IDLIST),
		0, 0)

	// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-sendmessagetimeoutw
	env, _ := windows.UTF16PtrFromString("Environment")
	windows.NewLazyDLL("user32.dll").NewProc("SendMessageTimeoutW").Call(
		uintptr(HWND_BROADCAST),
		uintptr(WM_SETTINGCHANGE),
		0,
		uintptr(unsafe.Pointer(env)),
		uintptr(SMTO_ABORTIFHUNG),
		uintptr(5000))
}
