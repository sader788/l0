package cache

import (
	"WildberriesL0/server/internal/order"
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type OrderCache struct {
	l      *logrus.Logger
	client *pgx.Conn
	cache  map[string]order.Order
}

func CacheInit(client *pgx.Conn, l *logrus.Logger) (*OrderCache, error) {
	orderCache := OrderCache{
		l:      l,
		client: client,
		cache:  make(map[string]order.Order),
	}

	rows, _ := client.Query(context.Background(), "select * from orders")

	for rows.Next() {
		var orderUUID string
		var orderBytes []byte
		var orderJson order.Order

		err := rows.Scan(&orderUUID, &orderBytes)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(orderBytes, &orderJson)
		if err != nil {
			return nil, err
		}

		orderCache.cache[orderUUID] = orderJson
	}

	return &orderCache, nil
}

func (orders *OrderCache) GetOrderByUUID(orderUUID string) *order.Order {
	order, found := orders.cache[orderUUID]
	if found == false {
		return nil
	}
	return &order
}

func (orders *OrderCache) AppendOrder(orderJson *order.Order) error {
	orderBytes, err := json.Marshal(orderJson)
	if err != nil {
		return err
	}
	orders.client.Exec(context.Background(), "insert into orders values($1, $2)", orderJson.OrderUID, orderBytes)
	orders.cache[orderJson.OrderUID] = *orderJson

	return nil
}

func (orders *OrderCache) CacheLen() int {
	return len(orders.cache)
}
