# Sprig-REPL

REPL go-template with [Sprig](http://masterminds.github.io/sprig/) functions

## Install

```shell
go get -u github.com/n0madic/sprig-repl
```

## Usage

```shell
$ sprig-repl
Use CTRL + D to exit
> {{ "hello!" | upper | repeat 5 }}
HELLO!HELLO!HELLO!HELLO!HELLO!
> "HELLO" | lower
hello
>
```

Brackets can be omitted.
