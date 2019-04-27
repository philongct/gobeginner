module github.com/philongct/golessions/modules/main

go 1.12

require (
	github.com/philongct/golessions/modules/local v0.0.0
	github.com/philongct/golessions/modules/remote v1.0.0
)

replace github.com/philongct/golessions/modules/local => ../local
