// code template version: v3.0.0 c1e9fa185ac15a16f8d77f036522e8e685879b26 1751872125-20250707150845
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

// Article | article | article comment
type Article struct {
	Id        int    `json:"id" db:"id"`                 // id comment
	AccountId int    `json:"account_id" db:"account_id"` // account_id comment
	Title     string `json:"title" db:"title"`           // title comment
	Content   string `json:"content" db:"content"`       // content comment
	Stars     int64  `json:"stars" db:"stars"`           // stars comment
	CreatedAt int64  `json:"created_at" db:"created_at"` // created_at comment
	UpdatedAt int64  `json:"updated_at" db:"updated_at"` // updated_at comment
}

type S0000001Article struct {
	ID         string // id comment
	ACCOUNT_ID string // account_id comment
	TITLE      string // title comment
	CONTENT    string // content comment
	STARS      string // stars comment
	CREATED_AT string // created_at comment
	UPDATED_AT string // updated_at comment

	table   string
	comment string

	border string

	columnMap   map[string]*struct{}
	columnSlice []string
	columnIndex map[string]int

	basic *abc.BASIC
	way   *hey.Way
}

func (s *S0000001Article) Basic() *abc.BASIC {
	return s.basic
}

func (s *S0000001Article) Table() string {
	return s.table
}

func (s *S0000001Article) Comment() string {
	return s.comment
}

