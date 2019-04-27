
# Introduction
In this lession, I'm gonna introduce modules and module dependencies.

# Lession 1: Version Specific Dependency
## Init repository
First **cd** to `golessions/modules/remote` and remove `go.mod`

Then use `go mod init` to init a module
```
C:\CODE\golessions\modules\remote> go mod init github.com/philongct/golessions/modules/remote
go: creating new go.mod: module github.com/philongct/golessions/modules/remote
```

## Run Test
Type `go test`
```
C:\CODE\golessions\modules\remote>go test
--- FAIL: TestHello (0.00s)
    hello_test.go:9: Hello() = "99 bottles of beer on the wall, 99 bottles of beer, ...", want "Hello, world."
FAIL
exit status 1
FAIL    github.com/philongct/golessions/modules/remote  0.305s
```
Test will failed because Go will download latest version of `rsc.io/sampler`. Fix it by forcing a specific version to `go.mod`

<pre>
module github.com/philongct/golessions/modules/remote

go 1.12

require rsc.io/sampler v1.99.99

<b>replace rsc.io/sampler => rsc.io/sampler v1.3.1</b>
</pre>

Run test again, it should PASS

# Lession 2: Local Dependencies
We already have a `local` module in `golessions/modules/local`, so we can use it directly by using `replace` keyword

We also have version `v2.0.0` of `remote` module on [golessions](http://github.com/philongct/golessions)

In `golessions/modules/main` Run `go run main`
```
C:\CODE\golessions\modules\main>go run main.go
go: finding github.com/philongct/golessions/modules/remote v2.0.0
go: finding rsc.io/sampler v1.3.1
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading github.com/philongct/golessions/modules/remote v0.0.0-20190427065944-146b4d2afc42
go: extracting github.com/philongct/golessions/modules/remote v0.0.0-20190427065944-146b4d2afc42
go: downloading rsc.io/sampler v1.3.1
go: extracting rsc.io/sampler v1.3.1
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
Hi, world.
I'm go1.12.4
```

Open `go.mod`, replace `github.com/philongct/golessions/modules/remote v2.0.0` with `github.com/philongct/golessions/modules/remote v1.0.0` to see result

```
C:\CODE\golessions\modules\main>go run main.go
go: finding github.com/philongct/golessions/modules/remote v1.0.0
go: downloading github.com/philongct/golessions/modules/remote v1.0.0
go: extracting github.com/philongct/golessions/modules/remote v1.0.0
Hello, world.
I'm go1.12.4
```