# Go Syncmap

An wrapper around `sync.Map` that uses generics for the key and value rather than `any`.

## Getting Started

### Installing

```sh
go get github.com/kagadar/go-syncmap
```

### Usage

```go
import "github.com/kagadar/go-syncmap"
```

As is the case with the wrapped type: "The zero Map is empty and ready for use. A Map must not be copied after first use."
