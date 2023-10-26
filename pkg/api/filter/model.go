package filter

import "fmt"

const (
	DataTypeStr  = "string"
	DataTypeInt  = "int"
	DataTypeBool = "bool"
	DataTypeDate = "date"

	OperatorEq            = "eq"
	OperatorNotEq         = "neq"
	OperatorLowerThan     = "lt"
	OperatorLowerThanEq   = "lte"
	OperatorGreaterThan   = "gt"
	OperatorGreaterThanEq = "gte"
	OperatorBetween       = "between"
	OperatorLike          = "like"
)

type options struct {
	limit  int
	fields []Field
}

func NewOption(limit int) Options {
	return &options{limit: limit}
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

func (o *options) Limit() int {
	return o.limit
}

func (o *options) AddField(name, operator, value, dtype string) error {
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
func (o *options) Fields() []Field {
	return o.fields
}
func validateOperator(raw string) error {
	switch raw {

	case OperatorEq:
	case OperatorNotEq:
	case OperatorLowerThan:
	case OperatorLowerThanEq:
	case OperatorGreaterThan:
	case OperatorGreaterThanEq:
	default:
		return fmt.Errorf("bad operator")
	}
	return nil
}
