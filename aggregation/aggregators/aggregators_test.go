package aggregators

import "testing"
import "github.com/rajch/learngo/aggregation/aggregators/intaggregators"

func TestSimpleAggregate(t *testing.T) {
	total, cnt := SimpleAggregate(intaggregators.IntSumAggregator{}, []int{1, 2, 3, 4, 5})
	if total != 15 || cnt != 5 {
		t.Error("Total not calculated properly")
	}
}

func TestPartionedAggregate(t *testing.T) {
	t.Error("Test not written")
}
