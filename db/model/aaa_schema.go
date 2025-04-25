// code template version: v3.0.0 67ab087b6ba2926de886c8a05e3188b18cd6567d 1745553000-20250425115000
// TEMPLATE CODE DO NOT EDIT IT.

package model

import (
	"context"
	"github.com/cd365/hey-example/db/abc"
	"github.com/cd365/hey/v3"
	"time"
)

type Database struct {
	schemaMap   map[string]abc.Table
	schemaSlice []string

	COMPANY  *S000001Company
	EMPLOYEE *S000001Employee
}

func NewDatabase(ctx context.Context, way *hey.Way, initialize func(db *Database) error) (*Database, error) {
	basic := abc.BASIC{
		Ctx:                   ctx,
		SqlExecuteMaxDuration: time.Minute,
	}
	tmp := &Database{
		COMPANY:  newS000001Company(basic, way),
		EMPLOYEE: newS000001Employee(basic, way),
	}
	tmp.schemaMap = map[string]abc.Table{
		tmp.COMPANY.Table():  tmp.COMPANY,
		tmp.EMPLOYEE.Table(): tmp.EMPLOYEE,
	}
	tmp.schemaSlice = []string{
		tmp.COMPANY.Table(),
		tmp.EMPLOYEE.Table(),
	}
	if initialize != nil {
		if err := initialize(tmp); err != nil {
			return nil, err
		}
	}
	return tmp, nil
}

func (s *Database) TableMap() map[string]abc.Table {
	length := len(s.schemaMap)
	result := make(map[string]abc.Table, length)
	for k, v := range s.schemaMap {
		result[k] = v
	}
	return result
}

func (s *Database) TableSlice() []string {
	length := len(s.schemaSlice)
	result := make([]string, length)
	_ = copy(result, s.schemaSlice)
	return result
}

func (s *Database) TableExists(table string) bool {
	_, ok := s.schemaMap[table]
	return ok
}
