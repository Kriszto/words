# Lottery test task

Lottery is a go program for a software engineering test.

## Installation and building

- Download this repo from github

- Build with make (requires make)
```bash
make build
```

or

- Build with command line
```bash
go build -o bin/scrmabledstrings main.go
```

## Usage

#### Basic usage:
```bash
bin/scrmabledstrings 10m-v2.txt
```


#### Extended usage
```
Usage: bin/lottery [-v] [-t 8] [-b 100000] [-h] <filename>
  filename: file to process that contains the test data
  -b int
        number of tickets in a batch process (default 100000)
  -h    help
  -t int
        number of threads, default set to the number of processors
  -v    prints out times needed to the calculations
```

## Useful utilities

#### run tests:
runs all unit tests
```bash
make test
```

#### linter:
runs linter and shows linter issues
```bash
make lint
```

#### test coverage:
runs all tests with coverage and shows the result in a browser
```bash
make coverage
```

## Asymptotic run time

### Processing the input file
The input file processing is simple, reads line by line from a file, parses the lines, and appends to a `[][]byte` slice.
The cost of processing is `O(n)`

### Reporting
The algorithm stores the result in a `map[byte][byte]`,
iterates through the `[][]byte` slice one by one, 
calculates the matching numbers in a for loop, and adds to the result map.

The execution of the algorithm is multi-threaded, but it doesn't affect its asymptotic complexity, it's `O(n)`.

