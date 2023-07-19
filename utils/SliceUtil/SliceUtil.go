package SliceUtil

func IsEmpty(slice []string) bool {
	return slice == nil || len(slice) == 0
}
