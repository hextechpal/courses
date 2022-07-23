package quick

func Sort(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		sort(arr, low, pivot-1)
		sort(arr, pivot+1, high)
	}

}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
