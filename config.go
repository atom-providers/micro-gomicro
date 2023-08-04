package microGoMicro

import (
	"go-micro.dev/v4"
	"go.uber.org/dig"
)

const DefaultPrefix = "GoMicro"

var (
	GroupGoMicroOptions  = dig.Group("go-micro-options")
	GroupGoMicroHandlers = dig.Group("go-micro-handlers")
)

type MicroOptions struct {
	dig.In
	Options []micro.Option `group:"go-micro-options"`
}

type Config struct {
	Port uint
	Tls  *Tls
}

type Tls struct {
	Cert string
	Key  string
}
