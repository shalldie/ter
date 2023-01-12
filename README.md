# ter

[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/ter?label=go&logo=go&style=flat-square)](https://github.com/shalldie/ter)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/ter.svg)](https://pkg.go.dev/github.com/shalldie/ter)
[![Test Status](https://img.shields.io/github/actions/workflow/status/shalldie/ter/ci.yml?branch=master&label=test&logo=github&style=flat-square)](https://github.com/shalldie/ter/actions)
[![License](https://img.shields.io/github/license/shalldie/ter?logo=github&style=flat-square)](https://github.com/shalldie/ter)

Ternary operator in golang. Golang 三元运算。

## Installation

```bash
go get github.com/shalldie/ter
```

## Methods&Examples

```go
// True(a, b, c) such as a ? b : c

ter.True(true, 1, 2)  // 1
ter.True(false, 1, 2) // 2
```

```go
ter.Truthy(1, 1, 2)     // 1

// Truth treat these as false: false, 0, "", nil

ter.Truthy(false, 1, 2) // 2
ter.Truthy(0, 1, 2)     // 2
ter.Truthy("", 1, 2)    // 2
ter.Truthy(nil, 1, 2)   // 2
```

## License

MIT
