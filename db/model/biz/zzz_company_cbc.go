// code template version: v3.0.0 876382ccafbc7ec905331e01d9c66afa58a11d6b 1744869629-20250417140029
// TEMPLATE CODE DO NOT EDIT IT.

package biz

import (
	"github.com/cd365/hey-example/db/model"
	"github.com/cd365/hey/v3"
)

type Company interface {
	Model() *model.S000001Company
	Debugger(cmder hey.Cmder) Company
	Filter(filters ...hey.Filter) func(f hey.Filter)
	SelectColumn(columns ...string) func(queryColumns hey.QueryColumns)
	SelectColumnCmder(custom func(f hey.Filter, g *hey.Get), columns ...string) hey.Cmder
	SelectTableColumn(table *hey.TableColumn, columns ...string) []string
	Transaction(way *hey.Way, transaction func(tx *hey.Way) error) error
	Insert(way *hey.Way, insert interface{}, custom ...func(add *hey.Add)) (int64, error)
	InsertOne(way *hey.Way, insert interface{}, custom ...func(add *hey.Add)) (int64, error)
	Delete(way *hey.Way, where func(f hey.Filter)) (int64, error)
	Update(way *hey.Way, update func(f hey.Filter, u *hey.Mod)) (int64, error)
	DeleteInsert(way *hey.Way, where func(f hey.Filter), insert interface{}, custom ...func(add *hey.Add)) error
	DeleteInsertOne(way *hey.Way, where func(f hey.Filter), insert interface{}, custom ...func(add *hey.Add)) (int64, error)
	Upsert(way *hey.Way, where func(f hey.Filter), update func(u *hey.Mod), insert interface{}, custom ...func(add *hey.Add)) (int64, error)
	UpsertOne(way *hey.Way, where func(f hey.Filter), update func(u *hey.Mod), insert interface{}, custom ...func(add *hey.Add)) (int64, error)
	Exists(way *hey.Way, where func(f hey.Filter), custom func(get *hey.Get)) (exists bool, err error)
	ScanOne(way *hey.Way, query func(f hey.Filter, g *hey.Get)) (*model.Company, error)
	ScanAll(way *hey.Way, query func(f hey.Filter, g *hey.Get)) ([]*model.Company, error)
	SelectOne(way *hey.Way, query func(f hey.Filter, g *hey.Get)) (*model.Company, error)
	SelectAll(way *hey.Way, query func(f hey.Filter, g *hey.Get)) ([]*model.Company, error)
	SelectGet(way *hey.Way, query func(f hey.Filter, g *hey.Get), receive interface{}) error
}

type company struct {
	table *model.S000001Company
}

func NewCompany(table *model.S000001Company) Company {
	return &company{
		table: table,
	}
}

func (s *company) Model() *model.S000001Company {
	return s.table
}

func (s *company) Debugger(cmder hey.Cmder) Company {
	s.table.Way().Debugger(cmder)
	return s
}

func (s *company) Filter(filters ...hey.Filter) func(f hey.Filter) {
	return func(f hey.Filter) { f.Use(filters...) }
}

func (s *company) SelectColumn(columns ...string) func(queryColumns hey.QueryColumns) {
	return func(queryColumns hey.QueryColumns) { queryColumns.AddAll(columns...) }
}

func (s *company) SelectColumnCmder(custom func(f hey.Filter, g *hey.Get), columns ...string) hey.Cmder {
	m := s.Model()
	where := m.Filter()
	result := m.Get().Select(columns...)
	if custom != nil {
		custom(where, result)
	}
	return result.Where(func(f hey.Filter) { f.Use(where) })
}

func (s *company) SelectTableColumn(table *hey.TableColumn, columns ...string) []string {
	if table == nil {
		table = s.table.Way().T()
	}
	return table.ColumnAll(columns...)
}

func (s *company) Transaction(way *hey.Way, transaction func(tx *hey.Way) error) error {
	if transaction == nil {
		return nil
	}
	return s.table.Way(way).Transaction(nil, transaction)
}

func (s *company) Insert(way *hey.Way, insert interface{}, custom ...func(add *hey.Add)) (int64, error) {
	if insert == nil {
		return 0, nil
	}
	add := s.table.Add(way)
	for _, tmp := range custom {
		if tmp != nil {
			tmp(add)
		}
	}
	return add.Create(insert).Add()
}

func (s *company) InsertOne(way *hey.Way, insert interface{}, custom ...func(add *hey.Add)) (int64, error) {
	if insert == nil {
		return 0, nil
	}
	return s.table.AddOne(func(add *hey.Add) {
		for _, tmp := range custom {
			if tmp != nil {
				tmp(add)
			}
		}
	}, insert, way)
}

func (s *company) Delete(way *hey.Way, where func(f hey.Filter)) (int64, error) {
	if where == nil {
		return 0, nil
	}
	filter := s.table.Filter()
	where(filter)
	return s.table.Delete(filter, way)
}

