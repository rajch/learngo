package aggregators

import "testing"
import "github.com/rajch/learngo/aggregation/aggregators/intaggregators"

func TestSimpleAggregate(t *testing.T) {
	ag := new(intaggregators.IntSumAggregator)
	total, cnt := SimpleAggregate(ag, []int{1, 2, 3, 4, 5})
	if total != 15 || cnt != 5 {
		t.Errorf("Total not calculated properly. Total was %v, count %v.", total, cnt)
	}
}

func TestPartionedAggregate(t *testing.T) {
	t.Skip("Test not written, but skipping for now.")
}
