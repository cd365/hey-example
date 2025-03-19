// code template version: v3.0.0 8232b1ce979cdaf7365eb708a30dddfa0cbaa290 1742361115-20250319131155
// TEMPLATE CODE DO NOT EDIT IT.

package model

import (
	"context"
	"database/sql"
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
	return s.Way(ways...).Get(s.Table()).Columns(func(columns hey.QueryColumns) { columns.AddAll(s.columnSlice...) })
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
func (s *S000001Company) AddOne(custom func(add *hey.Add), create interface{}, ways ...*hey.Way) (int64, error) {
	if create == nil {
		return 0, nil
	}
	add := s.Add(ways...)
	if custom != nil {
		custom(add)
	}
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()

	add.Context(ctx)
	add.Default(func(o *hey.Add) {
		timestamp := o.Way().Now().Unix()
		for _, v := range s.ColumnCreatedAt() {
			o.ColumnValue(v, timestamp)
		}
	}).Create(create)
	return add.Way().AddOne(
		ctx,
		add,
		func(cmder hey.Cmder) hey.Cmder {
			prepare, args := cmder.Cmd()
			return hey.NewCmder(hey.ConcatString(prepare, fmt.Sprintf(" RETURNING %s", s.way.NameReplace(s.PrimaryKey()))), args)
		},
		func(ctx context.Context, stmt *hey.Stmt, args []interface{}) (id int64, err error) {
			err = stmt.QueryRowContext(ctx, func(rows *sql.Row) error { return rows.Scan(&id) }, args...)
			return
		},
	)

}

// Insert SQL INSERT.
func (s *S000001Company) Insert(create interface{}, ways ...*hey.Way) (int64, error) {
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
		return s.AddOne(nil, create)
	}

	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()
	return s.Add(ways...).
		Context(ctx).
		Default(func(o *hey.Add) {
			timestamp := o.Way().Now().Unix()
			for _, v := range s.ColumnCreatedAt() {
				o.ColumnValue(v, timestamp)
			}
		}).
		Create(create).
		Add()
}

// Delete SQL DELETE.
func (s *S000001Company) Delete(where hey.Filter, ways ...*hey.Way) (int64, error) {
	if where.IsEmpty() {
		return 0, nil
	}
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()
	return s.Del(ways...).
		Context(ctx).
		Where(func(f hey.Filter) {
			f.Use(where.Use(s.Available()))
		}).
		Del()
}

// Update SQL UPDATE.
func (s *S000001Company) Update(update func(f hey.Filter, u *hey.Mod), ways ...*hey.Way) (int64, error) {
	filter := s.Filter()
	modify := s.Mod(ways...)
	if update != nil {
		update(filter, modify)
	}
	if filter.IsEmpty() {
		return 0, nil
	}
	modify.Default(func(o *hey.Mod) {
		timestamp := o.Way().Now().Unix()
		for _, v := range s.ColumnUpdatedAt() {
			o.Set(v, timestamp)
		}
	})
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()
	return modify.Context(ctx).Where(func(f hey.Filter) {
		f.Use(filter.Use(s.Available()))
	}).Mod()
}

// InsertSelect SQL INSERT SELECT.
func (s *S000001Company) InsertSelect(columns []string, get *hey.Get, ways ...*hey.Way) (int64, error) {
	if len(columns) == 0 || get == nil {
		return 0, nil
	}
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()
	return s.Add(ways...).Context(ctx).CmderValues(get, columns).Add()
}

// SelectCount SQL SELECT COUNT.
func (s *S000001Company) SelectCount(where hey.Filter, ways ...*hey.Way) (int64, error) {
	return s.Get(ways...).Columns(func(columns hey.QueryColumns) { columns.Add(s.columnSlice[0]) }).Where(func(f hey.Filter) { f.Use(where) }).Count()
}

// SelectQuery SQL SELECT.
func (s *S000001Company) SelectQuery(where hey.Filter, custom func(get *hey.Get), query func(rows *sql.Rows) error, ways ...*hey.Way) error {
	get := s.Get(ways...).Where(func(f hey.Filter) { f.Use(where) })
	if custom != nil {
		custom(get)
	}
	return get.Query(query)
}

