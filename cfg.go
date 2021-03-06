package mylucky

import "sync"

type Data struct {
	//单个连接未处理消息包的缓存队列
	//注意:[超过这个大小，包将丢弃，视为当前系统无法处理，默认100]
	ConnUndoQueueSize int
	//单个连接写入消息包队列大小
	ConnWriteQueueSize int
	//第一个包等待超时时间(s)[默认5秒，连接上来未读到正确包，断开连接]
	FirstPackageTimeOut int
	//连接读取超时[默认35秒，超时等待时间内，请发送任何数据包，如心跳包]
	ConnReadTimeout int
	//连接写超时[默认5秒，超时等待时间内，请发送任何数据包，如心跳包]
	ConnWriteTimeout int
	// 数据包最大限制，[默认2048]
	MaxDataPackageSize int
	// ws 最大header，[默认1024]
	MaxHeaderLen int
}

var (
	C *Data
	once sync.Once
)

func init() {
	C = &Data {
		ConnUndoQueueSize: 100,
		ConnWriteQueueSize: 10,
		FirstPackageTimeOut: 5,
		ConnReadTimeout: 35,
		ConnWriteTimeout: 5,
		MaxDataPackageSize: 4096,
		MaxHeaderLen: 1024,
	}
}

func SetConf(cfg *Data) {
	once.Do(func () {
		C = cfg
		if C.ConnUndoQueueSize == 0 {
			C.ConnUndoQueueSize = 1
		}
		if C.ConnWriteQueueSize == 0 {
			C.ConnWriteQueueSize = 1
		}
	})
}







