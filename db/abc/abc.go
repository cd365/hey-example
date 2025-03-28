// code template version: v3.0.0 c3e763620528071cd91f9f9535dd9700e721d7a5 1743124166-20250328090926
// TEMPLATE CODE DO NOT EDIT IT.

/*
The current package can be used by multiple database models.

SINGLE
db
├── abc
│    ├── abc.go
│    └── echo.go
└── model
    ├── aaa_schema.go
    ├── aaa_table_create.sql
    └── biz
        └── aaa_schema.go

MULTIPLE
db
├── abc
│    ├── abc.go
│    └── echo.go
├── model
│    ├── aaa_schema.go
│    ├── aaa_table_create.sql
│    └── biz
│        └── aaa_schema.go
├── model1
│    ├── aaa_schema.go
│    ├── aaa_table_create.sql
│    └── biz
│        └── aaa_schema.go
└── model2
    ├── aaa_schema.go
    ├── aaa_table_create.sql
    └── biz
        └── aaa_schema.go
*/

package abc

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/cd365/hey/v3"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Way Get the last non-empty element.
func Way(way *hey.Way, ways ...*hey.Way) *hey.Way {
	for i := len(ways) - 1; i >= 0; i-- {
		if ways[i] != nil {
			return ways[i]
		}
	}
	return way
}

// PrimaryKey Used to obtain the primary key column value of the database table.
type PrimaryKey interface {
	PrimaryKey() interface{}
}

type Table interface {
	Basic() *BASIC
	Table() string
	Comment() string
	Column(except ...string) []string
	ColumnMap() map[string]*struct{}
	ColumnString() string
	ColumnExist(column string) bool
	ColumnPermit(permit ...string) []string
	ColumnValue(columnValue ...interface{}) map[string]interface{}
	ColumnAutoIncr() []string
	ColumnCreatedAt() []string
	ColumnUpdatedAt() []string
	ColumnDeletedAt() []string

	Filter(filters ...func(f hey.Filter)) hey.Filter
	Way(ways ...*hey.Way) *hey.Way
	Add(ways ...*hey.Way) *hey.Add
	Del(ways ...*hey.Way) *hey.Del
	Mod(ways ...*hey.Way) *hey.Mod
	Get(ways ...*hey.Way) *hey.Get
	Available() hey.Filter
	Border() string
	Debugger(cmder hey.Cmder)

	AddOne(custom func(add *hey.Add), create interface{}, ways ...*hey.Way) (int64, error)
	Insert(create interface{}, ways ...*hey.Way) (int64, error)
	Delete(where hey.Filter, ways ...*hey.Way) (int64, error)
	Update(update func(f hey.Filter, u *hey.Mod), ways ...*hey.Way) (int64, error)
	InsertSelect(column []string, get *hey.Get, ways ...*hey.Way) (int64, error)
	SelectCount(where hey.Filter, ways ...*hey.Way) (int64, error)
	SelectQuery(where hey.Filter, custom func(get *hey.Get), query func(rows *sql.Rows) error, ways ...*hey.Way) error
	SelectGet(where hey.Filter, custom func(get *hey.Get), receive interface{}, ways ...*hey.Way) error
	SelectExists(where hey.Filter, custom func(get *hey.Get), ways ...*hey.Way) (bool, error)
	SelectCountGet(where hey.Filter, custom func(get *hey.Get), receive interface{}, ways ...*hey.Way) (int64, error)
	DeleteByColumn(column string, values interface{}, filters ...hey.Filter) (int64, error)
	UpdateByColumn(column string, values interface{}, modify interface{}, filters ...hey.Filter) (int64, error)
	SelectExistsByColumn(column string, values interface{}, customs ...func(f hey.Filter, g *hey.Get)) (bool, error)
	SelectGetByColumn(column string, values interface{}, receive interface{}, customs ...func(f hey.Filter, g *hey.Get)) error
	DeleteInsert(where hey.Filter, create interface{}, ways ...*hey.Way) (deleteResult int64, insertResult int64, err error)

	PrimaryKey() string
	PrimaryKeyUpdate(primaryKey PrimaryKey, filter hey.Filter, ways ...*hey.Way) (int64, error)
	PrimaryKeyHidden(primaryKey PrimaryKey, filter hey.Filter, ways ...*hey.Way) (int64, error)
	PrimaryKeyDelete(primaryKey PrimaryKey, filter hey.Filter, ways ...*hey.Way) (int64, error)
	PrimaryKeyUpsert(primaryKey PrimaryKey, filter hey.Filter, ways ...*hey.Way) (int64, error)
	PrimaryKeyUpdateAll(ctx context.Context, way *hey.Way, pks ...PrimaryKey) (int64, error)
	PrimaryKeyHiddenAll(ctx context.Context, way *hey.Way, pks ...PrimaryKey) (int64, error)
	PrimaryKeyDeleteAll(ctx context.Context, way *hey.Way, pks ...PrimaryKey) (int64, error)
	PrimaryKeyUpsertAll(ctx context.Context, way *hey.Way, pks ...PrimaryKey) (int64, error)
	PrimaryKeyEqual(value interface{}) hey.Filter
	PrimaryKeyIn(values ...interface{}) hey.Filter
	PrimaryKeyUpdateMap(primaryKey interface{}, modify map[string]interface{}, filter hey.Filter, ways ...*hey.Way) (int64, error)
	PrimaryKeyUpsertMap(primaryKey interface{}, upsert map[string]interface{}, filter hey.Filter, ways ...*hey.Way) (int64, error)
	PrimaryKeyDeleteFilter(primaryKeys interface{}, filter hey.Filter, ways ...*hey.Way) (int64, error)
	PrimaryKeySelectExists(primaryKey interface{}, filter hey.Filter, ways ...*hey.Way) (bool, error)
	PrimaryKeySelectCount(primaryKeys interface{}, filter hey.Filter, ways ...*hey.Way) (int64, error)
	PrimaryKeyExists(primaryKey interface{}, ways ...*hey.Way) (bool, error)

	ValueStruct() interface{}
	ValueStructPtr() interface{}
	ValueSliceStruct(capacities ...int) interface{}
	ValueSliceStructPtr(capacities ...int) interface{}
}

