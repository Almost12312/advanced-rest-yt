package model

import (
	"advanced-rest-yt/pkg/api/filter"
	"fmt"
)

type filterOptions struct {
	limit  int
	fields []Field
}

func NewOption(limit int) Options {
	return &filterOptions{limit: limit}
}

type Field struct {
	Name     string
	Operator string
	Value    string
	Type     string
}

type Options interface {
	Limit() int
	AddField(name, operator, value, dtype string) error
	Fields() []Field
}

func (o *filterOptions) Limit() int {
	return o.limit
}

func (o *filterOptions) AddField(name, operator, value, dtype string) error {
	err := validateOperator(operator)
	if err != nil {
		return err
	}

	o.fields = append(o.fields, Field{
		Name:     name,
		Operator: operator,
		Value:    value,
		Type:     dtype,
	})
	return nil
}
func (o *filterOptions) Fields() []Field {
	return o.fields
}
func validateOperator(raw string) error {
	switch raw {

	case filter.OperatorEq:
	case filter.OperatorNotEq:
	case filter.OperatorLowerThan:
	case filter.OperatorLowerThanEq:
	case filter.OperatorGreaterThan:
	case filter.OperatorGreaterThanEq:
	default:
		return fmt.Errorf("bad operator")
	}
	return nil
}
