package array

func Filter[T any](arr []T, predicate func(item T, idx int) bool) []T {
	var filteredArr []T

	for i, s := range arr {
		if predicate(s, i) {
			filteredArr = append(filteredArr, s)
		}
	}

	return filteredArr
}

func Reduce[T any, U any](arr []T, initialValue U, fn func(i int, acc U, item T) U) U {
	acc := initialValue

	for i, item := range arr {
		acc = fn(i, acc, item)
	}

	return acc
}
