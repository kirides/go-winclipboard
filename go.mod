module github.com/kirides/go-winclipboard

go 1.16

require (
    golang.org/x/sys v0.0.0-20210415045647-66c3f260301c
    github.com/kirides/hwnd-go v0.0.0
)

replace (
    github.com/kirides/hwnd-go => ../hwnd-go
)