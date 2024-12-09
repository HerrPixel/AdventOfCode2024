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

func Equal[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Reverse[T any](l []T) []T {
	reversed := make([]T, len(l))

	for i := range l {
		reversed[i] = l[len(l)-i-1]
	}
	return reversed
}

func Sum(l []int) int {
	total := 0
	for _, n := range l {
		total += n
	}

	return total
}
