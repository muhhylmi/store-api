package objects

func ToPointer[T any](v T) *T {
	return &v
}

func AnyInSlice(s []string, v string) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}
	return false
}
