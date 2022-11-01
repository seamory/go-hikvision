package emitter

type listener func(args ...interface{})

type Callback listener

type EmitterContext struct {
	listeners map[string][]listener
}

func New() *EmitterContext {
	return &EmitterContext{
		listeners: make(map[string][]listener),
	}
}

func (x *EmitterContext) On(eventName string, cb Callback) {
	list, ok := x.listeners[eventName]
	if !ok {

	}
}

func (x *EmitterContext) Off(eventName string, cb Callback) {

}

func (x *EmitterContext) OffAll(eventName string) {

}
