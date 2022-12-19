package functions

func RemoveWithOrderInt(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
func RemoveWithoutOrderInt[T comparable](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func IndexOf[T comparable](s []T, a T) int {
    for k, v := range s {
        if v == a {
            return k
        }
    }
    return -1
}