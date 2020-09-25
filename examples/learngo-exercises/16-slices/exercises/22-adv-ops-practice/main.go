
package main

import (
	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Practice advanced slice operations
//
//  Please follow the directions in the following code.
//
//  To see the expected output, please run:
//
//    go run solution/main.go
//
// ---------------------------------------------------------

func main() {
	// ########################################################
	//
	// #1: Create a string slice: `names` with a length and
	//     capacity of 5, and print it.
	//
	//
	// ...
	// s.Show("1st step", names)

	// ########################################################
	//
	// #2: Append the following names to the names slice:
	//
	//     "einstein", "tesla", "aristo"
	//
	//     Print the names slice.
	//
	//     Observe how the slice and its backing array change.
	//
	//
	// ...
	// s.Show("2nd step", names)

	// ########################################################
	//
	// #3: The previous code appends at the end of the names
	//     slice:
	//
	//     ["" "" "" "" "" "einstein", "tesla", "aristo"]
	//
	//     Append the new elements to the beginning of the names
	//     slice instead:
	//
	//     ["einstein", "tesla", "aristo" "" ""]
	//
	//     So: Overwrite and print the names slice.
	//
	// ...
	// s.Show("3rd step", names)

	// ########################################################
	//
	// #4: Copy only the first two elements of the following
	//     array to the last two elements of the `names` slice.
	//
	//     Print the names slice, you should see 5 elements.
	//     So, do not forget extending the slice.
	//
	//     Observe how its backing array stays the same.
	//
	//
	// Array (uncomment):
	// moreNames := [...]string{"plato", "khayyam", "ptolemy"}
	//
	// ...
	//
	// s.Show("4th step", names)

	// ########################################################
	//
	// #5:  Only copy the last 3 elements of the `names` slice
	//      to a new slice: `clone`.
	//
	//     Append the first two elements of the `names` to the
	//    `clone`.
	//
	//     Ensure that after appending no new backing array
	//     allocations occur for the `clone` slice.
	//
	//     Print the clone slice before and after the append.
	//
	//
	// ...
	// s.Show("5th step (before append)", clone)
	//
	// ...
	// s.Show("5th step (after append)", clone)

	// ########################################################
	//
	// #6: Slice the `clone` slice between 2nd and 4th (inclusive)
	//     elements into a new slice: `sliced`.
	//
	//     Append "hypatia" to the `sliced`.
	//
	//     Ensure that new backing array allocation "occurs".
	//
	//       Change the 3rd element of the `clone` slice
	//       to "elder".
	//
	//       Doing so should not change any elements of
	//       the `sliced` slice.
	//
	//     Print the `clone` and `sliced` slices.
	//
	//
	// ...
	// s.Show("6th step", clone, sliced)
}

//
// Don't mind about this function.
//
// For printing the slices: You can either use the
// prettyslice package or `fmt.Printf`.
//
func init() {
	s.PrintBacking = true // prints the backing array
	s.MaxPerLine = 10     // prints 10 slice elements per line
	s.Width = 60          // prints 60 character per line
}
