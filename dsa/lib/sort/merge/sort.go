package merge

func Sort(arr []int) []int{
	sort(arr, 0, len(arr)-1)
	return arr
}

func sort(arr []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		sort(arr, low, mid)
		sort(arr, mid + 1, high)

		merge(arr, low, mid, high)
	}
}

func merge(arr []int, low, mid, high int) {
	
	lSize := mid - low + 1
	rSize := high - mid

	lArr := make([]int, lSize)
	rArr := make([]int, rSize)


	for i := 0; i < lSize; i++ {
		lArr[i] = arr[low + i]
	}

	for j := 0; j < rSize; j++ {
		rArr[j] = arr[mid+ 1 + j]
	}

	var i, j int
	k := low

	for i < lSize && j < rSize {
		if lArr[i] < rArr[j]{
			arr[k] = lArr[i]
			i++	
		}else{
			arr[k] = rArr[j]
			j++	
		}
		k++
	}

	for i < lSize{
		arr[k] = lArr[i]
		i++
		k++	
	}

	for j < rSize{
		arr[k] = rArr[j]
		j++
		k++	
	}
}