type DatabaseManager interface {
	Backup(limit int64, backup func(cmder hey.Cmder) (affectedRows int64, err error)) error
}

type BASIC struct {
	// Ctx Execute sql default context.
	Ctx context.Context

	// SqlExecuteMaxDuration Execute sql max duration.
	SqlExecuteMaxDuration time.Duration
}

func (s *BASIC) SetSqlExecuteMaxDuration(duration time.Duration) *BASIC {
	if duration > 0 {
		s.SqlExecuteMaxDuration = duration
	}
	return s
}

type COUNT struct {
	Count int64 `json:"counts" db:"counts"` // total number of rows
}

/* common structures for querying data */

type SelectTotal struct {
	Total bool `json:"total" query:"total" form:"total" validate:"omitempty"` // 是否统计总条数
}

// SelectIndexValueMaxMin MAX or MIN index value.
type SelectIndexValueMaxMin struct {
	IndexValueMax *int64 `json:"index_max" query:"index_max" form:"index_max" validate:"omitempty,min=0"` // maximum (index or serial) value
	IndexValueMin *int64 `json:"index_min" query:"index_min" form:"index_min" validate:"omitempty,min=0"` // minimum (index or serial) value
}

func (s SelectIndexValueMaxMin) HasIndexValueMax() bool {
	return s.IndexValueMax != nil && s.IndexValueMin == nil
}

func (s SelectIndexValueMaxMin) HasIndexValueMin() bool {
	return s.IndexValueMin != nil && s.IndexValueMax == nil
}

func (s SelectIndexValueMaxMin) UseIndexValueMax(column string, filter hey.Filter) bool {
	if !s.HasIndexValueMax() || column == "" || filter == nil {
		return false
	}
	filter.LessThan(column, *s.IndexValueMax)
	return true
}

func (s SelectIndexValueMaxMin) UseIndexValueMin(column string, filter hey.Filter) bool {
	if !s.HasIndexValueMin() || column == "" || filter == nil {
		return false
	}
	filter.GreaterThan(column, *s.IndexValueMin)
	return true
}

type SelectKeyword struct {
	Keyword *string `json:"keyword" query:"keyword" form:"keyword" validate:"omitempty,min=1,max=32"` // search keyword
}

func (s SelectKeyword) GetKeyword() string {
	return fmt.Sprintf("%%%s%%", *s.Keyword)
}

func (s SelectKeyword) UseKeyword(column string, filter hey.Filter) bool {
	if s.Keyword == nil || *s.Keyword == "" || column == "" || filter == nil {
		return false
	}
	filter.Like(column, s.GetKeyword())
	return true
}

var (
	// regexpSelectTimeDuration column BETWEEN start AND end; "1701234567,1701320967" OR "1701234567000,1701320967000" OR "2006-01-02,2006-01-03" OR "2006-01-02 15:04:05,2006-01-03 15:04:05"
	regexpSelectTimeDuration = regexp.MustCompile(`^([1-9][0-9]{9},[1-9][0-9]{9})|([1-9][0-9]{12},[1-9][0-9]{12})|([12][0-9]{3}-[01][0-9]-[0123][0-9],[12][0-9]{3}-[01][0-9]-[0123][0-9])|([12][0-9]{3}-[01][0-9]-[0123][0-9] [012][0-9]:[012345][0-9]:[012345][0-9],[12][0-9]{3}-[01][0-9]-[0123][0-9] [012][0-9]:[012345][0-9]:[012345][0-9])$`)
)

