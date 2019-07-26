package validate

import "github.com/winstarshl/pip-services3-commons-go-vgo/convert"

func NewProjectionParamsSchema() *ArraySchema {
	return NewArraySchema(convert.String)
}
