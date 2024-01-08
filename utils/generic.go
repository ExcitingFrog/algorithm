package utils

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Sum[T Integer](a []T) T {
	var s T
	for _, v := range a {
		s += v
	}

	return s
}
