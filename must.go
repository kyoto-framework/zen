/*
	-

	Must

	Zen provides a simple  helper function that wraps a call to a function
	returning value and error, and panics if the error is non-nil.

	Example:

		func main() {
			zen.Must(strconv.Atoi("asd")) // panic
		}
*/
package zen

// Must is a helper that wraps a call to a function returning
// and panics if the error is non-nil.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
