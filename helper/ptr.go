package helper

func StringPtr(val string) *string {
	return &val
}

func IntPtr(val int) *int {
	return &val
}

func BoolPtr(val bool) *bool {
	return &val
}

// PtrOf returns pointer to value.
func PtrOf[T any](v T) *T {
	return &v
}

// PtrTo returns the value of the pointer
func PtrTo[T any](v *T) T {
	if v == nil {
		var zero T
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
