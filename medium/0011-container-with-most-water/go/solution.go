package simplifypath

func maxArea(height []int) int {
	if len(height) == 2 {
		return min(height[0], height[1])
	}

	left := 0
	right := len(height) - 1
	minHeight := min(height[left], height[right])
	maxVolume := minHeight * (right - left)

	for left < right {
		if height[left] < height[right] {
			left++
		} else if height[left] >= height[right] {
			right--
		}
		minHeight = min(height[left], height[right])
		tmpVolume := minHeight * (right - left)
		if maxVolume < tmpVolume {
			maxVolume = tmpVolume
		}
	}

	return maxVolume
}
