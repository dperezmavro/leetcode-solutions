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

func (this *MovingAverage) Next(val int) float64 {
	// add new value
	this.currentSum += val

	if len(this.numbers) < this.k {
		this.numbers = append(this.numbers, val)
	} else {
		// remove oldest value
		this.currentSum -= this.numbers[this.oldestElementPtr]

		// store current element in freshest position
		this.numbers[this.oldestElementPtr] = val
	}

	// increment to next position across the ring
	this.oldestElementPtr = (this.oldestElementPtr + 1) % this.k

	return float64(this.currentSum) / float64(len(this.numbers))
}
