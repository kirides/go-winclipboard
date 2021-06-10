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

// returns the all FileContents.
func GetFileContents() ([]NamedReadCloser, error)

// returns the FileContents from the specified index
func GetFileContent(index int) (NamedReadCloser, error)
```

## Building this module 

```
> go generate ./...
> go build ./cmd/demo/main.go
```

## Remarks

- Some APIs _do_ require a call to `clipboard.Init()`
    - `GetFileContents`
    - `GetFileContent`
    - they also _might_ require a call to `runtime.LockOSThread()` on the current goroutine
