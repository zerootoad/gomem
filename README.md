# GoMem

Fork of [github.com/jamesmoriarty/gomem](https://github.com/jamesmoriarty/gomem).

![Continuous Integration](https://github.com/jamesmoriarty/gomem/workflows/Continuous%20Integration/badge.svg?branch=master) ![Latest Tag](https://img.shields.io/github/v/tag/jamesmoriarty/gomem.svg?logo=github&label=latest) [![Go Report Card](https://goreportcard.com/badge/github.com/zerootoad/gomem)](https://goreportcard.com/report/github.com/zerootoad/gomem)

A Go package for manipulating Windows processes. Automated tests manipulate and verify their own process memory via Windows APIs.

```go
import "github.com/zerootoad/gomem"

process, err := gomem.GetOpenProcessFromName("example.exe")
if err != nil {
	panic(err)
}
defer process.Close()

base, err := process.GetModule("example.exe")
if err != nil {
	panic(err)
}

ptr, err := process.ResolvePointer(base+0x01234567, 0x10, 0x20, 0x8)
if err != nil {
	panic(err)
}

value, err := process.ReadUInt32(Ptr)
if err != nil {
	panic(err)
}

_ = process.WriteUInt32(ptr, value+25)
```

## Build

```
go build
```

## Test

```
go test
```

## Docs

[pkg.go.dev/github.com/zerootoad/gomem](https://pkg.go.dev/github.com/zerootoad/gomem)

## Examples

[github.com/jamesmoriarty/gohack](https://github.com/jamesmoriarty/gohack)
