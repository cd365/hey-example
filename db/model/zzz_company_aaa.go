// code template version: v3.0.0 a1e877e692cab7668466ba74010a8e88e78e039e 1748326418-20250527141338
// TEMPLATE CODE DO NOT EDIT IT.

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/cd365/hey-example/db/abc"
	"github.com/cd365/hey/v3"
	"reflect"
	"strings"
)

// Company | company | company
type Company struct {
	Id        int    `json:"id" db:"id"`                 // id
	Pid       int    `json:"pid" db:"pid"`               // pid
	Name      string `json:"name" db:"name"`             // name
	Country   string `json:"country" db:"country"`       // country
	City      string `json:"city" db:"city"`             // city
	Address   string `json:"address" db:"address"`       // address
	Logo      string `json:"logo" db:"logo"`             // logo
	State     int    `json:"state" db:"state"`           // state
	Remark    string `json:"remark" db:"remark"`         // remark
	CreatedAt int64  `json:"created_at" db:"created_at"` // created_at
	UpdatedAt int64  `json:"updated_at" db:"updated_at"` // updated_at
	DeletedAt int64  `json:"deleted_at" db:"deleted_at"` // deleted_at
}

type S000001Company struct {
	ID         string // id
	PID        string // pid
	NAME       string // name
	COUNTRY    string // country
	CITY       string // city
	ADDRESS    string // address
	LOGO       string // logo
	STATE      string // state
	REMARK     string // remark
	CREATED_AT string // created_at
	UPDATED_AT string // updated_at
	DELETED_AT string // deleted_at

	table   string
	comment string

	border string

	columnMap   map[string]*struct{}
	columnSlice []string
	columnIndex map[string]int

	basic *abc.BASIC
	way   *hey.Way
}

func (s *S000001Company) Basic() *abc.BASIC {
	return s.basic
}

func (s *S000001Company) Table() string {
	return s.table
}

func (s *S000001Company) Comment() string {
	return s.comment
}

func (s *S000001Company) Column(except ...string) []string {
	excepted := make(map[string]*struct{}, len(except))
	for _, v := range except {
		excepted[v] = &struct{}{}
	}
	result := make([]string, 0, len(s.columnSlice))
	for _, v := range s.columnSlice {
		if _, ok := excepted[v]; ok {
			continue
		}
		result = append(result, v)
	}
	return result
}

func (s *S000001Company) ColumnMap() map[string]*struct{} {
	result := make(map[string]*struct{}, len(s.columnMap))
	for k, v := range s.columnMap {
		result[k] = v
	}
	return result
}

func (s *S000001Company) ColumnString() string {
	return `"id", "pid", "name", "country", "city", "address", "logo", "state", "remark", "created_at", "updated_at", "deleted_at"`
}

func (s *S000001Company) ColumnExist(column string) bool {
	_, exist := s.columnMap[column]
	return exist
}

func (s *S000001Company) ColumnPermit(permit ...string) []string {
	result := make([]string, 0, len(permit))
	for _, v := range permit {
		if ok := s.ColumnExist(v); ok {
			result = append(result, v)
		}
	}
	return result
}

func (s *S000001Company) ColumnValue(columnValue ...interface{}) map[string]interface{} {
	length := len(columnValue)
	if length == 0 || length&1 == 1 {
		return nil
	}
	result := make(map[string]interface{}, length)
	for i := 0; i < length; i += 2 {
		if i >= length || i+1 >= length {
			continue
		}
		column, ok := columnValue[i].(string)
		if !ok {
			continue
		}
		if ok = s.ColumnExist(column); !ok {
			continue
		}
		result[column] = columnValue[i+1]
	}
	return result
}

func (s *S000001Company) ColumnAutoIncr() []string {
	return []string{s.ID}
}

func (s *S000001Company) ColumnCreatedAt() []string {
	return []string{s.CREATED_AT}
}

func (s *S000001Company) ColumnUpdatedAt() []string {
	return []string{s.UPDATED_AT}
}

func (s *S000001Company) ColumnDeletedAt() []string {
	return []string{s.DELETED_AT}
}

func (s *S000001Company) Filter(filters ...func(f hey.Filter)) hey.Filter {
	filter := s.way.F()
	for _, tmp := range filters {
		if tmp != nil {
			tmp(filter)
		}
	}
	return filter
}

func (s *S000001Company) Way(ways ...*hey.Way) *hey.Way {
	return abc.Way(s.way, ways...)
}

func (s *S000001Company) Add(ways ...*hey.Way) *hey.Add {
	excepts := s.ColumnAutoIncr()
	return s.Way(ways...).Add(s.Table()).
		ExceptPermit(
			func(except hey.UpsertColumns, permit hey.UpsertColumns) {
				except.Add(excepts...)
				permit.Add(s.Column(excepts...)...)
			},
		)
}

func (s *S000001Company) Del(ways ...*hey.Way) *hey.Del {
	return s.Way(ways...).Del(s.Table())
}

func (s *S000001Company) Mod(ways ...*hey.Way) *hey.Mod {
	excepts := s.ColumnAutoIncr()
	excepts = append(excepts, s.ColumnCreatedAt()...)
	return s.Way(ways...).Mod(s.Table()).
		ExceptPermit(
			func(except hey.UpsertColumns, permit hey.UpsertColumns) {
				except.Add(excepts...)
				permit.Add(s.Column(excepts...)...)
			},
		)
}

