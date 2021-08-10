package metrics

import (
	"sync"
	"testing"
)

func TestStartRepeatedReport(t *testing.T) {
	m := &Metrics{
		Mutex: sync.Mutex{},
		responseTimes: map[string][]int64{
			"1": []int64{1,2},
		},
		timestamps: map[string][]int64{
			"1": []int64{1},
		},
		apiNames: map[string]bool{"1": true},
	}
	m.startRepeatedReport()
}
