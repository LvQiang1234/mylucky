package mylucky

type Encryptor interface {
	Encode(bs []byte) []byte
	Decode(bs []byte) []byte
}

type Processor interface {
	SetBigEndian()
	GetBigEndian() bool
	SetEncryptor(enc Encryptor)
	OnReceivedPackage(interface{}, []byte) error
	WrapMsg(interface{}) ([]byte, error)
	RegisterHandler(id int, entity interface{}, handle func(args ...interface{}))
}
