
# Introduction
In this lession, I'm gonna introduce about modules and module dependencies.

# Lession 1: Module & Dependency
## Init repository
First **cd** to `golessions/modules/remote` and remove `go.mod`

Then use `go mod init` to init a module
```
C:\CODE\golessions\modules\remote> go mod init github.com/philongct/golessions/modules/remote
go: creating new go.mod: module github.com/philongct/golessions/modules/remote
```

`github.com/philongct/golessions/modules/remote` is module name, it should match with github location (don't ask why)

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
Test will failed because Go will download latest version of `rsc.io/sampler`. Fix it by forcing a specific version in `go.mod`

<pre>
module github.com/philongct/golessions/modules/remote

go 1.12

require rsc.io/sampler v1.99.99

<b>replace rsc.io/sampler => rsc.io/sampler v1.3.1</b>
</pre>

Run test again, it should PASS

# Lession 2: Local Dependencies
Sometimes we need local dependencies (to try somethings). In such cases we can use `replace` keyword to add local dependencies

In this lession, we'll use `local` module. Open `go.mod` and add location for `local` module:
<pre>
module github.com/philongct/golessions/modules/main

go 1.12

require (
	github.com/philongct/golessions/modules/local v0.0.0
	github.com/philongct/golessions/modules/remote v1.0.0
)

<b>replace github.com/philongct/golessions/modules/local => ../local</b>

</pre>

In `golessions/modules/main`, run it:
```
C:\CODE\golessions\modules\main>go run main.go
go: finding github.com/philongct/golessions/modules/remote v1.0.0
go: finding rsc.io/sampler v1.3.1
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading github.com/philongct/golessions/modules/remote v0.0.0-20190427065944-146b4d2afc42
go: extracting github.com/philongct/golessions/modules/remote v0.0.0-20190427065944-146b4d2afc42
go: downloading rsc.io/sampler v1.3.1
go: extracting rsc.io/sampler v1.3.1
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
Hello, world.
I'm go1.12.4
```

There is also `github.com/philongct/golessions/modules/remote v1.9.9`. You can try with that version

# Lession 3: Using `vendor`
`Vendor` is used when depended modules are offline (No internet connection, missing original modules). Basically this mode is redundant. Read more at [The End of Vendoring](https://research.swtch.com/vgo-module#end_of_vendoring)

First you need Go to copy dependencies to `vendor` folder
```
C:\CODE\golessions\modules\main>go mod vendor
go: finding github.com/philongct/golessions/modules/remote v2.0.0
go: finding rsc.io/sampler v1.3.1
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading github.com/philongct/golessions/modules/remote v0.0.0-20190427065944-146b4d2afc42
go: extracting github.com/philongct/golessions/modules/remote v0.0.0-20190427065944-146b4d2afc42
go: downloading rsc.io/sampler v1.3.1
go: extracting rsc.io/sampler v1.3.1
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
```

Then you can run with `vendor` mode
```
C:\CODE\golessions\modules\main>go run -mod=vendor main.go
Hello, world.
I'm go1.12.4
```
