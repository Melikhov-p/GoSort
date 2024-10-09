package GoSort

func Sort(arr *[]int) {
	var rightLimit = len(*arr) - 1

	for rightLimit >= 0 {
		for i := 0; i < rightLimit; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				(*arr)[i], (*arr)[i+1] = (*arr)[i+1], (*arr)[i]
			}
		}
		rightLimit--
	}
}
