package goshopee

func Ptr[T any](v T) *T {
	return &v
}
