package aggregators

// IntAggregator defines operations of any kind of aggregator that aggregates int values
type IntAggregator interface {
	Reset()
	Add(elem int)
	Aggregate() (int, int)
}

// SimpleAggregate calculates an aggregate of a slice of ints
func SimpleAggregate(ag IntAggregator, numbers []int) (result int, elementcount int) {
	return
}

// PartitionedAggregate calculates aggregates by slicing the input set into partitions
func PartitionedAggregate(ag IntAggregator, partitions int, numbers []int) (result int, elementcount int) {
	return
}
