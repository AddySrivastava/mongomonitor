package events

type IEventHandler interface {
	Handle(payload []byte) (bool, error)
}
