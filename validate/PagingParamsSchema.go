package validate

import "github.com/winstarshl/pip-services3-commons-go-vgo/convert"

func NewPagingParamsSchema() *ObjectSchema {
	return NewObjectSchema().
		WithOptionalProperty("skip", convert.Long).
		WithOptionalProperty("take", convert.Long).
		WithOptionalProperty("total", convert.Boolean)
}
