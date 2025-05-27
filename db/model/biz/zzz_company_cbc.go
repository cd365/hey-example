// code template version: v3.0.0 a1e877e692cab7668466ba74010a8e88e78e039e 1748326418-20250527141338
// TEMPLATE CODE DO NOT EDIT IT.

package biz

import (
	"database/sql"
	"github.com/cd365/hey-example/db/abc"
	"github.com/cd365/hey-example/db/model"
	"github.com/cd365/hey/v3"
)

type Company interface {
	abc.DatabaseTable
	F(filters ...hey.Filter) hey.Filter
	Model() *model.S000001Company
	SelectTableColumn(table *hey.TableColumn, columns ...string) []string
	Transaction(way *hey.Way, transaction func(tx *hey.Way) error, opts ...*sql.TxOptions) error
}

type company struct {
	abc.DatabaseTable
	table *model.S000001Company
}

func NewCompany(table *model.S000001Company) Company {
	return &company{
		DatabaseTable: table,
		table:         table,
	}
}

func (s *company) Way(ways ...*hey.Way) *hey.Way {
	return s.table.Way(ways...)
}

func (s *company) F(filters ...hey.Filter) hey.Filter {
	return s.Way().F(filters...)
}

func (s *company) Model() *model.S000001Company {
	return s.table
}

func (s *company) SelectTableColumn(table *hey.TableColumn, columns ...string) []string {
	if table == nil {
		table = s.table.Way().T()
	}
	return table.ColumnAll(columns...)
}

func (s *company) Transaction(way *hey.Way, transaction func(tx *hey.Way) error, opts ...*sql.TxOptions) error {
	if transaction == nil {
		return nil
	}
	if way != nil {
		return way.Transaction(nil, transaction, opts...)
	}
	return s.table.Way().Transaction(nil, transaction, opts...)
}