// EmptySlice Initialize an empty slice.
func (s *S000001Company) EmptySlice() []*Company {
	return make([]*Company, 0, 1<<5)
}

// SelectGet SQL SELECT.
func (s *S000001Company) SelectGet(where hey.Filter, custom func(get *hey.Get), receive interface{}, ways ...*hey.Way) error {
	get := s.Get(ways...).Where(func(f hey.Filter) { f.Use(where) })
	if custom != nil {
		custom(get)
	}
	return get.Get(receive)
}

// SelectAll SQL SELECT ALL.
func (s *S000001Company) SelectAll(where hey.Filter, custom func(get *hey.Get), ways ...*hey.Way) ([]*Company, error) {
	get := s.Get(ways...).Where(func(f hey.Filter) { f.Use(where) })
	if custom != nil {
		custom(get)
	}
	all := s.EmptySlice()
	err := get.Get(&all)
	if err != nil {
		return nil, err
	}
	return all, nil
}

// SelectOne SQL SELECT ONE.
func (s *S000001Company) SelectOne(where hey.Filter, custom func(get *hey.Get), ways ...*hey.Way) (*Company, error) {
	all, err := s.SelectAll(where, func(get *hey.Get) {
		if custom != nil {
			custom(get)
		}
		get.Limit(1)
	}, ways...)
	if err != nil {
		return nil, err
	}
	if len(all) == 0 {
		return nil, nil
	}
	return all[0], nil
}

// SelectExists SQL SELECT EXISTS.
func (s *S000001Company) SelectExists(where hey.Filter, custom func(get *hey.Get), ways ...*hey.Way) (bool, error) {
	exists, err := s.SelectOne(where, func(get *hey.Get) {
		if custom != nil {
			custom(get)
		}
		get.Columns(func(columns hey.QueryColumns) { columns.Add(s.columnSlice[0]) })
	}, ways...)
	if err != nil {
		return false, err
	}
	return exists != nil, nil
}

// SelectCountAll SQL SELECT COUNT + ALL.
func (s *S000001Company) SelectCountAll(where hey.Filter, custom func(get *hey.Get), ways ...*hey.Way) (int64, []*Company, error) {
	total, err := s.SelectCount(where, ways...)
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return 0, make([]*Company, 0), nil
	}
	all, err := s.SelectAll(where, custom, ways...)
	if err != nil {
		return 0, nil, err
	}
	return total, all, nil
}

// SelectCountGet SQL SELECT COUNT + GET.
func (s *S000001Company) SelectCountGet(where hey.Filter, custom func(get *hey.Get), receive interface{}, ways ...*hey.Way) (int64, error) {
	count, err := s.SelectCount(where, ways...)
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, nil
	}
	if err = s.SelectGet(where, custom, receive, ways...); err != nil {
		return 0, err
	}
	return count, nil
}

