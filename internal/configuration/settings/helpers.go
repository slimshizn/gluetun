package settings

func ptrTo[T any](value T) *T {
	return &value
}