func (s *S000001Company) Get(ways ...*hey.Way) *hey.Get {
	return s.Way(ways...).Get(s.Table()).Select(s.columnSlice...)
}

func (s *S000001Company) Available() hey.Filter {
	return s.Filter(func(f hey.Filter) {
		for _, v := range s.ColumnDeletedAt() {
			f.Equal(v, 0)
		}
	})
}

func (s *S000001Company) Debugger(cmder hey.Cmder) {
	s.way.Debugger(cmder)
}

// AddOne Insert a record and return the auto-increment id.
func (s *S000001Company) AddOne(create interface{}, custom func(add *hey.Add)) (int64, error) {
	if create == nil {
		return 0, nil
	}
	add := s.Add()
	if custom != nil {
		custom(add)
	}
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()

	add.Context(ctx)
	add.Default(func(o *hey.Add) {
		timestamp := o.GetWay().Now().Unix()
		for _, v := range s.ColumnCreatedAt() {
			o.ColumnValue(v, timestamp)
		}
	}).Create(create)
	return add.AddOne(func(add hey.AddOneReturnSequenceValue) {
		add.Adjust(func(prepare string, args []interface{}) (string, []interface{}) {
			return hey.ConcatString(prepare, fmt.Sprintf(" RETURNING %s", s.way.Replace(s.PrimaryKey()))), args
		}).Execute(func(ctx context.Context, stmt *hey.Stmt, args []interface{}) (id int64, err error) {
			err = stmt.QueryRowContext(ctx, func(rows *sql.Row) error { return rows.Scan(&id) }, args...)
			return
		})
	})

}

// Insert SQL INSERT.
func (s *S000001Company) Insert(create interface{}, custom func(add *hey.Add)) (int64, error) {
	if create == nil {
		return 0, nil
	}
	typeOf := reflect.TypeOf(create)
	kind := typeOf.Kind()
	for kind == reflect.Ptr {
		typeOf = typeOf.Elem()
		kind = typeOf.Kind()
	}
	if kind == reflect.Map || kind == reflect.Struct {
		return s.AddOne(create, custom)
	}
	add := s.Add()
	if custom != nil {
		custom(add)
	}
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()
	return add.Context(ctx).
		Default(func(o *hey.Add) {
			timestamp := o.GetWay().Now().Unix()
			for _, v := range s.ColumnCreatedAt() {
				o.ColumnValue(v, timestamp)
			}
		}).Create(create).Add()
}

// Delete SQL DELETE.
func (s *S000001Company) Delete(custom func(del *hey.Del, where hey.Filter)) (int64, error) {
	if custom == nil {
		return 0, nil
	}
	remove := s.Del()
	where := s.Filter()
	custom(remove, where)
	if where.IsEmpty() {
		return 0, nil
	}
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()
	return remove.Context(ctx).Where(func(f hey.Filter) { f.Use(where) }).Del()
}

// Update SQL UPDATE.
func (s *S000001Company) Update(update func(mod *hey.Mod, where hey.Filter)) (int64, error) {
	if update == nil {
		return 0, nil
	}
	filter := s.Filter()
	modify := s.Mod()
	update(modify, filter)
	if filter.IsEmpty() {
		return 0, nil
	}
	modify.Default(func(o *hey.Mod) {
		timestamp := o.GetWay().Now().Unix()
		for _, v := range s.ColumnUpdatedAt() {
			o.Set(v, timestamp)
		}
	})
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()
	return modify.Context(ctx).Where(func(f hey.Filter) { f.Use(filter) }).Mod()
}

// InsertSelect SQL INSERT and SELECT.
func (s *S000001Company) InsertSelect(columns []string, get *hey.Get, way *hey.Way) (int64, error) {
	if len(columns) == 0 || get == nil {
		return 0, nil
	}
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()
	return s.Add().SetWay(way).Context(ctx).CmderValues(get, columns).Add()
}

// SelectCount SQL SELECT COUNT.
func (s *S000001Company) SelectCount(custom func(get *hey.Get, where hey.Filter)) (int64, error) {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	return get.Select(s.columnSlice[0]).Where(func(f hey.Filter) { f.Use(where) }).Count()
}

// SelectQuery SQL SELECT.
func (s *S000001Company) SelectQuery(custom func(get *hey.Get, where hey.Filter), query func(rows *sql.Rows) error) error {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	return get.Where(func(f hey.Filter) { f.Use(where) }).Query(query)
}

// EmptySlice Initialize an empty slice.
func (s *S000001Company) EmptySlice() []*Company {
	return make([]*Company, 0)
}

// SelectGet SQL SELECT.
func (s *S000001Company) SelectGet(custom func(get *hey.Get, where hey.Filter), receive interface{}) error {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	return get.Where(func(f hey.Filter) { f.Use(where) }).Get(receive)
}

// SelectAll SQL SELECT ALL.
func (s *S000001Company) SelectAll(custom func(get *hey.Get, where hey.Filter)) ([]*Company, error) {
	lists := s.EmptySlice()
	if err := s.SelectGet(custom, &lists); err != nil {
		return nil, err
	}
	return lists, nil
}

