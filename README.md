# golocaterepo

Have you ever run across a Go binary and you want to look at the source code but you forgot where you installed it from? Then this is the project for you!

## Usage

```
❯ golocaterepo -h
Description: prints the source repository found in the strings of a Go binary.

Usage: golocaterepo [binary in your PATH]

Options:
  -v, --verbose   Verbose error output
```

### Example

	❯ golocaterepo chroma
	github.com/alecthomas/chroma/cmd/chroma

## Installation

	go get github.com/jakewarren/golocaterepo

## Credits

Almost all the code is adapted from the wonderful [christophberger/goman](https://github.com/christophberger/goman) project.
