package main

import (
	"context"

	"github.com/environment-toolkit/actuator/handler"

	"github.com/go-apis/utils/xgraceful"
	"github.com/go-apis/utils/xservice"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := xservice.NewConfig(ctx)
	if err != nil {
		panic(err)
	}

	handler, err := handler.NewHandler(ctx, cfg, nil)
	if err != nil {
		panic(err)
	}

	xgraceful.Serve(ctx, cfg, handler)
	cancel()

	<-ctx.Done()
}
