package SliceUtil

// IsEmpty checks if a string slice is empty.
// Parameter slice is the string slice to be checked.
// The return value bool indicates whether the slice is empty (true) or not (false).
func IsEmpty(slice []string) bool {
	return slice == nil || len(slice) == 0
}
