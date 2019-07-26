package validate

import "github.com/winstarshl/pip-services3-commons-go-vgo/convert"

func NewFilterParamsSchema() *MapSchema {
	return NewMapSchema(convert.String, nil)
}
