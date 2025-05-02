package repo

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/TheTeemka/hhChat/pkg/validator"
)

type Filter struct {
	Name        string
	Surname     string
	Age         string
	Gender      string
	Nationality string
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	limit  string
	offset string
}

func NewFilters() *Filter {
	return &Filter{}
}

func (f *Filter) ParseURL(vals url.Values, v *validator.Validator) {
	f.Name = vals.Get("name")
	f.Surname = vals.Get("surname")
	f.Gender = vals.Get("gender")
	f.Nationality = vals.Get("nationality")

	f.Age = vals.Get("age")
	if f.Age != "" {
		v.CheckWithRules("age", f.Age, validator.IsInt(0))
	}

	f.limit = vals.Get("limit")
	if f.limit != "" {
		v.CheckWithRules("limit", f.limit, validator.IsInt(0))
	}

	f.offset = vals.Get("offset")
	if f.offset != "" {
		v.CheckWithRules("offset", f.offset, validator.IsInt(0))
	}
}

func (f *Filter) String() string {
	var builder builder
	if f.Name != "" {
		builder.AddWhere("name", f.Name)
	}
	if f.Surname != "" {
		builder.AddWhere("surname", f.Surname)
	}
	if f.Age != "" {
		builder.AddWhere("age", f.Age)
	}
	if f.Gender != "" {
		builder.AddWhere("gender", f.Gender)
	}
	if f.Nationality != "" {
		builder.AddWhere("nationality", f.Nationality)
	}

	if f.limit != "" {
		builder.AddLimit(f.limit)
	}
	if f.offset != "" {
		builder.AddOffset(f.offset)
	}

	return builder.String()
}

type builder struct {
	where  strings.Builder
	limit  strings.Builder
	offset strings.Builder
}

func (b *builder) AddWhere(key, value string) {
	if b.where.Len() == 0 {
		b.where.WriteString("WHERE ")
	} else {
		b.where.WriteString("AND ")
	}
	b.where.WriteString(fmt.Sprintf("%s = '%s' ", key, value))
}

func (b *builder) AddLimit(value string) {
	if b.limit.Len() > 0 {
		panic("Only one limit can be")
	}
	b.limit.WriteString("LIMIT " + value + " ")
}

func (b *builder) AddOffset(value string) {
	if b.offset.Len() > 0 {
		panic("Only one offset can be")
	}
	b.limit.WriteString("OFFSET " + value + " ")
}

func (b *builder) String() string {
	return " " + b.where.String() + b.limit.String() + b.offset.String()
}
