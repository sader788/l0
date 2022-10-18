package natsstreaming

import (
	"WildberriesL0/server/internal/config"
	"github.com/nats-io/stan.go"
	"strconv"
)

type natsManager struct {
	sc  stan.Conn
	sub stan.Subscription
}

func NewNatsManager() *natsManager {
	return &natsManager{}
}

func (nm *natsManager) Register(cfg *config.ConfigNats, msgHandler stan.MsgHandler) error {
	var err error

	nm.sc, err = stan.Connect(cfg.ClusterID, cfg.ClusterID, stan.NatsURL("nats://"+cfg.NatsHost+":"+strconv.Itoa(cfg.NatsPort)))
	if err != nil {
		return err
	}

	nm.sub, err = nm.sc.Subscribe(cfg.SubjectName, msgHandler, stan.StartWithLastReceived())
	if err != nil {
		return err
	}
	return nil
}

func (nm *natsManager) Unregister() {
	nm.sub.Unsubscribe()
	nm.sc.Close()
}
