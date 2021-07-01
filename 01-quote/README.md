# 01-quote

Based on https://golang.org/doc/tutorial/getting-started#call

## What did we learn?

- `go mod tidy` finds references to modules in your code, downloads them, and adds them to `go.mod` with `go.sum` entries to validate them.
```
$ go mod tidy
go: finding module for package rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
```
- In my case those libraries were stored under `~/go/pkg/mod`, e.g.
```
└─▪ ls ~/go/pkg/mod/golang.org/x/
text@v0.0.0-20170915032832-14c0d48ead0c

└─▪ ls ~/go/pkg/mod/rsc.io/
quote@v1.5.2   sampler@v1.3.0
```
- `go run quote.go` will run your code
```
$ go run quote.go
Don't communicate by sharing memory, share memory by communicating.
```
- `go build quote.go` will build/compile your code and produce an executable `quote` (i.e. based on src filename)
```
$ file quote
main: Mach-O 64-bit executable x86_64
```
```
$ ./quote
Don't communicate by sharing memory, share memory by communicating.
```