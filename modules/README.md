
# Introduction
In this lession, I'm gonna introduce modules and module dependencies.

# Lession 1: Version Specific Dependency
## Init repository
First cd to `golessions\modules\remote` and remove `go.mod`

Then use `go mod init` to init a module
<pre>
C:\CODE\golessions\modules\remote> go mod init github.com/philongct/golessions/modules/remote
go: creating new go.mod: module github.com/philongct/golessions/modules/remote
</pre>

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
Test will be failed because Go will download latest version of `rsc.io/sampler`. Fix it by adding a specific version to `go.mod`

<pre>
module github.com/philongct/golessions/modules/remote

go 1.12

require rsc.io/sampler v1.99.99

<b>replace rsc.io/sampler => rsc.io/sampler v1.3.1</b>
</pre>

Run test again, it should PASS
