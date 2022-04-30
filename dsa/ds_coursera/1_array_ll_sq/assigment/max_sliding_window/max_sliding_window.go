package max_sliding_window

func maxSliding(nums []int, wsize int) []int{
	m1Idx := 0
	m2Idx := 0
	wmax := make([]int, 0)
	for i := 1; i < wsize; i++ {
		if nums[m1Idx] < nums[i] {
			m2Idx = m1Idx
			m1Idx = i
		}else if nums[m2Idx] < nums[i]{
			m2Idx = i
		}
	}

	wmax = append(wmax, nums[m1Idx])

	for i := wsize; i < len(nums); i++ {
		if m1Idx == i-wsize  {
			if nums[m2Idx] < nums[i] {
				wmax = append(wmax, nums[i])
				m1Idx = i
			}else{
				m1Idx = m2Idx
			}
		}else if m2Idx == i-wsize {
			if nums[m1Idx] < nums[i] {
				wmax = append(wmax, nums[i])
				m1Idx = i
				m2Idx = m1Idx
			}else if nums[m2Idx] < nums[i]{
				wmax = append(wmax, nums[m1Idx])
				m2Idx = i					
			}
		}else if (nums[m1Idx] < nums[i]){
			wmax = append(wmax, nums[i])
			m1Idx = i
			m2Idx = m1Idx
		}else if (nums[m2Idx] < nums[i]) {
			wmax = append(wmax, nums[m1Idx])
			m2Idx = i	
		}else{
			wmax = append(wmax, nums[m1Idx])
		}
	}
	return wmax
}