// SelectOne SQL SELECT ONE.
func (s *S000001Company) SelectOne(custom func(get *hey.Get, where hey.Filter)) (*Company, error) {
	all, err := s.SelectAll(func(get *hey.Get, where hey.Filter) {
		if custom != nil {
			custom(get, where)
		}
		get.Limit(1)
	})
	if err != nil {
		return nil, err
	}
	if len(all) == 0 {
		return nil, hey.RecordDoesNotExists
	}
	return all[0], nil
}

// SelectExists SQL SELECT EXISTS.
func (s *S000001Company) SelectExists(custom func(get *hey.Get, where hey.Filter)) (bool, error) {
	exists, err := s.SelectOne(func(get *hey.Get, where hey.Filter) {
		if custom != nil {
			custom(get, where)
		}
		get.Select(s.columnSlice[0])
	})
	if err != nil && !errors.Is(err, hey.RecordDoesNotExists) {
		return false, err
	}
	return exists != nil, nil
}

// SelectCountAll SQL SELECT COUNT + ALL.
func (s *S000001Company) SelectCountAll(custom func(get *hey.Get, where hey.Filter)) (int64, []*Company, error) {
	count, err := s.SelectCount(custom)
	if err != nil {
		return 0, nil, err
	}
	if count == 0 {
		return 0, make([]*Company, 0), nil
	}
	all, err := s.SelectAll(custom)
	if err != nil {
		return 0, nil, err
	}
	return count, all, nil
}

// SelectCountGet SQL SELECT COUNT + GET.
func (s *S000001Company) SelectCountGet(custom func(get *hey.Get, where hey.Filter), receive interface{}) (int64, error) {
	count, err := s.SelectCount(custom)
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, nil
	}
	if err = s.SelectGet(custom, receive); err != nil {
		return 0, err
	}
	return count, nil
}

// SelectAllMap Make map[string]*Company
func (s *S000001Company) SelectAllMap(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Company) string) (map[string]*Company, []*Company, error) {
	all, err := s.SelectAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[string]*Company, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

// SelectAllMapInt Make map[int]*Company
func (s *S000001Company) SelectAllMapInt(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Company) int) (map[int]*Company, []*Company, error) {
	all, err := s.SelectAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int]*Company, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

// SelectAllMapInt64 Make map[int64]*Company
func (s *S000001Company) SelectAllMapInt64(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Company) int64) (map[int64]*Company, []*Company, error) {
	all, err := s.SelectAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int64]*Company, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

// DeleteByColumn Delete by column values. Additional conditions can be added in the filters. No transaction support.
func (s *S000001Company) DeleteByColumn(column string, values interface{}, custom func(del *hey.Del, where hey.Filter)) (int64, error) {
	return s.Delete(func(del *hey.Del, where hey.Filter) {
		where.In(column, values)
		if custom != nil {
			custom(del, where)
		}
	})
}

// UpdateByColumn Update by column values. Additional conditions can be added in the filters. No transaction support.
func (s *S000001Company) UpdateByColumn(column string, values interface{}, update interface{}, custom func(mod *hey.Mod, where hey.Filter)) (int64, error) {
	if update == nil {
		return 0, nil
	}
	return s.Update(func(mod *hey.Mod, where hey.Filter) {
		where.In(column, values)
		if custom != nil {
			custom(mod, where)
		}
		mod.Update(update)
	})
}

// SelectAllByColumn Select all by column values. No transaction support.
func (s *S000001Company) SelectAllByColumn(column string, values interface{}, customs ...func(get *hey.Get, where hey.Filter)) ([]*Company, error) {
	return s.SelectAll(func(get *hey.Get, where hey.Filter) {
		where.In(column, values)
		for _, custom := range customs {
			if custom != nil {
				custom(get, where)
				break
			}
		}
	})
}

// SelectOneByColumn Select one by column values. No transaction support.
func (s *S000001Company) SelectOneByColumn(column string, values interface{}, customs ...func(get *hey.Get, where hey.Filter)) (*Company, error) {
	return s.SelectOne(func(get *hey.Get, where hey.Filter) {
		where.In(column, values)
		for _, custom := range customs {
			if custom != nil {
				custom(get, where)
				break
			}
		}
	})
}

// SelectExistsByColumn Select exists by column values. No transaction support.
func (s *S000001Company) SelectExistsByColumn(column string, values interface{}, customs ...func(get *hey.Get, where hey.Filter)) (bool, error) {
	return s.SelectExists(func(get *hey.Get, where hey.Filter) {
		where.In(column, values)
		for _, custom := range customs {
			if custom != nil {
				custom(get, where)
				break
			}
		}
	})
}

// SelectGetByColumn Select get by column values. No transaction support.
func (s *S000001Company) SelectGetByColumn(column string, values interface{}, receive interface{}, customs ...func(get *hey.Get, where hey.Filter)) error {
	return s.SelectGet(func(get *hey.Get, where hey.Filter) {
		where.In(column, values)
		for _, custom := range customs {
			if custom != nil {
				custom(get, where)
				break
			}
		}
	}, receive)
}

// DeleteInsert Delete first and then insert.
func (s *S000001Company) DeleteInsert(del func(del *hey.Del, where hey.Filter), create interface{}, add func(add *hey.Add)) (deleteResult int64, insertResult int64, err error) {
	if deleteResult, err = s.Delete(del); err != nil {
		return
	}
	insertResult, err = s.Insert(create, add)
	return
}

