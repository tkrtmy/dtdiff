[![Release](https://img.shields.io/github/release/tkrtmy/dtdiff.svg)](https://github.com/tkrtmy/dtdiff/releases)
[![Build Status](https://travis-ci.org/tkrtmy/dtdiff.svg?branch=master)](https://travis-ci.org/tkrtmy/dtdiff)
[![GoDoc](https://godoc.org/github.com/tkrtmy/dtdiff?status.svg)](https://godoc.org/github.com/tkrtmy/dtdiff)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg?style=plastic)][license]

[license]: https://github.com/tkrtmy/dtdiff/blob/master/LICENCE

# dtdiff
**dtdiff** is **D**ate**T**ime**Diff**

## CLI Usage

```shell
$ dtdiff -help
Usage: dtdiff [options] [datetime(from)] [datetime(to)]

  -h    display hours only
  -m    display minutes only
  -n    display nanoseconds only
  -q    display without time unit
  -s    display seconds only
  -short
        short description
  -until
        calculate until a given one
  -v    display current version

$ dtdiff '2018-10-14 12:40:50' '2018-10-16 15:48:20'
2 days 3 hours 7 minutes 30 seconds

# When you give one argument, calculate between a given and now.
# If you execute command at 2018-10-16 15:48:20
$ dtdiff '2018-10-14 12:40:50'
2 days 3 hours 7 minutes 30 seconds

# When you give one argument with -u option, calculate between now and a given.
# If you execute command at 2018-10-14 12:40:50
$ dtdiff -until '2018-10-16 15:48:20'
2 days 3 hours 7 minutes 30 seconds

$ dtdiff -m '12:40:50' '15:48:20'
187 minutes
$ dtdiff -m -q '12:40:50' '15:48:20'
187

$ dtdiff -m 11:47 18:23
396 minutes
```

## CLI Installation

### For Homebrew user

```shell
$ brew tap tkrtmy/dtdiff
$ brew install dtdiff
```

### For Go user
```shell
$ go get -u github.com/tkrtmy/dtdiff/cmd/dtdiff
```

### How to build & run

```shell
# build in container
$ docker build -t dtdiff ./cmd/dtdiff
# run in container
$ docker run -it --rm dtdiff 12:48 21:23
0 days 8 hours 35 minutes 0 seconds
```
