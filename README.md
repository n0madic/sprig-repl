# Sprig-REPL

REPL go-template with Sprig functions

## Install

```shell
go get -u github.com/n0madic/sprig-repl
```

## Usage

```shell
$ sprig-repl
Use CTRL + D to exit
>>> {{ "hello!" | upper | repeat 5 }}
HELLO!HELLO!HELLO!HELLO!HELLO!
>>>
```
