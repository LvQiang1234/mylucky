package mylucky

import (
	"mylucky/mylog"
	"reflect"
	"runtime/debug"
	"time"
)

const (
	Msg = iota
	Conn
	Raw
)

type msgInfo struct {
	msgId int
	msgType reflect.Type
	msgCallBack func(args ...interface{})
}

func execute(mInfo msgInfo, msg interface{}, writer interface{}, body []byte, id uint32) {
	defer func() {
		if r := recover(); r != nil {
			mylog.Error("%v", r)
			mylog.Error("panic at msg %d handler, stack %s", id, string(debug.Stack()))
		}
	}()
	begin := time.Now().UnixNano() / int64(time.Millisecond)
	mInfo.msgCallBack(msg, writer, body)
	costs := time.Now().UnixNano() / int64(time.Millisecond) - begin
	mylog.Debug("===> execute logic %d costs %dms, msgType %v <===", mInfo.msgId, costs, mInfo.msgType)
}