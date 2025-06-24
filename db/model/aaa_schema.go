// code template version: v3.0.0 e9ec97f8959c580123ea8ffbcfd1e2961fc08160 1750737071-20250624115111
// TEMPLATE CODE DO NOT EDIT IT.

package model

import (
	"context"
	_ "embed"
	"github.com/cd365/hey-example/db/abc"
	"github.com/cd365/hey/v3"
	"strings"
	"sync"
	"time"
)

//go:embed aaa_table_create.sql
var tableCreateSql []byte

type Database struct {
	schemaMap   map[string]abc.DatabaseTable
	schemaSlice []string

	ACCOUNT         *S0000001Account
	ARTICLE         *S0000001Article
	ARTICLE_COMMENT *S0000001ArticleComment
}

func NewDatabase(ctx context.Context, way *hey.Way, initialize func(db *Database) error) (*Database, error) {
	basic := abc.BASIC{
		Ctx:                   ctx,
		SqlExecuteMaxDuration: time.Minute,
	}
	tmp := &Database{
		ACCOUNT:         newS0000001Account(basic, way),
		ARTICLE:         newS0000001Article(basic, way),
		ARTICLE_COMMENT: newS0000001ArticleComment(basic, way),
	}
	tmp.schemaMap = map[string]abc.DatabaseTable{
		tmp.ACCOUNT.Table():         tmp.ACCOUNT,
		tmp.ARTICLE.Table():         tmp.ARTICLE,
		tmp.ARTICLE_COMMENT.Table(): tmp.ARTICLE_COMMENT,
	}
	tmp.schemaSlice = []string{
		tmp.ACCOUNT.Table(),
		tmp.ARTICLE.Table(),
		tmp.ARTICLE_COMMENT.Table(),
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

// InitializeAliasPrefix It is recommended to call the current function when initializing the object to prevent Replace.Set from panicking due to concurrent reading and writing.
func InitializeAliasPrefix(db *Database, aliasPrefix []string) {
	if db == nil || len(aliasPrefix) == 0 {
		return
	}
	obj := db.schemaMap[db.schemaSlice[0]]
	way := obj.Way()
	cfg := way.GetCfg()
	if cfg == nil || cfg.Manual == nil || cfg.Manual.Replace == nil {
		return
	}
	border := obj.Border()
	if border == hey.EmptyString {
		return
	}
	replace := cfg.Manual.Replace
	replaced := make(map[string]string)
	for key, value := range replace.Map() {
		replaced[key] = value
	}
	for _, aliasName := range aliasPrefix {
		aliasNameWithBorder := hey.ConcatString(border, aliasName, border)
		for key, value := range replaced {
			if strings.Contains(key, hey.SqlPoint) {
				continue
			}
			key1 := hey.ConcatString(aliasName, hey.SqlPoint, key)
			key2 := hey.ConcatString(aliasName, hey.SqlPoint, hey.ConcatString(border, key, border))
			targetValue := hey.ConcatString(aliasNameWithBorder, hey.SqlPoint, value)
			replace.Set(key1, targetValue)
			replace.Set(key2, targetValue)
		}
	}
}
