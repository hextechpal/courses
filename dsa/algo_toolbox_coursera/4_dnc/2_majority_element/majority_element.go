package main

//2 3 9 2 2

// {0, 2} [2, 3, 9]      {3, 4}[9, 2, 2]

// [
func majorityElement(arr []int) bool {

	return false
}

func majority(arr []int, l int, r int) int {

	if l == r {
		return arr[l]
	}

	m := l + (r-l)/2

	mLeft := majority(arr, l, m)
	mRight := majority(arr, l, m)

	if mLeft == mRight {
		return mLeft
	}

	return 0
}
