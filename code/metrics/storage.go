package metrics

type MetricsStorage interface {
	SaveRequestInfo(info *RequestInfo)
	GetRequestInfosByName(apiName string, startTimeInMills, endTimeInMills int64) []*RequestInfo
	GetRequestInfos(startTimeInMills, endTimeInMills int64) map[string][]*RequestInfo
}
