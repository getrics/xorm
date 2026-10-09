package core

import (
	"github.com/harluo/prometheus"
	"github.com/harluo/xorm"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Metrics struct {
	// 无字段
}

func newMetrics(
	selector *xorm.Selector, registry *prometheus.Registry,
) (metrics *Metrics, err error) {
	metrics = new(Metrics)
	selector.Foreach(func(name string, engine *xorm.Engine) {
		err = registry.Register(collectors.NewDBStatsCollector(engine.DB().DB, name))
	})

	return
}
