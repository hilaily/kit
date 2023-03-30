package helper

// PtrOf returns pointer to value.
func PtrOf[T any](v T) *T {
	return &v
}

// PtrTo returns the value of the pointer
func PtrTo[T any](v *T) T {
	var zero T
	if v == nil {
		return zero
	}
	return *v
}

func GetWithDef[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}
