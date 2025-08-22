package tools

func GetPointer[T any](t T) *T {
	return &t
}

func GetValue[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}
