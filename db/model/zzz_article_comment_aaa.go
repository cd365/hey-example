// code template version: v3.0.0 e9ec97f8959c580123ea8ffbcfd1e2961fc08160 1750737071-20250624115111
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

// ArticleComment | article_comment | article_comment comment
type ArticleComment struct {
	Id        int    `json:"id" db:"id"`                 // id comment
	AccountId int    `json:"account_id" db:"account_id"` // account_id comment
	ArticleId int    `json:"article_id" db:"article_id"` // article_id comment
	Content   string `json:"content" db:"content"`       // content comment
	CreatedAt int64  `json:"created_at" db:"created_at"` // created_at comment
}

type S0000001ArticleComment struct {
	ID         string // id comment
	ACCOUNT_ID string // account_id comment
	ARTICLE_ID string // article_id comment
	CONTENT    string // content comment
	CREATED_AT string // created_at comment

	table   string
	comment string

	border string

	columnMap   map[string]*struct{}
	columnSlice []string
	columnIndex map[string]int

	basic *abc.BASIC
	way   *hey.Way
}

func (s *S0000001ArticleComment) Basic() *abc.BASIC {
	return s.basic
}

func (s *S0000001ArticleComment) Table() string {
	return s.table
}

func (s *S0000001ArticleComment) Comment() string {
	return s.comment
}