func (s *S0000001Article) Column(except ...string) []string {
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

func (s *S0000001Article) ColumnMap() map[string]*struct{} {
	result := make(map[string]*struct{}, len(s.columnMap))
	for k, v := range s.columnMap {
		result[k] = v
	}
	return result
}

func (s *S0000001Article) ColumnString() string {
	return `"id", "account_id", "title", "content", "stars", "created_at", "updated_at"`
}

func (s *S0000001Article) ColumnExist(column string) bool {
	_, exist := s.columnMap[column]
	return exist
}

func (s *S0000001Article) ColumnPermit(permit ...string) []string {
	result := make([]string, 0, len(permit))
	for _, v := range permit {
		if ok := s.ColumnExist(v); ok {
			result = append(result, v)
		}
	}
	return result
}

func (s *S0000001Article) ColumnValue(columnValue ...interface{}) map[string]interface{} {
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

func (s *S0000001Article) ColumnAutoIncr() []string {
	return []string{s.ID}
}

func (s *S0000001Article) ColumnCreatedAt() []string {
	return []string{s.CREATED_AT}
}

func (s *S0000001Article) ColumnUpdatedAt() []string {
	return []string{s.UPDATED_AT}
}

func (s *S0000001Article) ColumnDeletedAt() []string {
	return nil
}

func (s *S0000001Article) Filter(filters ...func(f hey.Filter)) hey.Filter {
	filter := s.way.F()
	for _, tmp := range filters {
		if tmp != nil {
			tmp(filter)
		}
	}
	return filter
}

func (s *S0000001Article) Way(ways ...*hey.Way) *hey.Way {
	return abc.Way(s.way, ways...)
}

func (s *S0000001Article) Add(ways ...*hey.Way) *hey.Add {
	excepts := s.ColumnAutoIncr()
	return s.Way(ways...).Add(s.Table()).
		ExceptPermit(
			func(except hey.UpsertColumns, permit hey.UpsertColumns) {
				except.Add(excepts...)
				permit.Add(s.Column(excepts...)...)
			},
		)
}

func (s *S0000001Article) Del(ways ...*hey.Way) *hey.Del {
	return s.Way(ways...).Del(s.Table())
}

func (s *S0000001Article) Mod(ways ...*hey.Way) *hey.Mod {
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

func (s *S0000001Article) Get(ways ...*hey.Way) *hey.Get {
	return s.Way(ways...).Get(s.Table()).Select(s.columnSlice...)
}

func (s *S0000001Article) Available() hey.Filter {
	return s.Filter(func(f hey.Filter) {
		for _, v := range s.ColumnDeletedAt() {
			f.Equal(v, 0)
		}
	})
}

func (s *S0000001Article) Debug(cmder hey.Cmder) {
	s.way.Debug(cmder)
}

// AddOne Insert a record and return the auto-increment id.
func (s *S0000001Article) AddOne(create interface{}, custom func(add *hey.Add)) (int64, error) {
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
func (s *S0000001Article) Insert(create interface{}, custom func(add *hey.Add)) (int64, error) {
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
func (s *S0000001Article) Delete(custom func(del *hey.Del, where hey.Filter)) (int64, error) {
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
func (s *S0000001Article) Update(update func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001Article) InsertSelect(columns []string, get *hey.Get, way *hey.Way) (int64, error) {
	if len(columns) == 0 || get == nil {
		return 0, nil
	}
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()
	return s.Add().SetWay(way).Context(ctx).CmderValues(get, columns).Add()
}

// SelectCount SQL SELECT COUNT.
func (s *S0000001Article) SelectCount(custom func(get *hey.Get, where hey.Filter)) (int64, error) {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	return get.Select(s.columnSlice[0]).Where(func(f hey.Filter) { f.Use(where) }).Count()
}

// SelectQuery SQL SELECT.
func (s *S0000001Article) SelectQuery(custom func(get *hey.Get, where hey.Filter), query func(rows *sql.Rows) error) error {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	return get.Where(func(f hey.Filter) { f.Use(where) }).Query(query)
}

// EmptySlice Initialize an empty slice.
func (s *S0000001Article) EmptySlice() []*Article {
	return make([]*Article, 0)
}

// SelectGet SQL SELECT.
func (s *S0000001Article) SelectGet(custom func(get *hey.Get, where hey.Filter), receive interface{}) error {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	return get.Where(func(f hey.Filter) { f.Use(where) }).Get(receive)
}

// SelectAll SQL SELECT ALL.
func (s *S0000001Article) SelectAll(custom func(get *hey.Get, where hey.Filter)) ([]*Article, error) {
	lists := s.EmptySlice()
	if err := s.SelectGet(custom, &lists); err != nil {
		return nil, err
	}
	return lists, nil
}

// SelectOne SQL SELECT ONE.
func (s *S0000001Article) SelectOne(custom func(get *hey.Get, where hey.Filter)) (*Article, error) {
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
		return nil, hey.ErrNoRows
	}
	return all[0], nil
}

// SelectExists SQL SELECT EXISTS.
func (s *S0000001Article) SelectExists(custom func(get *hey.Get, where hey.Filter)) (bool, error) {
	exists, err := s.SelectOne(func(get *hey.Get, where hey.Filter) {
		if custom != nil {
			custom(get, where)
		}
		get.Select(s.columnSlice[0])
	})
	if err != nil && !errors.Is(err, hey.ErrNoRows) {
		return false, err
	}
	return exists != nil, nil
}

// SelectCountAll SQL SELECT COUNT + ALL.
func (s *S0000001Article) SelectCountAll(custom func(get *hey.Get, where hey.Filter)) (int64, []*Article, error) {
	count, err := s.SelectCount(custom)
	if err != nil {
		return 0, nil, err
	}
	if count == 0 {
		return 0, make([]*Article, 0), nil
	}
	all, err := s.SelectAll(custom)
	if err != nil {
		return 0, nil, err
	}
	return count, all, nil
}

// SelectCountGet SQL SELECT COUNT + GET.
func (s *S0000001Article) SelectCountGet(custom func(get *hey.Get, where hey.Filter), receive interface{}) (int64, error) {
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

// SelectAllMap Make map[string]*Article
func (s *S0000001Article) SelectAllMap(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Article) string) (map[string]*Article, []*Article, error) {
	all, err := s.SelectAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[string]*Article, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

// SelectAllMapInt Make map[int]*Article
func (s *S0000001Article) SelectAllMapInt(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Article) int) (map[int]*Article, []*Article, error) {
	all, err := s.SelectAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int]*Article, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

// SelectAllMapInt64 Make map[int64]*Article
func (s *S0000001Article) SelectAllMapInt64(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Article) int64) (map[int64]*Article, []*Article, error) {
	all, err := s.SelectAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int64]*Article, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

// DeleteByColumn Delete by column values. Additional conditions can be added in the filters. No transaction support.
func (s *S0000001Article) DeleteByColumn(column string, values interface{}, custom func(del *hey.Del, where hey.Filter)) (int64, error) {
	return s.Delete(func(del *hey.Del, where hey.Filter) {
		where.In(column, values)
		if custom != nil {
			custom(del, where)
		}
	})
}

// UpdateByColumn Update by column values. Additional conditions can be added in the filters. No transaction support.
func (s *S0000001Article) UpdateByColumn(column string, values interface{}, update interface{}, custom func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001Article) SelectAllByColumn(column string, values interface{}, customs ...func(get *hey.Get, where hey.Filter)) ([]*Article, error) {
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
func (s *S0000001Article) SelectOneByColumn(column string, values interface{}, customs ...func(get *hey.Get, where hey.Filter)) (*Article, error) {
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
func (s *S0000001Article) SelectExistsByColumn(column string, values interface{}, customs ...func(get *hey.Get, where hey.Filter)) (bool, error) {
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
func (s *S0000001Article) SelectGetByColumn(column string, values interface{}, receive interface{}, customs ...func(get *hey.Get, where hey.Filter)) error {
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
func (s *S0000001Article) DeleteInsert(del func(del *hey.Del, where hey.Filter), create interface{}, add func(add *hey.Add)) (deleteResult int64, insertResult int64, err error) {
	if deleteResult, err = s.Delete(del); err != nil {
		return
	}
	insertResult, err = s.Insert(create, add)
	return
}

// Border SQL identifier boundary characters.
func (s *S0000001Article) Border() string {
	return s.border
}

func (s *S0000001Article) initial() *S0000001Article {
	s.ID = "id"                 // id comment
	s.ACCOUNT_ID = "account_id" // account_id comment
	s.TITLE = "title"           // title comment
	s.CONTENT = "content"       // content comment
	s.STARS = "stars"           // stars comment
	s.CREATED_AT = "created_at" // created_at comment
	s.UPDATED_AT = "updated_at" // updated_at comment

	s.columnMap = map[string]*struct{}{
		s.ID:         {}, // id comment
		s.ACCOUNT_ID: {}, // account_id comment
		s.TITLE:      {}, // title comment
		s.CONTENT:    {}, // content comment
		s.STARS:      {}, // stars comment
		s.CREATED_AT: {}, // created_at comment
		s.UPDATED_AT: {}, // updated_at comment
	}

	s.columnSlice = []string{
		s.ID,         // id comment
		s.ACCOUNT_ID, // account_id comment
		s.TITLE,      // title comment
		s.CONTENT,    // content comment
		s.STARS,      // stars comment
		s.CREATED_AT, // created_at comment
		s.UPDATED_AT, // updated_at comment
	}

	s.columnIndex = map[string]int{
		s.ID:         0, // id comment
		s.ACCOUNT_ID: 1, // account_id comment
		s.TITLE:      2, // title comment
		s.CONTENT:    3, // content comment
		s.STARS:      4, // stars comment
		s.CREATED_AT: 5, // created_at comment
		s.UPDATED_AT: 6, // updated_at comment
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
		replace.Set(s.ID, `"id"`)                 // id comment
		replace.Set(s.ACCOUNT_ID, `"account_id"`) // account_id comment
		replace.Set(s.TITLE, `"title"`)           // title comment
		replace.Set(s.CONTENT, `"content"`)       // content comment
		replace.Set(s.STARS, `"stars"`)           // stars comment
		replace.Set(s.CREATED_AT, `"created_at"`) // created_at comment
		replace.Set(s.UPDATED_AT, `"updated_at"`) // updated_at comment

	}
	return s
}

func newS0000001Article(basic abc.BASIC, way *hey.Way) *S0000001Article {
	s := &S0000001Article{}
	s.table = "article"
	s.comment = "article comment"
	s.border = `"`
	s.basic = &basic
	s.way = way
	s.initial()
	return s
}

type INSERTArticle struct {
	AccountId int    `json:"account_id" db:"account_id" validate:"omitempty"`     // account_id comment
	Title     string `json:"title" db:"title" validate:"omitempty,min=0,max=255"` // title comment
	Content   string `json:"content" db:"content" validate:"omitempty"`           // content comment
	Stars     int64  `json:"stars" db:"stars" validate:"omitempty"`               // stars comment
}

func (s INSERTArticle) PrimaryKey() interface{} {
	return nil
}

type DELETEArticle struct {
	Id *int `json:"id" db:"id" validate:"required,min=1"` // id comment
}

type UPDATEArticle struct {
	DELETEArticle
	AccountId *int    `json:"account_id" db:"account_id" validate:"omitempty"`     // account_id comment
	Title     *string `json:"title" db:"title" validate:"omitempty,min=0,max=255"` // title comment
	Content   *string `json:"content" db:"content" validate:"omitempty"`           // content comment
	Stars     *int64  `json:"stars" db:"stars" validate:"omitempty"`               // stars comment
}

/* RowsScan, scan data directly, without using reflect. */

func (s *Article) rowsScanInitializePointer() {}

func (s *S0000001Article) RowsScanAll(custom func(get *hey.Get, where hey.Filter)) ([]*Article, error) {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	get.Where(func(f hey.Filter) { f.Use(where) }).Select(s.columnSlice...)
	return hey.RowsScanStructAllCmder(get.GetContext(), get.GetWay(), func(rows *sql.Rows, tmp *Article) error {
		tmp.rowsScanInitializePointer()
		return rows.Scan(
			&tmp.Id,
			&tmp.AccountId,
			&tmp.Title,
			&tmp.Content,
			&tmp.Stars,
			&tmp.CreatedAt,
			&tmp.UpdatedAt,
		)
	}, get)
}

func (s *S0000001Article) RowsScanOne(custom func(get *hey.Get, where hey.Filter)) (*Article, error) {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	get.Where(func(f hey.Filter) { f.Use(where) }).Select(s.columnSlice...).Limit(1)
	return hey.RowsScanStructOneCmder(get.GetContext(), get.GetWay(), func(rows *sql.Rows, tmp *Article) error {
		tmp.rowsScanInitializePointer()
		return rows.Scan(
			&tmp.Id,
			&tmp.AccountId,
			&tmp.Title,
			&tmp.Content,
			&tmp.Stars,
			&tmp.CreatedAt,
			&tmp.UpdatedAt,
		)
	}, get)
}

func (s *S0000001Article) RowsScanAllMap(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Article) string) (map[string]*Article, []*Article, error) {
	all, err := s.RowsScanAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[string]*Article, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s *S0000001Article) RowsScanAllMapInt(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Article) int) (map[int]*Article, []*Article, error) {
	all, err := s.RowsScanAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int]*Article, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s *S0000001Article) RowsScanAllMapInt64(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *Article) int64) (map[int64]*Article, []*Article, error) {
	all, err := s.RowsScanAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int64]*Article, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s DELETEArticle) PrimaryKey() interface{} {
	if s.Id != nil {
		return *s.Id
	}
	return nil
}

// PrimaryKey Table primary key column name.
func (s *S0000001Article) PrimaryKey() string {
	return s.ID
}

// PrimaryKeyUpdate Update based on the primary key as a condition. primaryKey can be any struct or struct pointer that implements the PrimaryKey interface. Additional conditions can be added in the filter.
func (s *S0000001Article) PrimaryKeyUpdate(primaryKey abc.PrimaryKey, custom func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001Article) PrimaryKeyHidden(primaryKey abc.PrimaryKey, custom func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001Article) PrimaryKeyDelete(primaryKey abc.PrimaryKey, custom func(del *hey.Del, where hey.Filter)) (int64, error) {
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
func (s *S0000001Article) PrimaryKeyUpsert(primaryKey abc.PrimaryKey, add func(add *hey.Add), get func(get *hey.Get, where hey.Filter), mod func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001Article) PrimaryKeyUpdateAll(ctx context.Context, way *hey.Way, update func(mod *hey.Mod, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
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
func (s *S0000001Article) PrimaryKeyHiddenAll(ctx context.Context, way *hey.Way, hidden func(del *hey.Mod, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
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
func (s *S0000001Article) PrimaryKeyDeleteAll(ctx context.Context, way *hey.Way, remove func(del *hey.Del, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
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
func (s *S0000001Article) PrimaryKeyUpsertAll(ctx context.Context, way *hey.Way, add func(add *hey.Add), get func(get *hey.Get, where hey.Filter), mod func(mod *hey.Mod, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
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
func (s *S0000001Article) PrimaryKeyEqual(value interface{}) hey.Filter {
	return s.way.F().Equal(s.PrimaryKey(), value)
}

// PrimaryKeyIn Build Filter PrimaryKey IN ( values... )
func (s *S0000001Article) PrimaryKeyIn(values ...interface{}) hey.Filter {
	return s.way.F().In(s.PrimaryKey(), values...)
}

// PrimaryKeyUpdateMap Update a row of data using map[string]interface{} by primary key value. Additional conditions can be added in the filter.
func (s *S0000001Article) PrimaryKeyUpdateMap(primaryKeyValue interface{}, modify map[string]interface{}, update func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001Article) PrimaryKeyUpsertMap(primaryKeyValue interface{}, upsert map[string]interface{}, way *hey.Way) (int64, error) {
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
func (s *S0000001Article) PrimaryKeySelectAll(primaryKeyValues interface{}, custom func(get *hey.Get, where hey.Filter)) ([]*Article, error) {
	return s.SelectAll(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyIn(primaryKeyValues))
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectOne Query a piece of data based on the primary key value. Additional conditions can be added in the filter.
func (s *S0000001Article) PrimaryKeySelectOne(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (*Article, error) {
	return s.SelectOne(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectOneAsc Query a piece of data based on the primary key value. Additional conditions can be added in the filter. ORDER BY PrimaryKey ASC
func (s *S0000001Article) PrimaryKeySelectOneAsc(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (*Article, error) {
	return s.SelectOne(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		get.Asc(s.PrimaryKey())
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectOneDesc Query a piece of data based on the primary key value. Additional conditions can be added in the filter. ORDER BY PrimaryKey DESC
func (s *S0000001Article) PrimaryKeySelectOneDesc(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (*Article, error) {
	return s.SelectOne(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		get.Desc(s.PrimaryKey())
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectExists Check whether the data exists based on the primary key value. Additional conditions can be added in the filter.
func (s *S0000001Article) PrimaryKeySelectExists(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (bool, error) {
	if primaryKeyValue == nil {
		return false, nil
	}
	exists, err := s.PrimaryKeySelectOne(primaryKeyValue, func(get *hey.Get, where hey.Filter) {
		if custom != nil {
			custom(get, where)
		}
		get.Select(s.PrimaryKey())
	})
	if err != nil && !errors.Is(err, hey.ErrNoRows) {
		return false, err
	}
	return exists != nil, nil
}

// PrimaryKeySelectCount The number of statistics based on primary key values. Additional conditions can be added in the filter.
func (s *S0000001Article) PrimaryKeySelectCount(primaryKeyValues interface{}, custom func(get *hey.Get, where hey.Filter)) (int64, error) {
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

// PrimaryKeySelectAllMap Make map[int]*Article and []*Article
func (s *S0000001Article) PrimaryKeySelectAllMap(primaryKeys interface{}, custom func(get *hey.Get, where hey.Filter)) (map[int]*Article, []*Article, error) {
	return s.SelectAllMapInt(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyIn(primaryKeys))
		if custom != nil {
			custom(get, where)
		}
	}, func(v *Article) int { return v.Id })
}

// PrimaryKeyUpsertOne Update or Insert one.
func (s *S0000001Article) PrimaryKeyUpsertOne(primaryKeyValue interface{}, upsert interface{}, get func(get *hey.Get, where hey.Filter), add func(add *hey.Add), mod func(mod *hey.Mod, where hey.Filter)) (exists bool, affectedRowsOrIdValue int64, err error) {
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
func (s *S0000001Article) Backup(limit int64, custom func(get *hey.Get, where hey.Filter), backup func(add *hey.Add, creates interface{}) (affectedRows int64, err error)) error {
	if backup == nil {
		return nil
	}
	var idMin int
	var affectedRows int64
	var err error
	var lists []*Article
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
func (s *S0000001Article) NotFoundInsert(create interface{}, get func(get *hey.Get, where hey.Filter), add func(add *hey.Add)) (exists bool, err error) {
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
func (s *S0000001Article) Truncate(ctx context.Context) (int64, error) {
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
func (s *S0000001Article) ValueStruct() interface{} {
	return Article{}
}

// ValueStructPtr struct pointer value
func (s *S0000001Article) ValueStructPtr() interface{} {
	return &Article{}
}

// ValueSliceStruct slice struct value
func (s *S0000001Article) ValueSliceStruct(capacities ...int) interface{} {
	capacity := 8
	for i := len(capacities) - 1; i >= 0; i++ {
		if capacities[i] >= 0 {
			capacity = capacities[i]
			break
		}
	}
	return make([]Article, 0, capacity)
}

// ValueSliceStructPtr slice struct pointer value
func (s *S0000001Article) ValueSliceStructPtr(capacities ...int) interface{} {
	capacity := 8
	for i := len(capacities) - 1; i >= 0; i++ {
		if capacities[i] >= 0 {
			capacity = capacities[i]
			break
		}
	}
	return make([]*Article, 0, capacity)
}

func (s *S0000001Article) Alias(aliases ...string) *S0000001ArticleAlias {
	alias := s.table
	if tmp := hey.LastNotEmptyString(aliases); tmp != "" {
		alias = tmp
	}
	table := s.way.T().SetAlias(alias)
	column := func(column string) string { return table.Column(column) }
	tmp := &S0000001ArticleAlias{
		ID:         column(s.ID),         // id comment
		ACCOUNT_ID: column(s.ACCOUNT_ID), // account_id comment
		TITLE:      column(s.TITLE),      // title comment
		CONTENT:    column(s.CONTENT),    // content comment
		STARS:      column(s.STARS),      // stars comment
		CREATED_AT: column(s.CREATED_AT), // created_at comment
		UPDATED_AT: column(s.UPDATED_AT), // updated_at comment

		table: s.table,
		alias: alias,
	}
	tmp.S0000001Article = s
	tmp.tableColumn = table
	return tmp
}

func (s *S0000001Article) AliasA() *S0000001ArticleAlias {
	return s.Alias(hey.AliasA)
}

func (s *S0000001Article) AliasB() *S0000001ArticleAlias {
	return s.Alias(hey.AliasB)
}

func (s *S0000001Article) AliasC() *S0000001ArticleAlias {
	return s.Alias(hey.AliasC)
}

func (s *S0000001Article) AliasD() *S0000001ArticleAlias {
	return s.Alias(hey.AliasD)
}

func (s *S0000001Article) AliasE() *S0000001ArticleAlias {
	return s.Alias(hey.AliasE)
}

func (s *S0000001Article) AliasF() *S0000001ArticleAlias {
	return s.Alias(hey.AliasF)
}

func (s *S0000001Article) AliasG() *S0000001ArticleAlias {
	return s.Alias(hey.AliasG)
}

type S0000001ArticleAlias struct {
	*S0000001Article
	tableColumn *hey.TableColumn

	ID         string // id comment
	ACCOUNT_ID string // account_id comment
	TITLE      string // title comment
	CONTENT    string // content comment
	STARS      string // stars comment
	CREATED_AT string // created_at comment
	UPDATED_AT string // updated_at comment

	table string
	alias string
}

func (s *S0000001ArticleAlias) Table() string {
	return s.table
}

func (s *S0000001ArticleAlias) Alias() string {
	if s.alias != "" {
		return s.alias
	}
	return s.Table()
}

func (s *S0000001ArticleAlias) Model() *S0000001Article {
	return s.S0000001Article
}

func (s *S0000001ArticleAlias) TableColumn() *hey.TableColumn {
	return s.tableColumn
}

func (s *S0000001ArticleAlias) Column(except ...string) []string {
	return s.TableColumn().ColumnAll(s.Model().Column(except...)...)
}
