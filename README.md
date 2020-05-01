# Nap Library

This workshop is a `CLI` project that encapsulates resources and actions for interacting with REST `API's`. Basically, this library allows us to build and call REST `API's` by conforming and acting as a bunch of cohesive layers that represents an abstraction of `GoLang`'s standard `HTTP` package.

## Dependencies

Check the `glide.yaml` file for it.

## Usage

I assume you already have Go installed, with the `GOPATH` set.

Clone this repository inside your `src` directory and install it using the following command

`go install ./cmd/...`

Then, inside `cmd/nap` directory just, `nap []`

For help: `nap -help`

### Available commands

1. `nap -help`: this command will display the help structure for available commands
2. `nap -list`: this shows all the commands available to use
3. `nap get [url]`: (new) the `get` command can also receive an optional POST payload in the request (body structure)

## Resource

The `baseURL` used for this example is the `https://httpbin.org` website.

## Running tests

For doing this, check the `go test -v` command.

This will check and evaluate all the test files for the project, referenced by `_test.go`

## Credits

 - [David E Lares](https://twitter.com/davidlares3)

## License

 - [MIT](https://opensource.org/licenses/MIT)
