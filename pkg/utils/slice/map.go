package slice

func Map[V any, U any](sources []U, transfer func(u U) V) []V {
	result := make([]V, len(sources))
	for i, el := range sources {
		result[i] = transfer(el)
	}
	return result
}
