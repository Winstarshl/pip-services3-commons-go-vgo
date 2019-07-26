package validate

import "github.com/winstarshl/pip-services3-commons-go-vgo/errors"

type ISchema interface {
	Validate(value interface{}) []*ValidationResult
	ValidateAndReturnError(correlationId string, value interface{}, strict bool) *errors.ApplicationError
	ValidateAndThrowError(correlationId string, value interface{}, strict bool)
}
