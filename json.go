package mylucky

import "reflect"

// JsonProcessor one of Processor implement
type JsonProcessor struct {
	bigEndian bool
	enc Encryptor
	msgTypes map[reflect.Type]int
	handlers map[int]msgInfo
}

// JsonProtocol is the protocol for JsonProcessor
type JsonProtocol struct {
	Id 	int	`json:"id"`
	Content interface{} `json:"content"`
}

// NewJSONProcessor return new JsonProcessor
func NewJSONProcessor() *JsonProcessor {
	pb := JsonProcessor {
		msgTypes: make(map[reflect.Type]int),
		handlers: make(map[int]msgInfo),
	}
	return &pb
}

func (jp *JsonProcessor) OnReceive


















