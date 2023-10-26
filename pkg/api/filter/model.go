package filter

const (
	DataTypeStr  = "string"
	DataTypeInt  = "int"
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
	isToApply bool
	limit     int
	fields    []Field
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
	IsToApply() bool
	AddField(name, operator, value, dtype string)
	Fields() []Field
}

func (o *options) Limit() int {
	return o.limit
}

func (o *options) IsToApply() bool {
	return o.isToApply
}
func (o *options) AddField(name, operator, value, dtype string) {
	o.fields = append(o.fields, Field{
		Name:     name,
		Operator: operator,
		Value:    value,
		Type:     dtype,
	})
}
func (o *options) Fields() []Field {
	return o.fields
}
