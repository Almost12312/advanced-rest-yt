package model

import (
	"advanced-rest-yt/internal/author/storage"
	"fmt"
)

type sortOptions struct {
	Field, Order string
}

func NewSortOptions(field, order string) storage.SortOptions {
	return &sortOptions{
		Field: field,
		Order: order,
	}
}

func (o *sortOptions) GetOrderBy() string {
	return fmt.Sprintf("%s %s", o.Field, o.Order)
}
