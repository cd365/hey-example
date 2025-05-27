// code template version: v3.0.0 a1e877e692cab7668466ba74010a8e88e78e039e 1748326418-20250527141338
// TEMPLATE CODE DO NOT EDIT IT.

package model

import (
	"context"
	_ "embed"
	"github.com/cd365/hey-example/db/abc"
	"github.com/cd365/hey/v3"
	"sync"
	"time"
)

//go:embed aaa_table_create.sql
var tableCreateSql []byte

type Database struct {
	schemaMap   map[string]abc.DatabaseTable
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
	tmp.schemaMap = map[string]abc.DatabaseTable{
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

func (s *Database) TableMap() map[string]abc.DatabaseTable {
	length := len(s.schemaMap)
	result := make(map[string]abc.DatabaseTable, length)
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

func (s *Database) TableCreate() []byte {
	return tableCreateSql
}

// CopyDatabase Copy all current table objects and their data to the target database.
func (s *Database) CopyDatabase(dst *hey.Way) error {
	_, resultErr := dst.GetDatabase().Exec(string(s.TableCreate()))
	if resultErr != nil {
		return resultErr
	}
	wg := &sync.WaitGroup{}
	once := &sync.Once{}
	callOnce := func(err error) {
		if err != nil {
			resultErr = err
		}
	}
	backup := func(table abc.DatabaseTable) {
		defer wg.Done()
		if tmp, ok := table.(abc.DatabaseBackup); ok {
			// TRUNCATE TABLE
			if _, err := dst.Exec(hey.ConcatString("TRUNCATE", hey.SqlSpace, "TABLE", hey.SqlSpace, dst.Replace(table.Table()))); err != nil {
				once.Do(func() { callOnce(err) })
				return
			}
			// WRITE DATA
			if err := tmp.Backup(1000, nil, func(add *hey.Add, creates interface{}) (affectedRows int64, err error) {
				return dst.Add(table.Table()).Create(creates).Add()
			}); err != nil {
				once.Do(func() { callOnce(err) })
				return
			}
		}
	}
	for _, name := range s.schemaSlice {
		if table, exists := s.schemaMap[name]; exists && table != nil {
			wg.Add(1)
			backup(table)
		}
	}
	wg.Wait()
	return resultErr
}
