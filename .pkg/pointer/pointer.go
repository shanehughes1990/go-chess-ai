package pointer

// Pointer is a generics way of referecing the pointer to a value
func Pointer[V any](v V) *V {
	return &v
}

// Deref is a generics nil safe way of dereferencing a pointer
func Deref[V any](v *V) V {
	if v == nil {
		return *new(V)
	}

	return *v
}
