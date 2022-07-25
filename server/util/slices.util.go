package util;

func PtrSliceToValueSlice[T interface{}](slice []*T) []T {
	result := make([]T, len(slice))

	for _, value := range(slice) {
		result = append(result, *value)
	}

	return result;
}
