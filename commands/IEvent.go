package commands

import "github.com/winstarshl/pip-services3-commons-go-vgo/run"

type IEvent interface {
	run.INotifiable

	Name() string
	Listeners() []IEventListener
	AddListener(listener IEventListener)
	RemoveListener(listener IEventListener)
}
