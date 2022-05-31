package lex

func Filter[T any](f func(T) bool, list []T) []T {
	out := make([]T, 0)
	for _, v := range list {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}
