package handlers

import (
	"WildberriesL0/server/internal/cache"
	"WildberriesL0/server/internal/order"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func StanHandler(cache *cache.OrderCache, l *logrus.Logger) func(msg *stan.Msg) {
	return func(msg *stan.Msg) {
		var orderJson order.Order

		err := json.Unmarshal(msg.Data, &orderJson)
		if err != nil {
			l.Warn("nats-streaming: received json isn`t correct")
			return
		}

		err = cache.AppendOrder(&orderJson)
		if err != nil {
			l.Warn("nats-streaming: failed to add order to cache")
			return
		}

		l.Info("nats-streaming: order received, orderUID: " + orderJson.OrderUID)
	}
}
