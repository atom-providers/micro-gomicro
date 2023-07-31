package microGoMicro

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	"go-micro.dev/v4"
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

	return container.Container.Provide(func() (contracts.Service, error) {
		service := &Service{
			conf:   &config,
			Engine: micro.NewService(),
		}
		container.AddCloseAble(service.Close)
		return service, nil
	}, o.DiOptions()...)
}

type Service struct {
	conf   *Config
	Engine micro.Service
}

func (s *Service) Serve() error {
	return s.Engine.Run()
}

func (s *Service) Close() {
	s.Engine.Server().Stop()
}

func (s *Service) GetEngine() any {
	return s.Engine
}

func (s *Service) Init(f func()) {
	s.Engine.Init(
		micro.Name("abc"),
	)
}