func (s *S0000001ArticleComment) Column(except ...string) []string {
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

func (s *S0000001ArticleComment) ColumnMap() map[string]*struct{} {
	result := make(map[string]*struct{}, len(s.columnMap))
	for k, v := range s.columnMap {
		result[k] = v
	}
	return result
}

func (s *S0000001ArticleComment) ColumnString() string {
	return `"id", "account_id", "article_id", "content", "created_at"`
}

func (s *S0000001ArticleComment) ColumnExist(column string) bool {
	_, exist := s.columnMap[column]
	return exist
}

func (s *S0000001ArticleComment) ColumnPermit(permit ...string) []string {
	result := make([]string, 0, len(permit))
	for _, v := range permit {
		if ok := s.ColumnExist(v); ok {
			result = append(result, v)
		}
	}
	return result
}

func (s *S0000001ArticleComment) ColumnValue(columnValue ...interface{}) map[string]interface{} {
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

func (s *S0000001ArticleComment) ColumnAutoIncr() []string {
	return []string{s.ID}
}

func (s *S0000001ArticleComment) ColumnCreatedAt() []string {
	return []string{s.CREATED_AT}
}

func (s *S0000001ArticleComment) ColumnUpdatedAt() []string {
	return nil
}

func (s *S0000001ArticleComment) ColumnDeletedAt() []string {
	return nil
}

func (s *S0000001ArticleComment) Filter(filters ...func(f hey.Filter)) hey.Filter {
	filter := s.way.F()
	for _, tmp := range filters {
		if tmp != nil {
			tmp(filter)
		}
	}
	return filter
}

func (s *S0000001ArticleComment) Way(ways ...*hey.Way) *hey.Way {
	return abc.Way(s.way, ways...)
}

func (s *S0000001ArticleComment) Add(ways ...*hey.Way) *hey.Add {
	excepts := s.ColumnAutoIncr()
	return s.Way(ways...).Add(s.Table()).
		ExceptPermit(
			func(except hey.UpsertColumns, permit hey.UpsertColumns) {
				except.Add(excepts...)
				permit.Add(s.Column(excepts...)...)
			},
		)
}

func (s *S0000001ArticleComment) Del(ways ...*hey.Way) *hey.Del {
	return s.Way(ways...).Del(s.Table())
}

func (s *S0000001ArticleComment) Mod(ways ...*hey.Way) *hey.Mod {
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

func (s *S0000001ArticleComment) Get(ways ...*hey.Way) *hey.Get {
	return s.Way(ways...).Get(s.Table()).Select(s.columnSlice...)
}

func (s *S0000001ArticleComment) Available() hey.Filter {
	return s.Filter(func(f hey.Filter) {
		for _, v := range s.ColumnDeletedAt() {
			f.Equal(v, 0)
		}
	})
}

func (s *S0000001ArticleComment) Debug(cmder hey.Cmder) {
	s.way.Debug(cmder)
}

// AddOne Insert a record and return the auto-increment id.
func (s *S0000001ArticleComment) AddOne(create interface{}, custom func(add *hey.Add)) (int64, error) {
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
func (s *S0000001ArticleComment) Insert(create interface{}, custom func(add *hey.Add)) (int64, error) {
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
func (s *S0000001ArticleComment) Delete(custom func(del *hey.Del, where hey.Filter)) (int64, error) {
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
func (s *S0000001ArticleComment) Update(update func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001ArticleComment) InsertSelect(columns []string, get *hey.Get, way *hey.Way) (int64, error) {
	if len(columns) == 0 || get == nil {
		return 0, nil
	}
	ctx, cancel := context.WithTimeout(s.basic.Ctx, s.basic.SqlExecuteMaxDuration)
	defer cancel()
	return s.Add().SetWay(way).Context(ctx).CmderValues(get, columns).Add()
}

// SelectCount SQL SELECT COUNT.
func (s *S0000001ArticleComment) SelectCount(custom func(get *hey.Get, where hey.Filter)) (int64, error) {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	return get.Select(s.columnSlice[0]).Where(func(f hey.Filter) { f.Use(where) }).Count()
}

// SelectQuery SQL SELECT.
func (s *S0000001ArticleComment) SelectQuery(custom func(get *hey.Get, where hey.Filter), query func(rows *sql.Rows) error) error {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	return get.Where(func(f hey.Filter) { f.Use(where) }).Query(query)
}

// EmptySlice Initialize an empty slice.
func (s *S0000001ArticleComment) EmptySlice() []*ArticleComment {
	return make([]*ArticleComment, 0)
}

// SelectGet SQL SELECT.
func (s *S0000001ArticleComment) SelectGet(custom func(get *hey.Get, where hey.Filter), receive interface{}) error {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	return get.Where(func(f hey.Filter) { f.Use(where) }).Get(receive)
}

// SelectAll SQL SELECT ALL.
func (s *S0000001ArticleComment) SelectAll(custom func(get *hey.Get, where hey.Filter)) ([]*ArticleComment, error) {
	lists := s.EmptySlice()
	if err := s.SelectGet(custom, &lists); err != nil {
		return nil, err
	}
	return lists, nil
}

// SelectOne SQL SELECT ONE.
func (s *S0000001ArticleComment) SelectOne(custom func(get *hey.Get, where hey.Filter)) (*ArticleComment, error) {
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
func (s *S0000001ArticleComment) SelectExists(custom func(get *hey.Get, where hey.Filter)) (bool, error) {
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
func (s *S0000001ArticleComment) SelectCountAll(custom func(get *hey.Get, where hey.Filter)) (int64, []*ArticleComment, error) {
	count, err := s.SelectCount(custom)
	if err != nil {
		return 0, nil, err
	}
	if count == 0 {
		return 0, make([]*ArticleComment, 0), nil
	}
	all, err := s.SelectAll(custom)
	if err != nil {
		return 0, nil, err
	}
	return count, all, nil
}

// SelectCountGet SQL SELECT COUNT + GET.
func (s *S0000001ArticleComment) SelectCountGet(custom func(get *hey.Get, where hey.Filter), receive interface{}) (int64, error) {
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

// SelectAllMap Make map[string]*ArticleComment
func (s *S0000001ArticleComment) SelectAllMap(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *ArticleComment) string) (map[string]*ArticleComment, []*ArticleComment, error) {
	all, err := s.SelectAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[string]*ArticleComment, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

// SelectAllMapInt Make map[int]*ArticleComment
func (s *S0000001ArticleComment) SelectAllMapInt(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *ArticleComment) int) (map[int]*ArticleComment, []*ArticleComment, error) {
	all, err := s.SelectAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int]*ArticleComment, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

// SelectAllMapInt64 Make map[int64]*ArticleComment
func (s *S0000001ArticleComment) SelectAllMapInt64(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *ArticleComment) int64) (map[int64]*ArticleComment, []*ArticleComment, error) {
	all, err := s.SelectAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int64]*ArticleComment, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

// DeleteByColumn Delete by column values. Additional conditions can be added in the filters. No transaction support.
func (s *S0000001ArticleComment) DeleteByColumn(column string, values interface{}, custom func(del *hey.Del, where hey.Filter)) (int64, error) {
	return s.Delete(func(del *hey.Del, where hey.Filter) {
		where.In(column, values)
		if custom != nil {
			custom(del, where)
		}
	})
}

// UpdateByColumn Update by column values. Additional conditions can be added in the filters. No transaction support.
func (s *S0000001ArticleComment) UpdateByColumn(column string, values interface{}, update interface{}, custom func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001ArticleComment) SelectAllByColumn(column string, values interface{}, customs ...func(get *hey.Get, where hey.Filter)) ([]*ArticleComment, error) {
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
func (s *S0000001ArticleComment) SelectOneByColumn(column string, values interface{}, customs ...func(get *hey.Get, where hey.Filter)) (*ArticleComment, error) {
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
func (s *S0000001ArticleComment) SelectExistsByColumn(column string, values interface{}, customs ...func(get *hey.Get, where hey.Filter)) (bool, error) {
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
func (s *S0000001ArticleComment) SelectGetByColumn(column string, values interface{}, receive interface{}, customs ...func(get *hey.Get, where hey.Filter)) error {
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
func (s *S0000001ArticleComment) DeleteInsert(del func(del *hey.Del, where hey.Filter), create interface{}, add func(add *hey.Add)) (deleteResult int64, insertResult int64, err error) {
	if deleteResult, err = s.Delete(del); err != nil {
		return
	}
	insertResult, err = s.Insert(create, add)
	return
}

// Border SQL identifier boundary characters.
func (s *S0000001ArticleComment) Border() string {
	return s.border
}

func (s *S0000001ArticleComment) initial() *S0000001ArticleComment {
	s.ID = "id"                 // id comment
	s.ACCOUNT_ID = "account_id" // account_id comment
	s.ARTICLE_ID = "article_id" // article_id comment
	s.CONTENT = "content"       // content comment
	s.CREATED_AT = "created_at" // created_at comment

	s.columnMap = map[string]*struct{}{
		s.ID:         {}, // id comment
		s.ACCOUNT_ID: {}, // account_id comment
		s.ARTICLE_ID: {}, // article_id comment
		s.CONTENT:    {}, // content comment
		s.CREATED_AT: {}, // created_at comment
	}

	s.columnSlice = []string{
		s.ID,         // id comment
		s.ACCOUNT_ID, // account_id comment
		s.ARTICLE_ID, // article_id comment
		s.CONTENT,    // content comment
		s.CREATED_AT, // created_at comment
	}

	s.columnIndex = map[string]int{
		s.ID:         0, // id comment
		s.ACCOUNT_ID: 1, // account_id comment
		s.ARTICLE_ID: 2, // article_id comment
		s.CONTENT:    3, // content comment
		s.CREATED_AT: 4, // created_at comment
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
		replace.Set(s.ARTICLE_ID, `"article_id"`) // article_id comment
		replace.Set(s.CONTENT, `"content"`)       // content comment
		replace.Set(s.CREATED_AT, `"created_at"`) // created_at comment

	}
	return s
}

func newS0000001ArticleComment(basic abc.BASIC, way *hey.Way) *S0000001ArticleComment {
	s := &S0000001ArticleComment{}
	s.table = "article_comment"
	s.comment = "article_comment comment"
	s.border = `"`
	s.basic = &basic
	s.way = way
	s.initial()
	return s
}

type INSERTArticleComment struct {
	AccountId int    `json:"account_id" db:"account_id" validate:"omitempty"` // account_id comment
	ArticleId int    `json:"article_id" db:"article_id" validate:"omitempty"` // article_id comment
	Content   string `json:"content" db:"content" validate:"omitempty"`       // content comment
}

func (s INSERTArticleComment) PrimaryKey() interface{} {
	return nil
}

type DELETEArticleComment struct {
	Id *int `json:"id" db:"id" validate:"required,min=1"` // id comment
}

type UPDATEArticleComment struct {
	DELETEArticleComment
	AccountId *int    `json:"account_id" db:"account_id" validate:"omitempty"` // account_id comment
	ArticleId *int    `json:"article_id" db:"article_id" validate:"omitempty"` // article_id comment
	Content   *string `json:"content" db:"content" validate:"omitempty"`       // content comment
}

/* RowsScan, scan data directly, without using reflect. */

func (s *ArticleComment) rowsScanInitializePointer() {}

func (s *S0000001ArticleComment) RowsScanAll(custom func(get *hey.Get, where hey.Filter)) ([]*ArticleComment, error) {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	get.Where(func(f hey.Filter) { f.Use(where) }).Select(s.columnSlice...)
	return hey.RowsScanStructAllCmder(get.GetContext(), get.GetWay(), func(rows *sql.Rows, tmp *ArticleComment) error {
		tmp.rowsScanInitializePointer()
		return rows.Scan(
			&tmp.Id,
			&tmp.AccountId,
			&tmp.ArticleId,
			&tmp.Content,
			&tmp.CreatedAt,
		)
	}, get)
}

func (s *S0000001ArticleComment) RowsScanOne(custom func(get *hey.Get, where hey.Filter)) (*ArticleComment, error) {
	get := s.Get()
	where := s.Filter()
	if custom != nil {
		custom(get, where)
	}
	get.Where(func(f hey.Filter) { f.Use(where) }).Select(s.columnSlice...).Limit(1)
	return hey.RowsScanStructOneCmder(get.GetContext(), get.GetWay(), func(rows *sql.Rows, tmp *ArticleComment) error {
		tmp.rowsScanInitializePointer()
		return rows.Scan(
			&tmp.Id,
			&tmp.AccountId,
			&tmp.ArticleId,
			&tmp.Content,
			&tmp.CreatedAt,
		)
	}, get)
}

func (s *S0000001ArticleComment) RowsScanAllMap(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *ArticleComment) string) (map[string]*ArticleComment, []*ArticleComment, error) {
	all, err := s.RowsScanAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[string]*ArticleComment, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s *S0000001ArticleComment) RowsScanAllMapInt(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *ArticleComment) int) (map[int]*ArticleComment, []*ArticleComment, error) {
	all, err := s.RowsScanAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int]*ArticleComment, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s *S0000001ArticleComment) RowsScanAllMapInt64(custom func(get *hey.Get, where hey.Filter), makeMapKey func(v *ArticleComment) int64) (map[int64]*ArticleComment, []*ArticleComment, error) {
	all, err := s.RowsScanAll(custom)
	if err != nil {
		return nil, nil, err
	}
	allMap := make(map[int64]*ArticleComment, len(all))
	for _, v := range all {
		allMap[makeMapKey(v)] = v
	}
	return allMap, all, nil
}

func (s DELETEArticleComment) PrimaryKey() interface{} {
	if s.Id != nil {
		return *s.Id
	}
	return nil
}

// PrimaryKey Table primary key column name.
func (s *S0000001ArticleComment) PrimaryKey() string {
	return s.ID
}

// PrimaryKeyUpdate Update based on the primary key as a condition. primaryKey can be any struct or struct pointer that implements the PrimaryKey interface. Additional conditions can be added in the filter.
func (s *S0000001ArticleComment) PrimaryKeyUpdate(primaryKey abc.PrimaryKey, custom func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001ArticleComment) PrimaryKeyHidden(primaryKey abc.PrimaryKey, custom func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001ArticleComment) PrimaryKeyDelete(primaryKey abc.PrimaryKey, custom func(del *hey.Del, where hey.Filter)) (int64, error) {
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
func (s *S0000001ArticleComment) PrimaryKeyUpsert(primaryKey abc.PrimaryKey, add func(add *hey.Add), get func(get *hey.Get, where hey.Filter), mod func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001ArticleComment) PrimaryKeyUpdateAll(ctx context.Context, way *hey.Way, update func(mod *hey.Mod, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
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
func (s *S0000001ArticleComment) PrimaryKeyHiddenAll(ctx context.Context, way *hey.Way, hidden func(del *hey.Mod, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
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
func (s *S0000001ArticleComment) PrimaryKeyDeleteAll(ctx context.Context, way *hey.Way, remove func(del *hey.Del, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
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
func (s *S0000001ArticleComment) PrimaryKeyUpsertAll(ctx context.Context, way *hey.Way, add func(add *hey.Add), get func(get *hey.Get, where hey.Filter), mod func(mod *hey.Mod, where hey.Filter), pks []abc.PrimaryKey) (int64, error) {
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
func (s *S0000001ArticleComment) PrimaryKeyEqual(value interface{}) hey.Filter {
	return s.way.F().Equal(s.PrimaryKey(), value)
}

// PrimaryKeyIn Build Filter PrimaryKey IN ( values... )
func (s *S0000001ArticleComment) PrimaryKeyIn(values ...interface{}) hey.Filter {
	return s.way.F().In(s.PrimaryKey(), values...)
}

// PrimaryKeyUpdateMap Update a row of data using map[string]interface{} by primary key value. Additional conditions can be added in the filter.
func (s *S0000001ArticleComment) PrimaryKeyUpdateMap(primaryKeyValue interface{}, modify map[string]interface{}, update func(mod *hey.Mod, where hey.Filter)) (int64, error) {
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
func (s *S0000001ArticleComment) PrimaryKeyUpsertMap(primaryKeyValue interface{}, upsert map[string]interface{}, way *hey.Way) (int64, error) {
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
func (s *S0000001ArticleComment) PrimaryKeySelectAll(primaryKeyValues interface{}, custom func(get *hey.Get, where hey.Filter)) ([]*ArticleComment, error) {
	return s.SelectAll(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyIn(primaryKeyValues))
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectOne Query a piece of data based on the primary key value. Additional conditions can be added in the filter.
func (s *S0000001ArticleComment) PrimaryKeySelectOne(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (*ArticleComment, error) {
	return s.SelectOne(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectOneAsc Query a piece of data based on the primary key value. Additional conditions can be added in the filter. ORDER BY PrimaryKey ASC
func (s *S0000001ArticleComment) PrimaryKeySelectOneAsc(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (*ArticleComment, error) {
	return s.SelectOne(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		get.Asc(s.PrimaryKey())
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectOneDesc Query a piece of data based on the primary key value. Additional conditions can be added in the filter. ORDER BY PrimaryKey DESC
func (s *S0000001ArticleComment) PrimaryKeySelectOneDesc(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (*ArticleComment, error) {
	return s.SelectOne(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyEqual(primaryKeyValue))
		get.Desc(s.PrimaryKey())
		if custom != nil {
			custom(get, where)
		}
	})
}

// PrimaryKeySelectExists Check whether the data exists based on the primary key value. Additional conditions can be added in the filter.
func (s *S0000001ArticleComment) PrimaryKeySelectExists(primaryKeyValue interface{}, custom func(get *hey.Get, where hey.Filter)) (bool, error) {
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
func (s *S0000001ArticleComment) PrimaryKeySelectCount(primaryKeyValues interface{}, custom func(get *hey.Get, where hey.Filter)) (int64, error) {
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

// PrimaryKeySelectAllMap Make map[int]*ArticleComment and []*ArticleComment
func (s *S0000001ArticleComment) PrimaryKeySelectAllMap(primaryKeys interface{}, custom func(get *hey.Get, where hey.Filter)) (map[int]*ArticleComment, []*ArticleComment, error) {
	return s.SelectAllMapInt(func(get *hey.Get, where hey.Filter) {
		where.Use(s.PrimaryKeyIn(primaryKeys))
		if custom != nil {
			custom(get, where)
		}
	}, func(v *ArticleComment) int { return v.Id })
}

// PrimaryKeyUpsertOne Update or Insert one.
func (s *S0000001ArticleComment) PrimaryKeyUpsertOne(primaryKeyValue interface{}, upsert interface{}, get func(get *hey.Get, where hey.Filter), add func(add *hey.Add), mod func(mod *hey.Mod, where hey.Filter)) (exists bool, affectedRowsOrIdValue int64, err error) {
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
func (s *S0000001ArticleComment) Backup(limit int64, custom func(get *hey.Get, where hey.Filter), backup func(add *hey.Add, creates interface{}) (affectedRows int64, err error)) error {
	if backup == nil {
		return nil
	}
	var idMin int
	var affectedRows int64
	var err error
	var lists []*ArticleComment
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
func (s *S0000001ArticleComment) NotFoundInsert(create interface{}, get func(get *hey.Get, where hey.Filter), add func(add *hey.Add)) (exists bool, err error) {
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
func (s *S0000001ArticleComment) Truncate(ctx context.Context) (int64, error) {
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
func (s *S0000001ArticleComment) ValueStruct() interface{} {
	return ArticleComment{}
}

// ValueStructPtr struct pointer value
func (s *S0000001ArticleComment) ValueStructPtr() interface{} {
	return &ArticleComment{}
}

// ValueSliceStruct slice struct value
func (s *S0000001ArticleComment) ValueSliceStruct(capacities ...int) interface{} {
	capacity := 8
	for i := len(capacities) - 1; i >= 0; i++ {
		if capacities[i] >= 0 {
			capacity = capacities[i]
			break
		}
	}
	return make([]ArticleComment, 0, capacity)
}

// ValueSliceStructPtr slice struct pointer value
func (s *S0000001ArticleComment) ValueSliceStructPtr(capacities ...int) interface{} {
	capacity := 8
	for i := len(capacities) - 1; i >= 0; i++ {
		if capacities[i] >= 0 {
			capacity = capacities[i]
			break
		}
	}
	return make([]*ArticleComment, 0, capacity)
}

func (s *S0000001ArticleComment) Alias(aliases ...string) *S0000001ArticleCommentAlias {
	alias := s.table
	if tmp := hey.LastNotEmptyString(aliases); tmp != "" {
		alias = tmp
	}
	table := s.way.T().SetAlias(alias)
	column := func(column string) string { return table.Column(column) }
	tmp := &S0000001ArticleCommentAlias{
		ID:         column(s.ID),         // id comment
		ACCOUNT_ID: column(s.ACCOUNT_ID), // account_id comment
		ARTICLE_ID: column(s.ARTICLE_ID), // article_id comment
		CONTENT:    column(s.CONTENT),    // content comment
		CREATED_AT: column(s.CREATED_AT), // created_at comment

		table: s.table,
		alias: alias,
	}
	tmp.S0000001ArticleComment = s
	tmp.tableColumn = table
	return tmp
}

func (s *S0000001ArticleComment) AliasA() *S0000001ArticleCommentAlias {
	return s.Alias(hey.AliasA)
}

func (s *S0000001ArticleComment) AliasB() *S0000001ArticleCommentAlias {
	return s.Alias(hey.AliasB)
}

func (s *S0000001ArticleComment) AliasC() *S0000001ArticleCommentAlias {
	return s.Alias(hey.AliasC)
}

func (s *S0000001ArticleComment) AliasD() *S0000001ArticleCommentAlias {
	return s.Alias(hey.AliasD)
}

func (s *S0000001ArticleComment) AliasE() *S0000001ArticleCommentAlias {
	return s.Alias(hey.AliasE)
}

func (s *S0000001ArticleComment) AliasF() *S0000001ArticleCommentAlias {
	return s.Alias(hey.AliasF)
}

func (s *S0000001ArticleComment) AliasG() *S0000001ArticleCommentAlias {
	return s.Alias(hey.AliasG)
}

type S0000001ArticleCommentAlias struct {
	*S0000001ArticleComment
	tableColumn *hey.TableColumn

	ID         string // id comment
	ACCOUNT_ID string // account_id comment
	ARTICLE_ID string // article_id comment
	CONTENT    string // content comment
	CREATED_AT string // created_at comment

	table string
	alias string
}

func (s *S0000001ArticleCommentAlias) Table() string {
	return s.table
}

func (s *S0000001ArticleCommentAlias) Alias() string {
	if s.alias != "" {
		return s.alias
	}
	return s.Table()
}

func (s *S0000001ArticleCommentAlias) Model() *S0000001ArticleComment {
	return s.S0000001ArticleComment
}

func (s *S0000001ArticleCommentAlias) TableColumn() *hey.TableColumn {
	return s.tableColumn
}

func (s *S0000001ArticleCommentAlias) Column(except ...string) []string {
	return s.TableColumn().ColumnAll(s.Model().Column(except...)...)
}
