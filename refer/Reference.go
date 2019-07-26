package refer

import "github.com/winstarshl/pip-services3-commons-go-vgo/data"

type Reference struct {
	locator   interface{}
	component interface{}
}

func NewReference(locator interface{}, component interface{}) *Reference {
	if component == nil {
		panic("Component cannot be null")
	}

	return &Reference{
		locator:   locator,
		component: component,
	}
}

func (c *Reference) Component() interface{} {
	return c.component
}

func (c *Reference) Locator() interface{} {
	return c.locator
}

func (c *Reference) Match(locator interface{}) bool {
	// Check for nil locator
	if locator == nil {
		return false
	}

	// Locate by direct reference matching
	if c.component == locator {
		return true
	}

	// Locate by direct locator matching
	equatable, ok := c.locator.(data.IEquatable)
	if ok {
		return equatable.Equals(locator)
	}

	// Locate by direct locator matching
	return c.locator == locator
}
