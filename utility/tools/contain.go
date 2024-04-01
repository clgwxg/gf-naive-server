package tools

func Contain[T comparable](list []T, item T) bool {
	inList := false
	for _, value := range list {
		if value == item {
			inList = true
		}

	}
	return inList
}
