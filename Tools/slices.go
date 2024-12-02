package Tools

func Remove[T any](l []T, i int) []T {
	removedList := make([]T, i, len(l)-1)
	_ = copy(removedList, l[:i])
	removedList = append(removedList, l[i+1:]...)
	return removedList
}

func Clone[T any](l []T) []T {
	clone := make([]T, len(l))
	_ = copy(clone, l)
	return clone
}