type SelectTimeDuration string // time duration `min-value,max-value` example format: `1701234567,1701320967` OR `1701234567000,1701320967000` OR `2006-01-02,2006-01-03` OR `2006-01-02 15:04:05,2006-01-03 15:04:05`

func (s SelectTimeDuration) location(locations ...*time.Location) (location *time.Location) {
	for i := len(locations) - 1; i >= 0; i-- {
		if locations[i] != nil {
			location = locations[i]
			break
		}
	}
	if location == nil {
		location = time.Local
	}
	return
}

func (s SelectTimeDuration) GetSplit() []string {
	if !regexpSelectTimeDuration.MatchString(string(s)) {
		return nil
	}
	splits := strings.Split(string(s), ",")
	if len(splits) != 2 {
		return nil
	}
	return splits
}

func (s SelectTimeDuration) GetSplitTimestamp() ([]int64, error) {
	splits := s.GetSplit()
	if splits == nil {
		return nil, errors.New("time format error")
	}
	timestamps := make([]int64, 0, len(splits))
	for _, tmp := range splits {
		i64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return nil, err
		}
		timestamps = append(timestamps, i64)
	}
	return timestamps, nil
}

func (s SelectTimeDuration) splitTimeDurationUseTimestamp(column string, filter hey.Filter, splits []string) bool {
	if column == "" || filter == nil {
		return false
	}
	length := len(splits)
	if length != 2 {
		return false
	}
	lists := make([]interface{}, 0, length)
	for _, tmp := range splits {
		i64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return false
		}
		if i64 <= 0 {
			return false
		}
		lists = append(lists, i64)
	}
	if len(splits[0]) != len(splits[1]) {
		return false
	}
	difference := lists[1].(int64) - lists[0].(int64)
	if difference < 0 {
		return false
	}
	if difference == 0 {
		filter.Equal(column, lists[0].(int64))
	} else {
		filter.Between(column, lists[0], lists[1])
	}
	return true
}

func (s SelectTimeDuration) splitTimeDurationUseDateOrDatetime(column string, filter hey.Filter, splits []string, timeLayout string, loc *time.Location) bool {
	if column == "" || filter == nil || len(splits) != 2 {
		return false
	}
	var err error
	var item time.Time
	lists := make([]time.Time, 0, 2)
	for _, tmp := range splits {
		if loc == nil {
			item, err = time.Parse(timeLayout, tmp)
		} else {
			item, err = time.ParseInLocation(timeLayout, tmp, loc)
		}
		if err != nil {
			return false
		}
		lists = append(lists, item)
	}
	if len(lists) != 2 {
		return false
	}
	difference := lists[1].Unix() - lists[0].Unix()
	if difference < 0 {
		return false
	}
	if difference == 0 {
		filter.Equal(column, splits[0])
	} else {
		filter.Between(column, splits[0], splits[1])
	}
	return true
}

func (s SelectTimeDuration) UseTimeDurationTimestamp(column string, filter hey.Filter) bool {
	return s.splitTimeDurationUseTimestamp(column, filter, s.GetSplit())
}

func (s SelectTimeDuration) UseTimeDurationDate(column string, filter hey.Filter, locations ...*time.Location) bool {
	return s.splitTimeDurationUseDateOrDatetime(column, filter, s.GetSplit(), time.DateOnly, s.location(locations...))
}

func (s SelectTimeDuration) UseTimeDurationDatetime(column string, filter hey.Filter, locations ...*time.Location) bool {
	return s.splitTimeDurationUseDateOrDatetime(column, filter, s.GetSplit(), time.DateTime, s.location(locations...))
}

func (s SelectTimeDuration) UseTimeDuration(column string, filter hey.Filter, locations ...*time.Location) bool {
	splits := s.GetSplit()
	if splits == nil {
		return false
	}
	if s.splitTimeDurationUseTimestamp(column, filter, splits) {
		return true
	}
	if s.splitTimeDurationUseDateOrDatetime(column, filter, splits, time.DateOnly, s.location(locations...)) {
		return true
	}
	if s.splitTimeDurationUseDateOrDatetime(column, filter, splits, time.DateTime, s.location(locations...)) {
		return true
	}
	return false
}

type SelectWhereIn string // column IN ( ?, ?, ? ... )  example-integer: 1,2,3... example-string: a,2,c... OR hex("1"),hex("2"),hex("3")...

