package commands

import (
	"github.com/winstarshl/pip-services3-commons-go-vgo/run"
	"github.com/winstarshl/pip-services3-commons-go-vgo/validate"
)

type ICommandInterceptor interface {
	Name(command ICommand) string
	Execute(correlationId string, command ICommand, args *run.Parameters) (interface{}, error)
	Validate(command ICommand, args *run.Parameters) []*validate.ValidationResult
}
