package metrics

import (
	"fmt"
	"sync"
)

type Metrics struct {
	sync.Mutex
	responseTimes map[string][]int64
	timestamps    map[string][]int64
	apiNames      map[string]bool
}

func (s *Metrics) recordResponseTime(apiName string, responseTime int64) {
	s.Lock()
	defer s.Unlock()

	if !s.apiNames[apiName] {
		s.apiNames[apiName] = true
	}
	s.responseTimes[apiName] = append(s.responseTimes[apiName], responseTime)
}

func (s *Metrics) recordTimestamp(apiName string, responseTime int64) {
	s.Lock()
	defer s.Unlock()

	s.timestamps[apiName] = append(s.timestamps[apiName], responseTime)
}

func (s *Metrics) startRepeatedReport() {
	stats := make(map[string]map[string]int64)
	for apiName, _ := range s.apiNames {
		stats[apiName] = make(map[string]int64)
	}
	for apiName, apiRespTimes := range s.responseTimes {
		stats[apiName]["max"] = max(apiRespTimes)
		stats[apiName]["avg"] = avg(apiRespTimes)
	}
	for apiName, apiTimestamps := range s.timestamps {
		stats[apiName]["count"] = int64(len(apiTimestamps))
	}
	fmt.Println(stats)
}

func max(dataset []int64) int64 {
	if len(dataset) <= 0 {
		return 0
	}
	m := dataset[0]
	for _, data := range dataset {
		if data > m {
			m = data
		}
	}
	return m
}

func avg(dataset []int64) int64 {
	if len(dataset) <= 0 {
		return 0
	}
	var sum int64
	for _, data := range dataset {
		sum = sum + data
	}
	return sum / int64(len(dataset))
}