// Border SQL identifier boundary characters.
func (s *S000001Company) Border() string {
	return s.border
}

func (s *S000001Company) initial() *S000001Company {
	s.ID = "id"                 // id
	s.PID = "pid"               // pid
	s.NAME = "name"             // name
	s.COUNTRY = "country"       // country
	s.CITY = "city"             // city
	s.ADDRESS = "address"       // address
	s.LOGO = "logo"             // logo
	s.STATE = "state"           // state
	s.REMARK = "remark"         // remark
	s.CREATED_AT = "created_at" // created_at
	s.UPDATED_AT = "updated_at" // updated_at
	s.DELETED_AT = "deleted_at" // deleted_at

	s.columnMap = map[string]*struct{}{
		s.ID:         {}, // id
		s.PID:        {}, // pid
		s.NAME:       {}, // name
		s.COUNTRY:    {}, // country
		s.CITY:       {}, // city
		s.ADDRESS:    {}, // address
		s.LOGO:       {}, // logo
		s.STATE:      {}, // state
		s.REMARK:     {}, // remark
		s.CREATED_AT: {}, // created_at
		s.UPDATED_AT: {}, // updated_at
		s.DELETED_AT: {}, // deleted_at
	}

	s.columnSlice = []string{
		s.ID,         // id
		s.PID,        // pid
		s.NAME,       // name
		s.COUNTRY,    // country
		s.CITY,       // city
		s.ADDRESS,    // address
		s.LOGO,       // logo
		s.STATE,      // state
		s.REMARK,     // remark
		s.CREATED_AT, // created_at
		s.UPDATED_AT, // updated_at
		s.DELETED_AT, // deleted_at
	}

	s.columnIndex = map[string]int{
		s.ID:         0,  // id
		s.PID:        1,  // pid
		s.NAME:       2,  // name
		s.COUNTRY:    3,  // country
		s.CITY:       4,  // city
		s.ADDRESS:    5,  // address
		s.LOGO:       6,  // logo
		s.STATE:      7,  // state
		s.REMARK:     8,  // remark
		s.CREATED_AT: 9,  // created_at
		s.UPDATED_AT: 10, // updated_at
		s.DELETED_AT: 11, // deleted_at
	}

	replace := s.way.GetCfg().Manual.Replace
	if replace != nil {

		table := s.table
		newest := table
		if strings.Contains(newest, hey.SqlPoint) {
			newest = strings.ReplaceAll(newest, hey.SqlPoint, fmt.Sprintf("%s%s%s", s.border, hey.SqlPoint, s.border))
		}
		newest = fmt.Sprintf("%s%s%s", s.border, newest, s.border)
		replace.Set(table, newest)
		replace.Set(s.ID, `"id"`)                 // id
		replace.Set(s.PID, `"pid"`)               // pid
		replace.Set(s.NAME, `"name"`)             // name
		replace.Set(s.COUNTRY, `"country"`)       // country
		replace.Set(s.CITY, `"city"`)             // city
		replace.Set(s.ADDRESS, `"address"`)       // address
		replace.Set(s.LOGO, `"logo"`)             // logo
		replace.Set(s.STATE, `"state"`)           // state
		replace.Set(s.REMARK, `"remark"`)         // remark
		replace.Set(s.CREATED_AT, `"created_at"`) // created_at
		replace.Set(s.UPDATED_AT, `"updated_at"`) // updated_at
		replace.Set(s.DELETED_AT, `"deleted_at"`) // deleted_at

	}
	return s
}

func newS000001Company(basic abc.BASIC, way *hey.Way) *S000001Company {
	s := &S000001Company{}
	s.table = "public.company"
	s.comment = "company"
	s.border = `"`
	s.basic = &basic
	s.way = way
	s.initial()
	return s
}

type INSERTCompany struct {
	Pid     int    `json:"pid" db:"pid" validate:"omitempty"`                       // pid
	Name    string `json:"name" db:"name" validate:"omitempty,min=0,max=128"`       // name
	Country string `json:"country" db:"country" validate:"omitempty,min=0,max=128"` // country
	City    string `json:"city" db:"city" validate:"omitempty,min=0,max=128"`       // city
	Address string `json:"address" db:"address" validate:"omitempty,min=0,max=128"` // address
	Logo    string `json:"logo" db:"logo" validate:"omitempty,min=0,max=255"`       // logo
	State   int    `json:"state" db:"state" validate:"omitempty"`                   // state
	Remark  string `json:"remark" db:"remark" validate:"omitempty"`                 // remark
}

func (s INSERTCompany) PrimaryKey() interface{} {
	return nil
}

type DELETECompany struct {
	Id *int `json:"id" db:"id" validate:"required,min=1"` // id
}

