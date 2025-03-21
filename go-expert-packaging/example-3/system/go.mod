module github.com/gilsondev/estudos-go/go-expert-packaging/example-3/system

go 1.22.1

// go mod edit -replace github.com/gilsondev/estudos-go/go-expert-packaging/example-3/math=../math
replace github.com/gilsondev/estudos-go/go-expert-packaging/example-3/math => ../math

require github.com/gilsondev/estudos-go/go-expert-packaging/example-3/math v0.0.0-00010101000000-000000000000
