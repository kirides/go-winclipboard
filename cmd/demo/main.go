package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	clipboard "github.com/kirides/go-winclipboard"
	"github.com/kirides/go-winclipboard/wm"

	"github.com/kirides/hwnd-go"
)

func main() {
	// Create a hWnd with a respective WndProc
	h, err := hwnd.New(func(h syscall.Handle, msg uint32, wParam, lParam uintptr) uintptr {
		if msg == wm.CLIPBOARDUPDATE {
			f, err := clipboard.Formats()
			if err != nil {
				panic(err)
			}

			for _, v := range f {
				name, err := clipboard.FormatName(v)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%s (%d)\n", name, v)
				if fn, ok := printClipFormat[name]; ok {
					fn()
				}
			}
		}

		fmt.Printf("WndProc(%v, %s, %v, %v)\n", h, wm.String(msg), wParam, lParam)

		// Should always call hwnd.DefWindowProc if not handled otherwise
		// refer to general Win32 API programming
		return hwnd.DefWindowProc(h, msg, wParam, lParam)
	})
	if err != nil {
		panic(err)
	}

	// Register for Clipboard change notification
	clipboard.AddClipboardFormatListener(h.Handle)
	defer clipboard.RemoveClipboardFormatListener(h.Handle)

	// support graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	// Process "window" message queue
	h.ProcessMessagesContext(ctx)
}

var printClipFormat = map[string]func(){
	"FileGroupDescriptorW": func() {
		d, err := clipboard.GetFileGroupDescriptor()
		if err != nil {
			fmt.Printf("    ERR: %v\n", err)
			return
		}
		for _, v := range d {
			fmt.Printf("    %v (%d b)\n", v.Name, v.Size)
		}
	},
	"CF_HDROP": func() {
		d, err := clipboard.GetHDROP()
		if err != nil {
			fmt.Printf("    ERR: %v\n", err)
			return
		}
		for _, v := range d {
			fmt.Printf("    %v\n", v)
		}
	},
	"Shell IDList Array": func() {
		d, err := clipboard.GetShellIDListArray()
		if err != nil {
			fmt.Printf("    ERR: %v\n", err)
			return
		}
		for _, v := range d {
			fmt.Printf("    %v\n", v)
		}
	},
}
