package config

import (
	"fmt"

	"github.com/anesthetised/toolkit/validate"
)

type Option[T any] func(v *T)

type Config[T validate.Validator] struct {
	val T
}

func New[T validate.Validator](opts ...Option[T]) *Config[T] {
	var zero T

	for _, opt := range opts {
		opt(&zero)
	}

	return &Config[T]{val: zero}
}

func (c *Config[T]) Val() T {
	return c.val
}

func (c *Config[T]) Ref() *T {
	ref := c.val
	return &ref
}

func (c *Config[T]) Validate() error {
	ref := c.Ref()
	validator, ok := any(ref).(validate.Validator)
	if !ok {
		return fmt.Errorf("type %T does not implement validate.Validator interface", ref)
	}
	return validator.Validate()
}
