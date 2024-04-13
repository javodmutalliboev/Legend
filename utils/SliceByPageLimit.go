package utils

func SliceByPageLimit[T any](s []T, page, limit int) []T {
	var slice []T
	for i := (page - 1) * limit; i < page*limit && i < len(s); i++ {
		slice = append(slice, s[i])
	}
	return slice
}
