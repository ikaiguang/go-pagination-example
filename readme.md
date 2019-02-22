# go-pagination

The fantastic pagination library for Golang, aims to be developer friendly.

## Install

install && get source code

`go get -v github.com/ikaiguang/go-pagination`

> go_1.11+ : `GO111MODULE=off go get -v github.com/ikaiguang/go-pagination`

## Test

before run test : you must rewrite database connection and generate test data

- [models/common.go](models/common.go)
- [models/test_data.go](models/test_data.go)

go test -v .

## License

Released under the [MIT License](License)