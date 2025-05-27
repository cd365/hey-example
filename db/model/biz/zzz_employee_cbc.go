// code template version: v3.0.0 a1e877e692cab7668466ba74010a8e88e78e039e 1748326418-20250527141338
// TEMPLATE CODE DO NOT EDIT IT.

package biz

import (
	"database/sql"
	"github.com/cd365/hey-example/db/abc"
	"github.com/cd365/hey-example/db/model"
	"github.com/cd365/hey/v3"
)

type Employee interface {
	abc.DatabaseTable
	F(filters ...hey.Filter) hey.Filter
	Model() *model.S000001Employee
	SelectTableColumn(table *hey.TableColumn, columns ...string) []string
	Transaction(way *hey.Way, transaction func(tx *hey.Way) error, opts ...*sql.TxOptions) error
}

type employee struct {
	abc.DatabaseTable
	table *model.S000001Employee
}

func NewEmployee(table *model.S000001Employee) Employee {
	return &employee{
		DatabaseTable: table,
		table:         table,
	}
}

func (s *employee) Way(ways ...*hey.Way) *hey.Way {
	return s.table.Way(ways...)
}

func (s *employee) F(filters ...hey.Filter) hey.Filter {
	return s.Way().F(filters...)
}

func (s *employee) Model() *model.S000001Employee {
	return s.table
}

func (s *employee) SelectTableColumn(table *hey.TableColumn, columns ...string) []string {
	if table == nil {
		table = s.table.Way().T()
	}
	return table.ColumnAll(columns...)
}

func (s *employee) Transaction(way *hey.Way, transaction func(tx *hey.Way) error, opts ...*sql.TxOptions) error {
	if transaction == nil {
		return nil
	}
	if way != nil {
		return way.Transaction(nil, transaction, opts...)
	}
	return s.table.Way().Transaction(nil, transaction, opts...)
}
