package handler

import (
	"context"
	"net/http"

	"github.com/environment-toolkit/actuator/data"

	"go.uber.org/zap"

	"github.com/go-apis/eventsourcing/es"
	_ "github.com/go-apis/eventsourcing/es/providers/data/pg"
	_ "github.com/go-apis/eventsourcing/es/providers/data/sqlite"
	_ "github.com/go-apis/eventsourcing/es/providers/stream/apub"
	_ "github.com/go-apis/eventsourcing/es/providers/stream/mpub"
	_ "github.com/go-apis/eventsourcing/es/providers/stream/noop"
	_ "github.com/go-apis/eventsourcing/es/providers/stream/npub"

	"github.com/go-apis/utils/xlog"
	"github.com/go-apis/utils/xopenapi"
	"github.com/go-apis/utils/xservice"
)

func TrySet(cfg *es.ProviderConfig, pubsub es.MemoryBusPubSub) {
	if cfg.Stream.Memory == nil || pubsub == nil {
		return
	}

	cfg.Stream.Memory.PubSub = pubsub
}

type Config struct {
}

// NewHandler creates a new http handler
func NewHandler(ctx context.Context, svc *xservice.Service, pubsub es.MemoryBusPubSub) (http.Handler, error) {
	log := xlog.Logger(ctx)

	appcfg := &Config{}
	if err := svc.Parse(appcfg); err != nil {
		log.Error("failed to parse config", zap.Error(err))
		return nil, err
	}

	pcfg := &es.ProviderConfig{}
	if err := svc.Parse(pcfg); err != nil {
		return nil, err
	}

	TrySet(pcfg, pubsub)

	cli, err := data.NewClient(ctx, pcfg)
	if err != nil {
		return nil, err
	}

	s := xopenapi.New(svc.ServiceConfig)

	// Setup middlewares.
	s.Wrap(
		es.CreateUnit(cli),
	)

	// todo maybe a way to view what's running?

	return s, nil
}
