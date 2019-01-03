package sortmy

func Insertionsort(a []int) []int {
	ln := len(a) - 1

	for i := range a {
		for j := i; j >= 0; j-- {
			if j < ln && a[j+1] < a[j] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}
