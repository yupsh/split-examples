package split_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/split"
)

func ExampleSplit_basic() {
	// split -l 10 input.txt output_prefix
	// Note: This writes to files
	yup.MustRun(
		Split("testdata/input.txt", "output_prefix", Lines(10)),
	)
	// Output:
	//
}

