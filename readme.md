## Exported functions

```go
package winclipboard

// Formats returns a slice that contains all formats currently avaiable in the clipboard
func Formats() ([]int, error)

// FormatName returns a readable name for the passed id.
// Being either a pre-defined name, or through a call to GetClipboardFormatNameW)
func FormatName(id int) (string, error)

// returns a slice containing the filepaths in the H_DROP(15) slot
func GetHDROP() ([]string, error)

// returns a slice containing file metadata (filename + filesize) in the FileGroupDescriptorW slot
func GetFileGroupDescriptor() ([]FileInfo, error)
```

## Generating windows syscall wrappers 

This package uses `golang.org/x/sys/windows/mkwinsyscall` to generate the win32 syscall wrappers

```
> go install golang.org/x/sys/windows/mkwinsyscall
> mkwinsyscall -output zsyscall_windows.go syscall_windows.go
```
