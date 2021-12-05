# Words
Code Challenge Solution

## Basic commands

### Run the tests

#### in docker
```shell
make docker-test
```
#### in bash
```shell
make test
```

### Run golangci-lint

#### in docker
```shell
make docker-lint
```
#### in bash
```shell
make lint
```

### Build the program

#### in docker
```shell
make docker-build
```

Builds a docker image  with version tag

#### in bash
```shell
makeb uild
```

Builds the program in the current directory


### Run the program
There are several test files in the [testdata/generated/ts1](testdata/generated/ts1) directory.
Use them in pair, eg. [testdata/generated/ts1/dictionary_1.txt](testdata/generated/ts1) with [testdata/generated/ts1/input_1.txt](testdata/generated/ts1)

#### Run in docker
```shell
make docker-run INPUT=testdata/generated/ts1/input_1.txt DICT=testdata/generated/ts1/dictionary_1.txt LOG_LEVEL=1
```

**Parameters**:
- **INPUT**: mandatory, path to input file
- **DICT**: mandatory, path to input dictionary file
- **LOG_LEVEL**: optional, level of logging
  - **0**: no logging
  - **1**: info level logging
  - **2**: debug level logging

The output is a line per input string:
```
Case #1: 2
```
Where first number is the line number and the second one is the number of words from the dictionary that appear (in their original or scrambled form) as substrings of the given string.

#### Run in bash
```shell
./scrambled-strings --input testdata/generated/ts1/input_1.txt  -d testdata/generated/ts1/dictionary_1.txt -l 1
```

Usage:
```shell
Usage:
  scrambled-strings [flags]
  scrambled-strings [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  generate    Generates a set of test cases from original tests.
  help        Help about any command

Flags:
  -d, --dictionary string   Dictionary filename
  -h, --help                help for scrambled-strings
  -i, --input string        Input filename
  -l, --log-level int       log level, 0 (default): no logging, 1: info level, 2: debug level
  -v, --verbose             Verbose output

Use "scrambled-strings [command] --help" for more information about a command.
```

## Helper commands

### Show test coverage

#### in bash

```shell
make coverage
```
Shows test coverage in the default browser.

### Generating test files from original [test sets](testdata/original)

#### in bash

```shell
./scrambled-strings generate
```
Generates 20 test from original dataset to use as input and dictionary test.