func (s SelectWhereIn) UseInInt(column string, filter hey.Filter, adjusts ...func(item int64) (int64, error)) bool {
	if column == "" || filter == nil {
		return false
	}
	splits := strings.Split(strings.ReplaceAll(string(s), " ", ""), ",")
	length := len(splits)
	if length == 0 {
		return false
	}
	var adjust func(item int64) (int64, error)
	for i := len(adjusts) - 1; i >= 0; i-- {
		if adjusts[i] != nil {
			adjust = adjusts[i]
			break
		}
	}
	lists := make([]interface{}, 0, length)
	for _, tmp := range splits {
		if i64, err := strconv.ParseInt(tmp, 10, 64); err != nil {
			return false
		} else {
			if adjust != nil {
				if i64, err = adjust(i64); err == nil {
					return false
				}
			}
			lists = append(lists, i64)
		}
	}
	filter.In(column, lists...)
	return true
}

func (s SelectWhereIn) UseInString(column string, filter hey.Filter, adjusts ...func(item string) (string, error)) bool {
	if column == "" || filter == nil {
		return false
	}
	splits := strings.Split(strings.TrimSpace(string(s)), ",")
	length := len(splits)
	if length == 0 {
		return false
	}
	var adjust func(item string) (string, error)
	for i := len(adjusts) - 1; i >= 0; i-- {
		if adjusts[i] != nil {
			adjust = adjusts[i]
			break
		}
	}
	lists := make([]interface{}, 0, length)
	for _, tmp := range splits {
		if adjust != nil {
			if item, err := adjust(tmp); err != nil {
				return false
			} else {
				tmp = item
			}
		}
		lists = append(lists, tmp)
	}
	filter.In(column, lists...)
	return true
}

type SelectOrder struct {
	Order *string `json:"order" query:"order" form:"order" validate:"omitempty,min=1,max=1000"` // order field1:a,field2:d,field3:a... OR StringToHexFunc(field1:a,field2:d,field3:a...)
}

func (s SelectOrder) GetOrder(adjusts ...func(item string) (string, error)) string {
	if s.Order == nil {
		return ""
	}
	orderBy := strings.TrimSpace(*s.Order)
	if orderBy == "" {
		return ""
	}
	if adjusts == nil {
		// try calling hex.DecodeString()
		if bts, err := hex.DecodeString(orderBy); err == nil {
			orderBy = string(bts)
		}
	}
	for i := len(adjusts) - 1; i >= 0; i-- {
		if adjusts[i] != nil {
			adjusted, err := adjusts[i](orderBy)
			if err != nil {
				return ""
			}
			orderBy = adjusted
			break
		}
	}
	return orderBy
}

type SelectLimitOffset struct {
	Limit *int64 `json:"limit" query:"limit" form:"limit" validate:"omitempty,min=1,max=5000"` // page size [1,5000]

	Offset *int64 `json:"offset" query:"offset" form:"offset" validate:"omitempty,min=0,max=100000"` // offset value [0,100000]

	Page *int64 `json:"page" query:"page" form:"page" validate:"omitempty,min=1"` // page value [1,+∞)
}

func (s SelectLimitOffset) GetLimit() int64 {
	if s.Limit == nil {
		return 1
	}
	return *s.Limit
}

func (s SelectLimitOffset) GetPage() int64 {
	if s.Page == nil {
		return 1
	}
	return *s.Page
}

func (s SelectLimitOffset) GetOffset() int64 {
	if s.Page != nil {
		return (*s.Page - 1) * s.GetLimit()
	}
	if s.Offset != nil && *s.Offset >= 0 {
		return *s.Offset
	}
	return 0
}

// Batches Batch process a set of data.
type Batches[V interface{}] struct {
	lists []V
	mutex *sync.Mutex
}

func NewBatches[V interface{}]() *Batches[V] {
	return &Batches[V]{
		lists: make([]V, 0, 1<<5),
	}
}

func (s *Batches[V]) GetLists() []V {
	return s.lists[:]
}

func (s *Batches[V]) SetLists(lists []V) *Batches[V] {
	if lists == nil {
		s.lists = nil
	} else {
		s.lists = lists[:]
	}
	return s
}

func (s *Batches[V]) GetMutex() *sync.Mutex {
	if s.mutex == nil {
		s.mutex = &sync.Mutex{}
	}
	return s.mutex
}

func (s *Batches[V]) SetMutex(mutex *sync.Mutex) *Batches[V] {
	if mutex != nil && s.mutex == nil {
		s.mutex = mutex
	}
	return s
}

func (s *Batches[V]) Iterator(custom func(i int, v V)) *Batches[V] {
	if custom == nil {
		return s
	}
	for index, value := range s.lists {
		custom(index, value)
	}
	return s
}

func (s *Batches[V]) WithLock(custom func(i int, v V), lists ...V) *Batches[V] {
	if custom == nil || lists == nil {
		return s
	}
	mutex := s.GetMutex()
	mutex.Lock()
	defer mutex.Unlock()
	for index, value := range lists {
		custom(index, value)
	}
	return s
}