// SelectAllMap Make map[string]*Company
func (s *S000001Company) SelectAllMap(where hey.Filter, makeMapKey func(v *Company) string, custom func(get *hey.Get), ways ...*hey.Way) (map[string]*Company, []*Company, error) {
	all, err := s.SelectAll(where, custom, ways...)
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
func (s *S000001Company) SelectAllMapInt(where hey.Filter, makeMapKey func(v *Company) int, custom func(get *hey.Get), ways ...*hey.Way) (map[int]*Company, []*Company, error) {
	all, err := s.SelectAll(where, custom, ways...)
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
func (s *S000001Company) SelectAllMapInt64(where hey.Filter, makeMapKey func(v *Company) int64, custom func(get *hey.Get), ways ...*hey.Way) (map[int64]*Company, []*Company, error) {
	all, err := s.SelectAll(where, custom, ways...)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int64]*Company, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

// DeleteByColumn Delete by column values. Additional conditions can be added in the filters. no transaction support.
func (s *S000001Company) DeleteByColumn(column string, values interface{}, filters ...hey.Filter) (int64, error) {
	return s.Delete(s.Filter().In(column, values).Use(filters...))
}

// UpdateByColumn Update by column values. Additional conditions can be added in the filters. no transaction support.
func (s *S000001Company) UpdateByColumn(column string, values interface{}, update interface{}, filters ...hey.Filter) (int64, error) {
	if update == nil {
		return 0, nil
	}
	return s.Update(func(f hey.Filter, u *hey.Mod) {
		f.In(column, values).Use(filters...)
		u.Update(update)
	})
}

// SelectAllByColumn Select all by column values. no transaction support.
func (s *S000001Company) SelectAllByColumn(column string, values interface{}, customs ...func(f hey.Filter, g *hey.Get)) ([]*Company, error) {
	where := s.Filter().In(column, values)
	return s.SelectAll(where, func(get *hey.Get) {
		for _, custom := range customs {
			if custom != nil {
				custom(where, get)
				break
			}
		}
	})
}

// SelectOneByColumn Select one by column values. no transaction support.
func (s *S000001Company) SelectOneByColumn(column string, values interface{}, customs ...func(f hey.Filter, g *hey.Get)) (*Company, error) {
	where := s.Filter().In(column, values)
	return s.SelectOne(where, func(get *hey.Get) {
		for _, custom := range customs {
			if custom != nil {
				custom(where, get)
				break
			}
		}
	})
}

// SelectExistsByColumn Select exists by column values. no transaction support.
func (s *S000001Company) SelectExistsByColumn(column string, values interface{}, customs ...func(f hey.Filter, g *hey.Get)) (bool, error) {
	where := s.Filter().In(column, values)
	return s.SelectExists(where, func(get *hey.Get) {
		for _, custom := range customs {
			if custom != nil {
				custom(where, get)
				break
			}
		}
	})
}

// SelectGetByColumn Select get by column values. no transaction support.
func (s *S000001Company) SelectGetByColumn(column string, values interface{}, receive interface{}, customs ...func(f hey.Filter, g *hey.Get)) error {
	where := s.Filter().In(column, values)
	return s.SelectGet(where, func(get *hey.Get) {
		for _, custom := range customs {
			if custom != nil {
				custom(where, get)
				break
			}
		}
	}, receive)
}

// DeleteInsert Delete first and then insert.
func (s *S000001Company) DeleteInsert(where hey.Filter, create interface{}, ways ...*hey.Way) (deleteResult int64, insertResult int64, err error) {
	if where != nil && !where.IsEmpty() {
		if deleteResult, err = s.Delete(where, ways...); err != nil {
			return
		}
	}
	insertResult, err = s.Insert(create, ways...)
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

	replacer := s.way.GetCfg().Replacer
	if replacer != nil {

		table := s.table
		newest := table
		if strings.Contains(newest, hey.SqlPoint) {
			newest = strings.ReplaceAll(newest, hey.SqlPoint, fmt.Sprintf("%s%s%s", s.border, hey.SqlPoint, s.border))
		}
		newest = fmt.Sprintf("%s%s%s", s.border, newest, s.border)
		replacer.Add(table, newest)
		replacer.Add(s.ID, `"id"`)                 // id
		replacer.Add(s.PID, `"pid"`)               // pid
		replacer.Add(s.NAME, `"name"`)             // name
		replacer.Add(s.COUNTRY, `"country"`)       // country
		replacer.Add(s.CITY, `"city"`)             // city
		replacer.Add(s.ADDRESS, `"address"`)       // address
		replacer.Add(s.LOGO, `"logo"`)             // logo
		replacer.Add(s.STATE, `"state"`)           // state
		replacer.Add(s.REMARK, `"remark"`)         // remark
		replacer.Add(s.CREATED_AT, `"created_at"`) // created_at
		replacer.Add(s.UPDATED_AT, `"updated_at"`) // updated_at
		replacer.Add(s.DELETED_AT, `"deleted_at"`) // deleted_at

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

func (s *S000001Company) rowsScan() *Company {
	tmp := &Company{}

	return tmp
}

func (s *S000001Company) RowsScanAll(where hey.Filter, custom func(get *hey.Get), ways ...*hey.Way) (lists []*Company, err error) {
	get := s.Get(ways...).Where(func(f hey.Filter) { f.Use(where) })
	if custom != nil {
		custom(get)
	}
	get.Columns(func(columns hey.QueryColumns) { columns.AddAll(s.columnSlice...) })
	lists = make([]*Company, 0, 1<<5)
	if err = get.Query(func(rows *sql.Rows) (err error) {
		for rows.Next() {
			tmp := s.rowsScan()
			if err = rows.Scan(
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
			); err != nil {
				return err
			}
			lists = append(lists, tmp)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return lists, nil
}

func (s *S000001Company) RowsScanOne(where hey.Filter, custom func(get *hey.Get), ways ...*hey.Way) (*Company, error) {
	lists, err := s.RowsScanAll(where, custom, ways...)
	if err != nil {
		return nil, err
	}
	if len(lists) == 0 {
		return nil, nil
	}
	return lists[0], nil
}

func (s *S000001Company) RowsScanAllMap(where hey.Filter, makeMapKey func(v *Company) string, custom func(get *hey.Get), ways ...*hey.Way) (map[string]*Company, []*Company, error) {
	all, err := s.RowsScanAll(where, custom, ways...)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[string]*Company, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s *S000001Company) RowsScanAllMapInt(where hey.Filter, makeMapKey func(v *Company) int, custom func(get *hey.Get), ways ...*hey.Way) (map[int]*Company, []*Company, error) {
	all, err := s.RowsScanAll(where, custom, ways...)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int]*Company, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s *S000001Company) RowsScanAllMapInt64(where hey.Filter, makeMapKey func(v *Company) int64, custom func(get *hey.Get), ways ...*hey.Way) (map[int64]*Company, []*Company, error) {
	all, err := s.RowsScanAll(where, custom, ways...)
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
func (s *S000001Company) PrimaryKeyUpdate(primaryKey abc.PrimaryKey, filter hey.Filter, ways ...*hey.Way) (int64, error) {
	if primaryKey == nil {
		return 0, nil
	}
	pk := primaryKey.PrimaryKey()
	if pk == nil {
		return 0, nil
	}
	return s.Update(func(f hey.Filter, u *hey.Mod) {
		f.Equal(s.PrimaryKey(), pk).Use(filter)
		u.Update(primaryKey)
	}, ways...)
}

// PrimaryKeyHidden Hidden based on the primary key as a condition. primaryKey can be any struct or struct pointer that implements the PrimaryKey interface. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyHidden(primaryKey abc.PrimaryKey, filter hey.Filter, ways ...*hey.Way) (int64, error) {
	if primaryKey == nil {
		return 0, nil
	}
	pk := primaryKey.PrimaryKey()
	if pk == nil {
		return 0, nil
	}
	updates := make(map[string]interface{}, 1<<3)
	way := s.Way(ways...)
	now := way.Now()
	for _, tmp := range s.ColumnDeletedAt() {
		updates[tmp] = now.Unix()
	}
	if len(updates) == 0 {
		return 0, nil
	}
	return s.Update(func(f hey.Filter, u *hey.Mod) {
		f.Equal(s.PrimaryKey(), pk).Use(filter)
		u.Update(updates)
	}, way)
}

// PrimaryKeyDelete Delete based on the primary key as a condition. primaryKey can be any struct or struct pointer that implements the PrimaryKey interface. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyDelete(primaryKey abc.PrimaryKey, filter hey.Filter, ways ...*hey.Way) (int64, error) {
	if primaryKey == nil {
		return 0, nil
	}
	pk := primaryKey.PrimaryKey()
	if pk == nil {
		return 0, nil
	}
	return s.Delete(s.Filter(func(f hey.Filter) {
		f.Equal(s.PrimaryKey(), pk).Use(filter)
	}), ways...)
}

// PrimaryKeyUpsert Upsert based on the primary key as a condition. primaryKey can be any struct or struct pointer that implements the PrimaryKey interface. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyUpsert(primaryKey abc.PrimaryKey, filter hey.Filter, ways ...*hey.Way) (int64, error) {
	if primaryKey == nil {
		return 0, nil
	}
	pk := primaryKey.PrimaryKey()
	if pk == nil {
		return s.AddOne(nil, primaryKey, ways...)
	}
	return s.PrimaryKeyUpdate(primaryKey, filter, ways...)
}

// PrimaryKeyUpdateAll Batch update based on primary key value.
func (s *S000001Company) PrimaryKeyUpdateAll(ctx context.Context, way *hey.Way, pks ...abc.PrimaryKey) (int64, error) {
	var total int64
	err := s.Way(way).Transaction(ctx, func(tx *hey.Way) error {
		for _, tmp := range pks {
			if num, err := s.PrimaryKeyUpdate(tmp, nil, tx); err != nil {
				return err
			} else {
				total += num
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return total, nil
}

// PrimaryKeyHiddenAll Batch hidden based on primary key value.
func (s *S000001Company) PrimaryKeyHiddenAll(ctx context.Context, way *hey.Way, pks ...abc.PrimaryKey) (int64, error) {
	var total int64
	err := s.Way(way).Transaction(ctx, func(tx *hey.Way) error {
		for _, tmp := range pks {
			if num, err := s.PrimaryKeyHidden(tmp, nil, tx); err != nil {
				return err
			} else {
				total += num
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return total, nil
}

// PrimaryKeyDeleteAll Batch delete based on primary key value.
func (s *S000001Company) PrimaryKeyDeleteAll(ctx context.Context, way *hey.Way, pks ...abc.PrimaryKey) (int64, error) {
	var total int64
	err := s.Way(way).Transaction(ctx, func(tx *hey.Way) error {
		for _, tmp := range pks {
			if num, err := s.PrimaryKeyDelete(tmp, nil, tx); err != nil {
				return err
			} else {
				total += num
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return total, nil
}

// PrimaryKeyUpsertAll Batch upsert based on primary key value.
func (s *S000001Company) PrimaryKeyUpsertAll(ctx context.Context, way *hey.Way, pks ...abc.PrimaryKey) (int64, error) {
	var total int64
	var err error
	var num int64
	err = s.Way(way).Transaction(ctx, func(tx *hey.Way) error {
		for _, tmp := range pks {
			if tmp == nil {
				continue
			}
			pk := tmp.PrimaryKey()
			if pk == nil {
				num, err = s.Insert(tmp, tx)
			} else {
				num, err = s.PrimaryKeyUpdate(tmp, nil, tx)
			}
			if err != nil {
				return err
			} else {
				total += num
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return total, nil
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
func (s *S000001Company) PrimaryKeyUpdateMap(primaryKeyValue interface{}, modify map[string]interface{}, filter hey.Filter, ways ...*hey.Way) (int64, error) {
	if primaryKeyValue == nil || len(modify) == 0 {
		return 0, nil
	}
	return s.Update(func(f hey.Filter, u *hey.Mod) {
		f.Use(s.PrimaryKeyEqual(primaryKeyValue), filter)
		u.Update(modify)
	}, ways...)
}

// PrimaryKeyUpsertMap Upsert a row of data using map[string]interface{} by primary key value. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyUpsertMap(primaryKeyValue interface{}, upsert map[string]interface{}, filter hey.Filter, ways ...*hey.Way) (int64, error) {
	if len(upsert) == 0 {
		return 0, nil
	}
	if primaryKeyValue == nil {
		return s.Insert(upsert, ways...)
	}
	exists, err := s.PrimaryKeySelectExists(primaryKeyValue, filter, ways...)
	if err != nil {
		return 0, err
	}
	if !exists {
		return s.Insert(upsert, ways...)
	}
	return s.Update(func(f hey.Filter, u *hey.Mod) {
		f.Use(s.PrimaryKeyEqual(primaryKeyValue), filter)
		u.Update(upsert)
	}, ways...)
}

// PrimaryKeyDeleteFilter Delete one or more records based on the primary key values. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeyDeleteFilter(primaryKeyValues interface{}, filter hey.Filter, ways ...*hey.Way) (int64, error) {
	return s.Delete(s.PrimaryKeyIn(primaryKeyValues).Use(filter), ways...)
}

// PrimaryKeySelectAll Query multiple records based on primary key values. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeySelectAll(primaryKeyValues interface{}, custom func(get *hey.Get), filter hey.Filter, ways ...*hey.Way) ([]*Company, error) {
	return s.SelectAll(s.PrimaryKeyIn(primaryKeyValues).Use(filter, s.Available()), custom, ways...)
}

// PrimaryKeySelectOne Query a piece of data based on the primary key value. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeySelectOne(primaryKeyValue interface{}, custom func(get *hey.Get), filter hey.Filter, ways ...*hey.Way) (*Company, error) {
	return s.SelectOne(s.PrimaryKeyEqual(primaryKeyValue).Use(filter, s.Available()), custom, ways...)
}

// PrimaryKeySelectOneAsc Query a piece of data based on the primary key value. Additional conditions can be added in the filter. ORDER BY PrimaryKey ASC
func (s *S000001Company) PrimaryKeySelectOneAsc(primaryKeyValue interface{}, custom func(get *hey.Get), filter hey.Filter, ways ...*hey.Way) (*Company, error) {
	return s.PrimaryKeySelectOne(primaryKeyValue, func(get *hey.Get) {
		if custom != nil {
			custom(get)
		}
		get.Asc(s.PrimaryKey())
	}, filter, ways...)
}

// PrimaryKeySelectOneDesc Query a piece of data based on the primary key value. Additional conditions can be added in the filter. ORDER BY PrimaryKey DESC
func (s *S000001Company) PrimaryKeySelectOneDesc(primaryKeyValue interface{}, custom func(get *hey.Get), filter hey.Filter, ways ...*hey.Way) (*Company, error) {
	return s.PrimaryKeySelectOne(primaryKeyValue, func(get *hey.Get) {
		if custom != nil {
			custom(get)
		}
		get.Desc(s.PrimaryKey())
	}, filter, ways...)
}

// PrimaryKeySelectExists Check whether the data exists based on the primary key value. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeySelectExists(primaryKeyValue interface{}, filter hey.Filter, ways ...*hey.Way) (bool, error) {
	if primaryKeyValue == nil {
		return false, nil
	}
	exists, err := s.PrimaryKeySelectOne(primaryKeyValue, func(get *hey.Get) { get.Columns(func(columns hey.QueryColumns) { columns.Add(s.PrimaryKey()) }) }, filter, ways...)
	if err != nil {
		return false, err
	}
	return exists != nil, nil
}

// PrimaryKeySelectCount The number of statistics based on primary key values. Additional conditions can be added in the filter.
func (s *S000001Company) PrimaryKeySelectCount(primaryKeyValues interface{}, filter hey.Filter, ways ...*hey.Way) (int64, error) {
	if primaryKeyValues == nil {
		return 0, nil
	}
	return s.SelectCount(s.PrimaryKeyIn(primaryKeyValues).Use(filter, s.Available()), ways...)
}

// PrimaryKeySelectAllMap Make map[int]*Company and []*Company
func (s *S000001Company) PrimaryKeySelectAllMap(primaryKeys interface{}, custom func(get *hey.Get), filter hey.Filter, ways ...*hey.Way) (map[int]*Company, []*Company, error) {
	return s.SelectAllMapInt(s.PrimaryKeyIn(primaryKeys).Use(filter, s.Available()), func(v *Company) int { return v.Id }, custom, ways...)
}

// PrimaryKeyGetAll Query multiple records based on primary key values.
func (s *S000001Company) PrimaryKeyGetAll(primaryKeys interface{}, ways ...*hey.Way) ([]*Company, error) {
	return s.PrimaryKeySelectAll(primaryKeys, nil, nil, ways...)
}

// PrimaryKeyGetOne Query a piece of data based on the primary key value.
func (s *S000001Company) PrimaryKeyGetOne(primaryKey interface{}, ways ...*hey.Way) (*Company, error) {
	return s.PrimaryKeySelectOne(primaryKey, nil, nil, ways...)
}

// PrimaryKeyGetOneAsc Query a piece of data based on the primary key value. ORDER BY PrimaryKey ASC
func (s *S000001Company) PrimaryKeyGetOneAsc(primaryKey interface{}, ways ...*hey.Way) (*Company, error) {
	return s.PrimaryKeySelectOneAsc(primaryKey, nil, nil, ways...)
}

// PrimaryKeyGetOneDesc Query a piece of data based on the primary key value. ORDER BY PrimaryKey DESC
func (s *S000001Company) PrimaryKeyGetOneDesc(primaryKey interface{}, ways ...*hey.Way) (*Company, error) {
	return s.PrimaryKeySelectOneDesc(primaryKey, nil, nil, ways...)
}

// PrimaryKeyGetAllMap Make map[int]*Company and []*Company
func (s *S000001Company) PrimaryKeyGetAllMap(primaryKeys interface{}, ways ...*hey.Way) (map[int]*Company, []*Company, error) {
	return s.PrimaryKeySelectAllMap(primaryKeys, nil, nil, ways...)
}

// PrimaryKeyExists Check whether the data exists based on the primary key value.
func (s *S000001Company) PrimaryKeyExists(primaryKey interface{}, ways ...*hey.Way) (bool, error) {
	return s.PrimaryKeySelectExists(primaryKey, nil, ways...)
}

// UpsertOne Update or Insert one.
func (s *S000001Company) UpsertOne(filter func(f hey.Filter, g *hey.Get), upsert interface{}, ways ...*hey.Way) (exists bool, affectedRowsOrIdValue int64, err error) {
	where := s.Filter()
	first, err := s.SelectOne(where, func(get *hey.Get) {
		get.Columns(func(columns hey.QueryColumns) { columns.Add(s.PrimaryKey()) })
		if filter != nil {
			filter(where, get)
		}
	}, ways...)
	if err != nil {
		return false, 0, err
	}
	exists = first != nil
	if exists {
		affectedRowsOrIdValue, err = s.Update(func(f hey.Filter, u *hey.Mod) {
			f.Equal(s.PrimaryKey(), first.Id)
			u.Update(upsert)
		}, ways...)
	} else {
		affectedRowsOrIdValue, err = s.AddOne(nil, upsert, ways...)
	}
	if err != nil {
		return false, 0, err
	}
	return exists, affectedRowsOrIdValue, nil
}

// NotFoundInsert If it does not exist, it will be created.
func (s *S000001Company) NotFoundInsert(filter func(f hey.Filter, g *hey.Get), create interface{}, ways ...*hey.Way) (exists bool, err error) {
	where := s.Filter()
	first, err := s.SelectOne(where, func(get *hey.Get) {
		get.Columns(func(columns hey.QueryColumns) { columns.Add(s.PrimaryKey()) })
		if filter != nil {
			filter(where, get)
		}
	}, ways...)
	if err != nil {
		return false, err
	}
	exists = first != nil
	if exists {
		return true, nil
	}
	if err = hey.MustAffectedRows(s.Insert(create, ways...)); err != nil {
		return false, err
	}
	return true, nil
}

// Backup Constructing a backup statement.
func (s *S000001Company) Backup(limit int64, backup func(cmder hey.Cmder) (affectedRows int64, err error)) error {
	if backup == nil {
		return nil
	}
	var idMin int
	var affectedRows int64
	var err error
	var lists []*Company
	for {
		lists, err = s.RowsScanAll(
			s.Filter().GreaterThan(s.PrimaryKey(), idMin),
			func(get *hey.Get) {
				get.Asc(s.PrimaryKey()).Limit(limit)
			},
		)
		if err != nil {
			return err
		}
		length := len(lists)
		if length == 0 {
			return nil
		}
		affectedRows, err = backup(s.way.Add(s.table).Create(lists))
		if err != nil {
			return err
		}
		if affectedRows != int64(length) {
			return fmt.Errorf("expected %d row(s), got %d", length, affectedRows)
		}
		idMin = lists[length-1].Id
	}
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
