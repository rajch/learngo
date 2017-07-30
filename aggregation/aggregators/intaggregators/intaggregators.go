package intaggregators

type intAccAggregator struct {
	accumulator int
	count       int
}

// Reset resets the total to 0
func (ag *intAccAggregator) Reset() {
	ag.accumulator, ag.count = 0, 0
}

// Add adds the number to the aggregator
func (ag *intAccAggregator) Add(elem int) {
	ag.accumulator += elem
	ag.count++
}

// IntSumAggregator calculates totals
type IntSumAggregator struct {
	intAccAggregator
}

// Aggregate returns the total and the number of elements
// added to the aggregator
func (ag *IntSumAggregator) Aggregate() (int, int) {
	return ag.accumulator, ag.count
}

// IntAvgAggregator calculates averages
type IntAvgAggregator struct {
	intAccAggregator
}

// Aggregate returns the average and the number of elements
// added to the aggregator
func (ag *IntAvgAggregator) Aggregate() (int, int) {
	return int(float32(ag.accumulator) / float32(ag.count)), ag.count
}
