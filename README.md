# gopipe-redis
[![Build Status](https://travis-ci.org/duncanleo/gopipe-redis.svg?branch=master)](https://travis-ci.org/duncanleo/gopipe-redis)

gopipe-redis is a small utility written in Go to generate a file for mass insertion into the key-value store Redis. It takes in a source file written using the human-readable Redis syntax.

### Installation
```
$ go get github.com/duncanleo/gopipe-redis
```

### Usage
```
$ gopipe-redis -i [source file]
```

### Example
A sample source file, `sample_source.txt`, is included. Here's what it looks like:  

```
SET price 99.99
SET color red
SET unit Celsius
```

Running this through the utility produces the following output:

```
*3
$3
SET
$5
price
$5
99.99
*3
$3
SET
$5
color
$3
red
*3
$3
SET
$4
unit
$7
Celsius
```

Using the command below, it can be piped into Redis' `redis-cli` command-line utility for mass insertion.

```
$ gopipe-redis -i sample_source.txt | redis-cli --pipe
```
