package moving_average_from_data_stream

type MovingAverage struct {
	k                int
	currentSum       int
	numbers          []int
	oldestElementPtr int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		k:       size,
		numbers: make([]int, 0, size),
	}
}

func (ma *MovingAverage) Next(val int) float64 {
	// add new value
	ma.currentSum += val

	if len(ma.numbers) < ma.k {
		ma.numbers = append(ma.numbers, val)
	} else {
		// remove oldest value
		ma.currentSum -= ma.numbers[ma.oldestElementPtr]

		// store current element in freshest position
		ma.numbers[ma.oldestElementPtr] = val
	}

	// increment to next position across the ring
	ma.oldestElementPtr = (ma.oldestElementPtr + 1) % ma.k

	return float64(ma.currentSum) / float64(len(ma.numbers))
}