func (s *company) Update(way *hey.Way, update func(f hey.Filter, u *hey.Mod)) (int64, error) {
	if update == nil {
		return 0, nil
	}
	return s.table.Update(update, way)
}

func (s *company) DeleteInsert(way *hey.Way, where func(f hey.Filter), insert interface{}, custom ...func(add *hey.Add)) error {
	if where != nil {
		filter := s.table.Filter()
		where(filter)
		exists, err := s.table.SelectExists(filter, nil, way)
		if err != nil {
			return err
		}
		if exists {
			if _, err = s.table.Delete(filter, way); err != nil {
				return err
			}
		}
	}
	if _, err := s.Insert(way, insert, custom...); err != nil {
		return err
	}
	return nil
}

func (s *company) DeleteInsertOne(way *hey.Way, where func(f hey.Filter), insert interface{}, custom ...func(add *hey.Add)) (int64, error) {
	if insert == nil {
		return 0, nil
	}
	if where != nil {
		filter := s.table.Filter()
		where(filter)
		exists, err := s.table.SelectExists(filter, nil, way)
		if err != nil {
			return 0, err
		}
		if exists {
			if _, err = s.table.Delete(filter, way); err != nil {
				return 0, err
			}
		}
	}
	return s.InsertOne(way, insert, custom...)
}

func (s *company) Upsert(way *hey.Way, where func(f hey.Filter), update func(u *hey.Mod), insert interface{}, custom ...func(add *hey.Add)) (int64, error) {
	filter := s.table.Filter()
	if where != nil {
		where(filter)
	}
	exists, err := s.Exists(way, s.Filter(filter), nil)
	if err != nil {
		return 0, err
	}
	if exists {
		return s.Update(way, func(f hey.Filter, u *hey.Mod) {
			f.Use(filter)
			if update != nil {
				update(u)
			}
		})
	}
	return s.Insert(way, insert, func(add *hey.Add) {
		for _, tmp := range custom {
			if tmp != nil {
				tmp(add)
			}
		}
	})
}

func (s *company) UpsertOne(way *hey.Way, where func(f hey.Filter), update func(u *hey.Mod), insert interface{}, custom ...func(add *hey.Add)) (int64, error) {
	filter := s.table.Filter()
	if where != nil {
		where(filter)
	}
	exists, err := s.Exists(way, s.Filter(filter), nil)
	if err != nil {
		return 0, err
	}
	if exists {
		return s.Update(way, func(f hey.Filter, u *hey.Mod) {
			f.Use(filter)
			if update != nil {
				update(u)
			}
		})
	}
	return s.InsertOne(way, insert, func(add *hey.Add) {
		for _, tmp := range custom {
			if tmp != nil {
				tmp(add)
			}
		}
	})
}

func (s *company) Exists(way *hey.Way, where func(f hey.Filter), custom func(get *hey.Get)) (exists bool, err error) {
	filter := s.table.Filter()
	if where != nil {
		where(filter)
	}
	return s.table.SelectExists(filter, func(get *hey.Get) {
		get.Where(func(f hey.Filter) { f.Use(filter) })
		if custom != nil {
			custom(get)
		}
	}, way)
}

func (s *company) ScanOne(way *hey.Way, query func(f hey.Filter, g *hey.Get)) (*model.Company, error) {
	m := s.Model()
	where := m.Filter()
	return m.RowsScanOne(where, func(get *hey.Get) {
		if query != nil {
			query(where, get)
		}
		get.Where(func(f hey.Filter) { f.Use(where) })
	}, way)
}

func (s *company) ScanAll(way *hey.Way, query func(f hey.Filter, g *hey.Get)) ([]*model.Company, error) {
	m := s.Model()
	where := m.Filter()
	return m.RowsScanAll(where, func(get *hey.Get) {
		if query != nil {
			query(where, get)
		}
		get.Where(func(f hey.Filter) { f.Use(where) })
	}, way)
}

func (s *company) SelectOne(way *hey.Way, query func(f hey.Filter, g *hey.Get)) (*model.Company, error) {
	m := s.Model()
	where := m.Filter()
	return m.SelectOne(where, func(get *hey.Get) {
		if query != nil {
			query(where, get)
		}
		get.Where(func(f hey.Filter) { f.Use(where) })
	}, way)
}

func (s *company) SelectAll(way *hey.Way, query func(f hey.Filter, g *hey.Get)) ([]*model.Company, error) {
	m := s.Model()
	where := m.Filter()
	return m.SelectAll(where, func(get *hey.Get) {
		if query != nil {
			query(where, get)
		}
		get.Where(func(f hey.Filter) { f.Use(where) })
	}, way)
}

func (s *company) SelectGet(way *hey.Way, query func(f hey.Filter, g *hey.Get), receive interface{}) error {
	m := s.Model()
	where := m.Filter()
	return m.SelectGet(where, func(get *hey.Get) {
		if query != nil {
			query(where, get)
		}
		get.Where(func(f hey.Filter) { f.Use(where) })
	}, receive, way)
}
