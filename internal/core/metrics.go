package core

import (
	"github.com/harluo/prometheus"
	"github.com/harluo/xorm"
)

type Metrics struct {
	// 无字段
}

func newMetrics(
	selector *xorm.Selector, registry *prometheus.Registry,
) (metrics *Metrics, err error) {
	metrics = new(Metrics)
	selector.Foreach(func(name string, engine *xorm.Engine) {
		err = registry.Register(collectors.NewDBStatsCollector(db, name))
		if err != nil {
			break
		}
	})

	return
}
