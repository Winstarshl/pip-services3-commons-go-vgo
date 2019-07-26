package commands

import "github.com/winstarshl/pip-services3-commons-go-vgo/run"

type IEventListener interface {
	OnEvent(correlationId string, e IEvent, value *run.Parameters)
}
