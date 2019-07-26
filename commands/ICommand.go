package commands

import (
	"github.com/winstarshl/pip-services3-commons-go-vgo/run"
	"github.com/winstarshl/pip-services3-commons-go-vgo/validate"
)

type ICommand interface {
	run.IExecutable

	Name() string
	Validate(args *run.Parameters) []*validate.ValidationResult
}