type UPDATECompany struct {
	DELETECompany
	Pid     *int    `json:"pid" db:"pid" validate:"omitempty"`                       // pid
	Name    *string `json:"name" db:"name" validate:"omitempty,min=0,max=128"`       // name
	Country *string `json:"country" db:"country" validate:"omitempty,min=0,max=128"` // country
	City    *string `json:"city" db:"city" validate:"omitempty,min=0,max=128"`       // city
	Address *string `json:"address" db:"address" validate:"omitempty,min=0,max=128"` // address
	Logo    *string `json:"logo" db:"logo" validate:"omitempty,min=0,max=255"`       // logo
	State   *int    `json:"state" db:"state" validate:"omitempty"`                   // state
	Remark  *string `json:"remark" db:"remark" validate:"omitempty"`                 // remark
}

/* RowsScan, scan data directly, without using reflect. */

func (s *Company) rowsScanInitializePointer() {}

func (s *S000001Company) RowsScanAll(custom func(get *hey.Get, where hey.Filter)) ([]*Company, error) {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	get.Where(func(f hey.Filter) { f.Use(where) }).Select(s.columnSlice...)
	return hey.RowsScanStructAllCmder(get.GetContext(), get.GetWay(), func(rows *sql.Rows, tmp *Company) error {
		tmp.rowsScanInitializePointer()
		return rows.Scan(
			&tmp.Id,
			&tmp.Pid,
			&tmp.Name,
			&tmp.Country,
			&tmp.City,
			&tmp.Address,
			&tmp.Logo,
			&tmp.State,
			&tmp.Remark,
			&tmp.CreatedAt,
			&tmp.UpdatedAt,
			&tmp.DeletedAt,
		)
	}, get)
}

func (s *S000001Company) RowsScanOne(custom func(get *hey.Get, where hey.Filter)) (*Company, error) {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	get.Where(func(f hey.Filter) { f.Use(where) }).Select(s.columnSlice...).Limit(1)
	return hey.RowsScanStructOneCmder(get.GetContext(), get.GetWay(), func(rows *sql.Rows, tmp *Company) error {
		tmp.rowsScanInitializePointer()
		return rows.Scan(
			&tmp.Id,
			&tmp.Pid,
			&tmp.Name,
			&tmp.Country,
			&tmp.City,
			&tmp.Address,
			&tmp.Logo,
			&tmp.State,
			&tmp.Remark,
			&tmp.CreatedAt,
			&tmp.UpdatedAt,
			&tmp.DeletedAt,
		)
	}, get)
}

func (s *S000001Company) RowsScanAllMap(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Company) string) (map[string]*Company, []*Company, error) {
	all, err := s.RowsScanAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[string]*Company, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s *S000001Company) RowsScanAllMapInt(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Company) int) (map[int]*Company, []*Company, error) {
	all, err := s.RowsScanAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int]*Company, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s *S000001Company) RowsScanAllMapInt64(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Company) int64) (map[int64]*Company, []*Company, error) {
	all, err := s.RowsScanAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int64]*Company, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s DELETECompany) PrimaryKey() interface{} {
	if s.Id != nil {
		return *s.Id
	}
	return nil
}

// PrimaryKey Table primary key column name.
func (s *S000001Company) PrimaryKey() string {
	return s.ID
}

// PrimaryKeyUpdate Update based on the primary key as a condition. primaryKey can be any struct or struct pointer that implements the PrimaryKey interface. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyUpdate(primaryKey abc.PrimaryKey, custom func(mod *hey.Mod, where hey.Filter)) (int64, error) {
	if primaryKey == nil {
		return 0, nil
	}
	pk := primaryKey.PrimaryKey()
	if pk == nil {
		return 0, nil
	}
	return s.Update(func(mod *hey.Mod, where hey.Filter) {
		where.Equal(s.PrimaryKey(), pk)
		if custom != nil {
			custom(mod, where)
		}
		mod.Update(primaryKey)
	})
}

// PrimaryKeyHidden Hidden based on the primary key as a condition. primaryKey can be any struct or struct pointer that implements the PrimaryKey interface. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyHidden(primaryKey abc.PrimaryKey, custom func(mod *hey.Mod, where hey.Filter)) (int64, error) {
	if primaryKey == nil {
		return 0, nil
	}
	pk := primaryKey.PrimaryKey()
	if pk == nil {
		return 0, nil
	}
	updates := make(map[string]interface{}, 1<<3)
	way := s.Way()
	now := way.Now()
	for _, tmp := range s.ColumnDeletedAt() {
		updates[tmp] = now.Unix()
	}
	if len(updates) == 0 {
		return 0, nil
	}
	return s.Update(func(mod *hey.Mod, where hey.Filter) {
		where.Equal(s.PrimaryKey(), pk)
		if custom != nil {
			custom(mod, where)
		}
		mod.Update(updates)
	})
}

// PrimaryKeyDelete Delete based on the primary key as a condition. primaryKey can be any struct or struct pointer that implements the PrimaryKey interface. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyDelete(primaryKey abc.PrimaryKey, custom func(del *hey.Del, where hey.Filter)) (int64, error) {
	if primaryKey == nil {
		return 0, nil
	}
	pk := primaryKey.PrimaryKey()
	if pk == nil {
		return 0, nil
	}
	return s.Delete(func(del *hey.Del, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(pk))
		if custom != nil {
			custom(del, where)
		}
	})
}

