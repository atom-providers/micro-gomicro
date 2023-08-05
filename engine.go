package microGoMicro

import (
	"context"
	"fmt"
	"time"

	gClient "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/logger/zap"
	gServer "github.com/go-micro/plugins/v4/server/grpc"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
	"github.com/samber/lo"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
)

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	return container.Container.Provide(func(ctx context.Context, opts MicroOptions) (contracts.MicroService, error) {
		logger, _ := zap.NewLogger(
			zap.WithLogger(opts.Log.Logger),
		)

		generateUUID := opts.Uuid.MustGenerate()
		serverOptions := []server.Option{
			server.Context(ctx),
			server.Name(atom.AppName),
			server.Version(atom.AppVersion),
			server.Id(generateUUID),
			server.Registry(opts.Registry),
		}
		if config.Port > 0 {
			addr := fmt.Sprintf(":%d", config.Port)
			serverOptions = append(serverOptions, server.Address(addr))
		}

		defaultOptions := []micro.Option{
			micro.Name(atom.AppName),
			micro.Version(atom.AppVersion),
			micro.Context(ctx),
			micro.Logger(logger),
			micro.Registry(opts.Registry),
			micro.RegisterTTL(time.Second * 30),
			micro.RegisterInterval(time.Second * 15),
			micro.Server(gServer.NewServer(serverOptions...)),
			micro.Client(gClient.NewClient()),
		}

		opts.Options = append(defaultOptions, opts.Options...)
		service := &Service{
			conf:   &config,
			ctx:    ctx,
			Engine: micro.NewService(opts.Options...),
		}
		container.AddCloseAble(service.Close)
		return service, nil
	}, o.DiOptions()...)
}

type Service struct {
	ctx    context.Context
	conf   *Config
	Engine micro.Service
}

func (s *Service) Serve() error {
	if err := s.Engine.Server().Start(); err != nil {
		return err
	}
	<-s.ctx.Done()
	return nil
}

func (s *Service) Close() {
	lo.Must0(s.Engine.Server().Stop())
}

func (s *Service) GetEngine() any {
	return s.Engine
}
