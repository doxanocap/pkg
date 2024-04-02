package sorting

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, l, h int) []int {
	if l < h {
		p := partition(a, l, h)
		a = quickSort(a, l, p-1)
		a = quickSort(a, p+1, h)
	}
	return a
}

func partition(a []int, l, h int) int {
	if l == h {
		return l
	}
	// pivot
	p := a[l]
	i, j := l, h

	for i < j {
		for a[i] <= p && i < h {
			i++
		}
		for a[j] > p && j > l {
			j--

			if i < j {
				a[i], a[j] = a[j], a[i]
			}
		}

		a[l], a[j] = a[j], a[l]
		return j
	}
	return j
}
