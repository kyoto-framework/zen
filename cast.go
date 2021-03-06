/*
	-

	Cast

	Zen provides a comfortable way to work with slices casting ([]any).

	Example:

		func main() {
			// Define a sample slice
			values := []any{1, 2, 3}
			// Demonstrate casting
			castedValues := zen.CastSlice[int](values) []int{1, 2, 3}
		}
*/
package zen

// CastSice is a function to cast a slice of any values ([]any)
// to a slice of the given type.
//
// Usage:
//  zen.CastSlice[int]([]any{1, 2, 3}) // []int{1, 2, 3}
func CastSlice[T any](slice []any) []T {
	_slice := make([]T, 0, len(slice))
	for _, el := range slice {
		_slice = append(_slice, el.(T))
	}
	return _slice
}

// CastPSlice is a function to cast a slice of any values ([]any)
// to a slice of the given type pointers.
//
// Usage:
//  zen.CastPSlice(int)([]any{1, 2, nil}) []*int{1, 2, nil}
func CastPSlice[T any](slice []any) []*T {
	_slice := make([]*T, 0, len(slice))
	for _, v := range slice {
		if v == nil {
			_slice = append(_slice, nil)
		} else {
			_v := v.(T)
			_slice = append(_slice, &_v)
		}
	}
	return _slice
}
