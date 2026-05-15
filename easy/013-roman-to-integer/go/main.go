package romantoint

var numberMappings = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	total := 0
	previous := 0
	for _, c := range s {
		current := numberMappings[string(c)]
		total += current
		if previous == 0 {
			previous = current
			continue
		}

		if current > previous {
			total -= 2 * previous
			previous = 0
		} else {
			previous = current
		}
	}
	return total
}
