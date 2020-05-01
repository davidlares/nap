# Nap Library

This workshop is a `CLI` project that encapsulates resources and actions for interacting with REST `API's`. Basically, this library allows us to build and call REST `API's` by shaping and acting like a bunch of cohesive layers that present an abstraction of `GoLang`'s standard `HTTP` package, without worrying about the mechanics of the Req/Res cycle.

It also features on calling REST API endpoints, authentication (basic and by tokens) and response callbacks

## Components

The `NAP library` acts as a `client`, this is the interface for the resources.
The `REST resources` are stored in the API structure and the `API endpoints` are the functionality involved.

## HTTP Requests

`Golang` have a standard library for handling `HTTP` requests (send and receive). The main goal of the workshop is preparing a layer for `REST API's` components, resolving `HTTP Headers` for authentication and encryption, and for `Marshal` and `Unmarshal` processes (data serialization and data structure population).


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
