package intaggregators

// IntSumAggregator calculates totals
type IntSumAggregator struct {
	accumulator int
	count       int
}

// Reset resets the total to 0
func (ag *IntSumAggregator) Reset() {
	ag.accumulator, ag.count = 0, 0
}

// Add adds the number to the aggregator
func (ag *IntSumAggregator) Add(elem int) {
	ag.accumulator += elem
	ag.count++
}

// Aggregate returns the total and the number of elements
// added to the aggregator
func (ag *IntSumAggregator) Aggregate() (int, int) {
	return ag.accumulator, ag.count
}
