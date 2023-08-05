package microGoMicro

import (
	"github.com/atom-providers/log"
	"github.com/atom-providers/uuid"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go.uber.org/dig"
)

const DefaultPrefix = "GoMicro"

var GroupGoMicroOptions = dig.Group("go-micro-options")

type MicroOptions struct {
	dig.In
	Uuid     *uuid.Generator
	Log      *log.Logger
	Registry registry.Registry
	Options  []micro.Option `group:"go-micro-options"`
}

type Config struct {
	Port uint
}
