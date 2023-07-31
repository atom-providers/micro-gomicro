package microGoMicro

const DefaultPrefix = "GoMicro"

type Config struct {
	Name    string
	Version string
	Port    uint
	Tls     *Tls
}

type Tls struct {
	Cert string
	Key  string
}
