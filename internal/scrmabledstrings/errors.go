package scrmabledstrings

import "fmt"

// NumberOutOfRangeError error for cases where a number is out of range
type NumberOutOfRangeError struct {
	Input string
	Err   error
}

func (r *NumberOutOfRangeError) Error() string {
	return fmt.Sprintf("invalid input %s, must be a number and between 1 and 90", r.Input)
}

// NumberInvalidFormatError error for cases where a number's format is not a valid int
type NumberInvalidFormatError struct {
	Input string
	Err   error
}

func (r *NumberInvalidFormatError) Error() string {
	return fmt.Sprintf("invalid input %s, must be an integer", r.Input)
}

// InvalidRowCountError error for cases where a row contains less or more number than expected
type InvalidRowCountError struct {
	Numbers []byte
	Err     error
}

func (r *InvalidRowCountError) Error() string {
	return fmt.Sprintf("invalid input length: %d in %v", len(r.Numbers), r.Numbers)
}

// DuplicatedNumberError error for cases where a row contains one or more duplicated numbers
type DuplicatedNumberError struct {
	Numbers []byte
	Err     error
}

func (r *DuplicatedNumberError) Error() string {
	return fmt.Sprintf("duplicated number in row: %v", r.Numbers)
}
