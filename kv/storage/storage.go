package storage

type Ctx interface{}

type Storage interface {
	Start() error
	Stop() error
	Write(ctx *Ctx, batch []any) error
	Reader(ctx *Ctx) (StorageReader, error)
}

type StorageReader interface {
	GetCF(cf string, key []byte) ([]byte, error)
	IterCF(cf string) any // should return an interator type
	Close()
}
