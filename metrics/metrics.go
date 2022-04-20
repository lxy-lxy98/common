package metrics

import (
	"fmt"
	"sort"
	"time"

	metrics "github.com/armon/go-metrics"
	"github.com/spf13/cast"
)

type metricsFunc func(key []string, val float32, labels []metrics.Label)

// MapToMetricsLables map转换成Labels
func MapToMetricsLables(maps map[string]interface{}) []metrics.Label {

	labels := make([]metrics.Label, len(maps))
	index := 0
	for k, v := range maps {
		labels[index].Name = k
		sv, err := cast.ToStringE(v)
		if err != nil {
			labels[index].Value = fmt.Sprintf("%v", v)
		} else {
			labels[index].Value = sv
		}
		index++
	}

	sort.Slice(labels[:], func(i, j int) bool {
		return labels[i].Name < labels[j].Name
	})

	return labels
}

// PairsToMetricsLables 通过传入的kv返回一组标签
// len(kv)必须为偶数，传入参数格式为key, value交替出现，且key类型必须为string
func PairsToMetricsLables(kv ...interface{}) []metrics.Label {
	if len(kv)%2 == 1 {
		panic(fmt.Sprintf("metrics: Pairs got the odd number of input pairs for metrics: %d", len(kv)))
	}
	maps := map[string]interface{}{}
	var key string
	for i, s := range kv {
		if i%2 == 0 {
			key = s.(string)
			continue
		}
		maps[key] = s
	}
	return MapToMetricsLables(maps)
}

// MetricsMilliSecondsCost 上报某个函数执行的耗时
// 调用方式
// func fake() {
// 	defer MetricsMilliSecondsCost(metrics.AddSampleWithLabels, []string{"searchManager"}, "appId", zhst.AppID_AppID_HUMMINGBIRD, "alg", zhst.AlgorithmVersion_VERSION_BNN_PRO_ATTR_SCORE)()
// 	// cost operation
// }
func MetricsMilliSecondsCost(metricscb metricsFunc, key []string, kv ...interface{}) func() {
	start := time.Now()
	return func() {
		metricscb(key, float32(time.Since(start).Milliseconds()), PairsToMetricsLables(kv))
	}
}
