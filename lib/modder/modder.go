package modder

type Modder interface {
	Init(module string) error
	Graph() error
	Tidy() error
	Vendor() error
	Verify() error
	Load(dir string) error
}
