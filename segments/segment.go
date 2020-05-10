package segments

// GetSegmentv4 for value
func GetSegmentv4(i float64, MAX float64) int {
	//revenue = rand.Intn(max-min) + min
	//value := max - min // max - min = 1 000 000 - 100 000 = 900 000
	value := float64(i)
	_max := float64(MAX)
	percent := value / _max
	segment := 0
	if (percent > float64(0.0)) && (percent <= float64(0.25)) {
		segment = 0
	}
	if (percent > float64(0.25)) && (percent <= float64(0.50)) {
		segment = 1
	}
	if (percent > float64(0.50)) && (percent <= float64(0.75)) {
		segment = 2
	}
	if (percent > float64(0.75)) && (percent <= float64(1.00)) {
		segment = 3
	}
	//fmt.Println(fmt.Sprintf("value: %d - %d - %f (%d)", max, i, percent, segment))
	return segment
}

// getSegmentv5 documentation
// Här ska vi kunna returnera baserad på en distribution av 4 värden 0.0 - 1.0
// detta för att kunna så regioner och kontor med olika mycket omsättning
func GetSegmentv5(i float64, MAX float64, distribution []float64) int {

	segment := 0
	value := float64(i)
	_max := float64(MAX)
	percent := value / _max
	if (percent > float64(0.0)) && (percent <= distribution[0]) {
		segment = 0
	}
	if (percent > distribution[0]) && (percent <= distribution[1]) {
		segment = 1
	}
	if (percent > distribution[1]) && (percent <= distribution[2]) {
		segment = 2
	}
	if (percent > distribution[2]) && (percent <= distribution[3]) {
		segment = 3
	}
	//fmt.Println(fmt.Sprintf("value: %d - %d - %f (%d)", max, i, percent, segment))
	return segment
}
