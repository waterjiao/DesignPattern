package metrics

import "sync"

type Memory struct {
	sync.Mutex
	infos map[string][]*RequestInfo
}

func NewMemory() *Memory {
	return &Memory{
		Mutex: sync.Mutex{},
		infos: make(map[string][]*RequestInfo),
	}
}

func (s *Memory) SaveRequestInfo(info *RequestInfo) {
	s.Lock()
	defer s.Unlock()

	s.infos[info.apiName] = append(s.infos[info.apiName], info)
}

func (s *Memory) GetRequestInfosByName(apiName string, startTimeInMills, endTimeInMills int64) []*RequestInfo {
	s.Lock()
	defer s.Unlock()

	infos := s.infos[apiName]
	return infos
}

func (s *Memory) GetRequestInfos(startTimeInMills, endTimeInMills int64) map[string][]*RequestInfo {
	panic("implement me")
}
