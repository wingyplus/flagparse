# flagparse

Analysis tool to detecting calls of `flag.Parse()` during `init()`.

## Install

```
$ go get github.com/wingyplus/flagparse/cmd/flagparse
$ go vet -vettool=$(which flagparse) ./...
```

## NOTE

This tool is workaround during wait for [CL218757](https://go-review.googlesource.com/c/tools/+/218757/). to be
merge.
