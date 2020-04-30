# Nap Library

This workshop is a `CLI` project that encapsulates resources and actions for interacting with REST `API's`. Basically, this library allows to build and call REST `API's` by conforming and acting as a bunch of cohesive layers that represents an abstraction of the `GoLang`'s standard `http` package.

## Dependencies

Check the `glide.yaml` file for it.

## Usage

I assume you already have Go installed, with the `GOPATH` set.

Clone this repository inside your `src` directory and install it using the following command

`go install ./cmd/...`

Then

`go build` or `go run` to the `main.go` file inside the `cmd/nap` directory.

## Resource

The `baseURL` used for this example is the `https://httpbin.org` website.

## Running tests

For doing this, check the `go test -v` command.

This will check and evaluate all the test files for the project, referenced by `_test.go`

## Credits

 - [David E Lares](https://twitter.com/davidlares3)

## License

 - [MIT](https://opensource.org/licenses/MIT)
