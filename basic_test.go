package split_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/split"
)

func ExampleSplit_basic() {
	// split -l 10 input.txt output_prefix
	// Note: This writes to files
	gloo.MustRun(
		Split("testdata/input.txt", "output_prefix", Lines(10)),
	)
	// Output:
	//
}
