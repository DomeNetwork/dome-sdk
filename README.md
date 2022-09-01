# ḎOME Network - SDK

## Requirements

* `docker`
* `docker-compose`
* `go v1.18+`
* `godoc`
* `swag`
* `swagger`

## Build

To build the WASM module run `bash scripts/build.sh wasm`.

The build system takes advantages of build tags to allow the inclusion of specific files for targeted architectures.

## Test

To test the Go code run `bash scripts/test.sh` in a console.

## Structure

A generalized overview of the folder structure.

* `/build` - holds the built WASM and related files
* `/pkg` - various packages
  * `/wasm` - adapter code that handles JS <-> Go
* `/scripts` - helpful bash scripts
  * `/build.sh` - simplified building options
  * `/test.sh` - run tests, runs `go test ./...` but used to do more
* `/sdk` - DOME SDK
* `/ui` - some example UIs taking advantage of WASM
  * `/test` - simplified test JS and Go server
  * `/wallet` - an example wallet UI that uses Vue and WASM
* `/wasm` - container to build WASM
