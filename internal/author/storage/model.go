package storage

import "fmt"

type sortOptions struct {
	Field, Order string
}

func NewSortOptions(field, order string) SortOptions {
	return &sortOptions{
		Field: field,
		Order: order,
	}
}

func (o *sortOptions) GetOrderBy() string {
	return fmt.Sprintf("%s %s", o.Field, o.Order)
}