// PrimaryKeyUpsert Upsert based on the primary key as a condition. primaryKey can be any struct or struct pointer that implements the PrimaryKey interface. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyUpsert(primaryKey abc.PrimaryKey, add func(add *hey.Add), get func(get *hey.Get, where hey.Filter), mod func(mod *hey.Mod, where hey.Filter)) (int64, error) {
	if primaryKey == nil {
		return 0, nil
	}
	pk := primaryKey.PrimaryKey()
	if pk == nil {
		return s.AddOne(primaryKey, add)
	}
	exists, err := s.SelectExists(func(query *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(pk))
		if get != nil {
			get(query, where)
		}
		query.Select(s.columnSlice[0]).Limit(1)
	})
	if err != nil {
		return 0, err
	}
	if !exists {
		return s.AddOne(primaryKey, add)
	}
	return s.PrimaryKeyUpdate(primaryKey, mod)
}

// PrimaryKeyUpdateAll Batch update based on primary key value.
func (s *S000001Company) PrimaryKeyUpdateAll(ctx context.Context, way *hey.Way, update func(mod *hey.Mod, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
	if len(pks) == 0 {
		return 0, nil
	}
	var total int64
	batch := func(tx *hey.Way) error {
		for _, tmp := range pks {
			if num, err := s.PrimaryKeyHidden(tmp, func(mod *hey.Mod, where hey.Filter) {
				if update != nil {
					update(mod, where)
				}
				mod.SetWay(tx)
			}); err != nil {
				return err
			} else {
				total += num
			}
		}
		return nil
	}
	if way != nil && way.IsInTransaction() {
		if err := batch(way); err != nil {
			return total, err
		}
		return total, nil
	}
	err := s.Way().Transaction(ctx, func(tx *hey.Way) error { return batch(tx) })
	return total, err
}

// PrimaryKeyHiddenAll Batch hidden based on primary key value.
func (s *S000001Company) PrimaryKeyHiddenAll(ctx context.Context, way *hey.Way, hidden func(del *hey.Mod, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
	if len(pks) == 0 {
		return 0, nil
	}
	var total int64
	batch := func(tx *hey.Way) error {
		for _, tmp := range pks {
			if num, err := s.PrimaryKeyHidden(tmp, func(mod *hey.Mod, where hey.Filter) {
				if hidden != nil {
					hidden(mod, where)
				}
				mod.SetWay(tx)
			}); err != nil {
				return err
			} else {
				total += num
			}
		}
		return nil
	}
	if way != nil && way.IsInTransaction() {
		if err := batch(way); err != nil {
			return total, err
		}
		return total, nil
	}
	err := s.Way().Transaction(ctx, func(tx *hey.Way) error { return batch(tx) })
	return total, err
}

// PrimaryKeyDeleteAll Batch deletes it based on primary key value.
func (s *S000001Company) PrimaryKeyDeleteAll(ctx context.Context, way *hey.Way, remove func(del *hey.Del, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
	if len(pks) == 0 {
		return 0, nil
	}
	var total int64
	batch := func(tx *hey.Way) error {
		for _, tmp := range pks {
			if num, err := s.PrimaryKeyDelete(tmp, func(del *hey.Del, where hey.Filter) {
				if remove != nil {
					remove(del, where)
				}
				del.SetWay(tx)
			}); err != nil {
				return err
			} else {
				total += num
			}
		}
		return nil
	}
	if way != nil && way.IsInTransaction() {
		if err := batch(way); err != nil {
			return total, err
		}
		return total, nil
	}
	err := s.Way().Transaction(ctx, func(tx *hey.Way) error { return batch(tx) })
	return total, err
}

// PrimaryKeyUpsertAll Batch upsert based on primary key value.
func (s *S000001Company) PrimaryKeyUpsertAll(ctx context.Context, way *hey.Way, add func(add *hey.Add), get func(get *hey.Get, where hey.Filter), mod func(mod *hey.Mod, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
	if len(pks) == 0 {
		return 0, nil
	}
	var total int64
	var err error
	var num int64
	batch := func(tx *hey.Way) error {
		for _, tmp := range pks {
			if tmp == nil {
				continue
			}
			num, err = s.PrimaryKeyUpsert(
				tmp,
				func(obj *hey.Add) {
					if add != nil {
						add(obj)
					}
					obj.SetWay(tx)
				},
				func(obj *hey.Get, where hey.Filter) {
					if get != nil {
						get(obj, where)
					}
					obj.SetWay(tx)
				},
				func(obj *hey.Mod, where hey.Filter) {
					if mod != nil {
						mod(obj, where)
					}
					obj.SetWay(tx)
				},
			)
			if err != nil {
				return err
			} else {
				total += num
			}
		}
		return nil
	}
	if way != nil && way.IsInTransaction() {
		if err = batch(way); err != nil {
			return total, err
		}
		return total, nil
	}
	err = s.Way().Transaction(ctx, func(tx *hey.Way) error { return batch(tx) })
	return total, err
}

// PrimaryKeyEqual Build Filter PrimaryKey = value
func (s *S000001Company) PrimaryKeyEqual(value interface{}) hey.Filter {
	return s.way.F().Equal(s.PrimaryKey(), value)
}

// PrimaryKeyIn Build Filter PrimaryKey IN ( values... )
func (s *S000001Company) PrimaryKeyIn(values ...interface{}) hey.Filter {
	return s.way.F().In(s.PrimaryKey(), values...)
}

// PrimaryKeyUpdateMap Update a row of data using map[string]interface{} by primary key value. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyUpdateMap(primaryKeyValue interface{}, modify map[string]interface{}, update func(mod *hey.Mod, where hey.Filter)) (int64, error) {
	if primaryKeyValue == nil || len(modify) == 0 {
		return 0, nil
	}
	return s.Update(func(mod *hey.Mod, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		if update != nil {
			update(mod, where)
		}
		mod.Update(modify)
	})
}

// PrimaryKeyUpsertMap Upsert a row of data using map[string]interface{} by primary key value. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyUpsertMap(primaryKeyValue interface{}, upsert map[string]interface{}, way *hey.Way) (int64, error) {
	if len(upsert) == 0 {
		return 0, nil
	}
	if primaryKeyValue == nil {
		return s.Insert(upsert, func(add *hey.Add) { add.SetWay(way) })
	}
	exists, err := s.PrimaryKeySelectExists(primaryKeyValue, func(get *hey.Get, where hey.Filter) { get.SetWay(way) })
	if err != nil {
		return 0, err
	}
	if !exists {
		return s.Insert(upsert, func(add *hey.Add) { add.SetWay(way) })
	}
	return s.Update(func(mod *hey.Mod, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		mod.SetWay(way).Update(upsert)
	})
}

// PrimaryKeySelectAll Query multiple records based on primary key values. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeySelectAll(primaryKeyValues interface{}, custom func(get *hey.Get, where hey.Filter)) ([]*Company, error) {
	return s.SelectAll(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyIn(primaryKeyValues))
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectOne Query a piece of data based on the primary key value. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeySelectOne(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (*Company, error) {
	return s.SelectOne(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectOneAsc Query a piece of data based on the primary key value. Additional conditions can be added in the filter. ORDER BY PrimaryKey ASC
func (s *S000001Company) PrimaryKeySelectOneAsc(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (*Company, error) {
	return s.SelectOne(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		get.Asc(s.PrimaryKey())
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectOneDesc Query a piece of data based on the primary key value. Additional conditions can be added in the filter. ORDER BY PrimaryKey DESC
func (s *S000001Company) PrimaryKeySelectOneDesc(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (*Company, error) {
	return s.SelectOne(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		get.Desc(s.PrimaryKey())
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectExists Check whether the data exists based on the primary key value. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeySelectExists(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (bool, error) {
	if primaryKeyValue == nil {
		return false, nil
	}
	exists, err := s.PrimaryKeySelectOne(primaryKeyValue, func(get *hey.Get, where hey.Filter) {
		if custom != nil {
			custom(get, where)
		}
		get.Select(s.PrimaryKey())
	})
	if err != nil && !errors.Is(err, hey.RecordDoesNotExists) {
		return false, err
	}
	return exists != nil, nil
}

// PrimaryKeySelectCount The number of statistics based on primary key values. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeySelectCount(primaryKeyValues interface{}, custom func(get *hey.Get, where hey.Filter)) (int64, error) {
	if primaryKeyValues == nil {
		return 0, nil
	}
	return s.SelectCount(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyIn(primaryKeyValues))
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectAllMap Make map[int]*Company and []*Company
func (s *S000001Company) PrimaryKeySelectAllMap(primaryKeys interface{}, custom func(get *hey.Get, where hey.Filter)) (map[int]*Company, []*Company, error) {
	return s.SelectAllMapInt(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyIn(primaryKeys))
		if custom != nil {
			custom(get, where)
		}
	}, func(v *Company) int { return v.Id })
}

// PrimaryKeyUpsertOne Update or Insert one.
func (s *S000001Company) PrimaryKeyUpsertOne(primaryKeyValue interface{}, upsert interface{}, get func(get *hey.Get, where hey.Filter), add func(add *hey.Add), mod func(mod *hey.Mod, where hey.Filter)) (exists bool, affectedRowsOrIdValue int64, err error) {
	exists, err = s.SelectExists(func(query *hey.Get, where hey.Filter) {
		query.Select(s.PrimaryKey())
		where.Equal(s.PrimaryKey(), primaryKeyValue)
		if get != nil {
			get(query, where)
		}
	})
	if err != nil {
		return
	}
	if exists {
		affectedRowsOrIdValue, err = s.Update(func(update *hey.Mod, where hey.Filter) {
			where.Equal(s.PrimaryKey(), primaryKeyValue)
			if mod != nil {
				mod(update, where)
			}
			update.Update(upsert)
		})
		if err != nil {
			return
		}
		return
	}
	affectedRowsOrIdValue, err = s.AddOne(upsert, add)
	return
}

// Backup Constructing a backup statement.
func (s *S000001Company) Backup(limit int64, custom func(get *hey.Get, where hey.Filter), backup func(add *hey.Add, creates interface{}) (affectedRows int64, err error)) error {
	if backup == nil {
		return nil
	}
	var idMin int
	var affectedRows int64
	var err error
	var lists []*Company
	for {
		lists, err = s.RowsScanAll(func(get *hey.Get, where hey.Filter) {
			where.GreaterThan(s.PrimaryKey(), idMin)
			if custom != nil {
				custom(get, where)
			}
			get.Asc(s.PrimaryKey()).Limit(limit)
		})
		if err != nil {
			return err
		}
		length := len(lists)
		if length == 0 {
			return nil
		}
		affectedRows, err = backup(s.way.Add(s.table), lists)
		if err != nil {
			return err
		}
		if affectedRows != int64(length) {
			return fmt.Errorf("expected %d row(s), got %d", length, affectedRows)
		}
		idMin = lists[length-1].Id
	}
}

// NotFoundInsert If it does not exist, it will be created.
func (s *S000001Company) NotFoundInsert(create interface{}, get func(get *hey.Get, where hey.Filter), add func(add *hey.Add)) (exists bool, err error) {
	exists, err = s.SelectExists(get)
	if err != nil {
		return
	}
	if !exists {
		err = hey.MustAffectedRows(s.Insert(create, add))
		return
	}
	return
}

// Truncate Clear all data in the table.
func (s *S000001Company) Truncate(ctx context.Context) (int64, error) {
	table := s.way.Replace(s.table)
	if ctx == nil {
		ctx = context.Background()
	} else {
		if name := ctx.Value("table"); name != nil {
			if tmp, ok := name.(string); ok && name != hey.EmptyString {
				table = s.way.Replace(tmp)
			}
		}
	}
	result, err := s.way.GetDatabase().ExecContext(ctx, hey.ConcatString("TRUNCATE", hey.SqlSpace, "TABLE", hey.SqlSpace, table))
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// ValueStruct struct value
func (s *S000001Company) ValueStruct() interface{} {
	return Company{}
}

// ValueStructPtr struct pointer value
func (s *S000001Company) ValueStructPtr() interface{} {
	return &Company{}
}

// ValueSliceStruct slice struct value
func (s *S000001Company) ValueSliceStruct(capacities ...int) interface{} {
	capacity := 8
	for i := len(capacities) - 1; i >= 0; i++ {
		if capacities[i] >= 0 {
			capacity = capacities[i]
			break
		}
	}
	return make([]Company, 0, capacity)
}

// ValueSliceStructPtr slice struct pointer value
func (s *S000001Company) ValueSliceStructPtr(capacities ...int) interface{} {
	capacity := 8
	for i := len(capacities) - 1; i >= 0; i++ {
		if capacities[i] >= 0 {
			capacity = capacities[i]
			break
		}
	}
	return make([]*Company, 0, capacity)
}

func (s *S000001Company) Alias(aliases ...string) *S000001CompanyAlias {
	alias := s.table
	if tmp := hey.LastNotEmptyString(aliases); tmp != "" {
		alias = tmp
	}
	table := s.way.T().SetAlias(alias)
	column := func(column string) string { return table.Column(column) }
	tmp := &S000001CompanyAlias{
		ID:         column(s.ID),         // id
		PID:        column(s.PID),        // pid
		NAME:       column(s.NAME),       // name
		COUNTRY:    column(s.COUNTRY),    // country
		CITY:       column(s.CITY),       // city
		ADDRESS:    column(s.ADDRESS),    // address
		LOGO:       column(s.LOGO),       // logo
		STATE:      column(s.STATE),      // state
		REMARK:     column(s.REMARK),     // remark
		CREATED_AT: column(s.CREATED_AT), // created_at
		UPDATED_AT: column(s.UPDATED_AT), // updated_at
		DELETED_AT: column(s.DELETED_AT), // deleted_at

		table: s.table,
		alias: alias,
	}
	tmp.S000001Company = s
	tmp.tableColumn = table
	return tmp
}

func (s *S000001Company) AliasA() *S000001CompanyAlias {
	return s.Alias(hey.AliasA)
}

func (s *S000001Company) AliasB() *S000001CompanyAlias {
	return s.Alias(hey.AliasB)
}

func (s *S000001Company) AliasC() *S000001CompanyAlias {
	return s.Alias(hey.AliasC)
}

func (s *S000001Company) AliasD() *S000001CompanyAlias {
	return s.Alias(hey.AliasD)
}

func (s *S000001Company) AliasE() *S000001CompanyAlias {
	return s.Alias(hey.AliasE)
}

func (s *S000001Company) AliasF() *S000001CompanyAlias {
	return s.Alias(hey.AliasF)
}

func (s *S000001Company) AliasG() *S000001CompanyAlias {
	return s.Alias(hey.AliasG)
}

type S000001CompanyAlias struct {
	*S000001Company
	tableColumn *hey.TableColumn

	ID         string // id
	PID        string // pid
	NAME       string // name
	COUNTRY    string // country
	CITY       string // city
	ADDRESS    string // address
	LOGO       string // logo
	STATE      string // state
	REMARK     string // remark
	CREATED_AT string // created_at
	UPDATED_AT string // updated_at
	DELETED_AT string // deleted_at

	table string
	alias string
}

func (s *S000001CompanyAlias) Table() string {
	return s.table
}

func (s *S000001CompanyAlias) Alias() string {
	if s.alias != "" {
		return s.alias
	}
	return s.Table()
}

func (s *S000001CompanyAlias) Model() *S000001Company {
	return s.S000001Company
}

func (s *S000001CompanyAlias) TableColumn() *hey.TableColumn {
	return s.tableColumn
}

func (s *S000001CompanyAlias) Column(except ...string) []string {
	return s.TableColumn().ColumnAll(s.Model().Column(except...)...)
}
