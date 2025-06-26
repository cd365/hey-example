package examples

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/cd365/hey-example/db/model"
	"github.com/cd365/hey/v3"
	"github.com/cd365/logger/v9"
	_ "github.com/lib/pq"
	"time"
)

/*

1. All codes in the db directory can be dynamically constructed according to the database table structure, and only need to be constructed once each time the table structure is changed;
2. After the table structure is changed, the rebuilt template code can immediately detect the changes in the table name, field name, and field type;
3. All business operation databases should be encapsulated according to business logic, rather than directly relying on the shortcut methods provided by `function-struct`;

*model.Database: All exposed properties are database table objects.

origin-table-name   data-struct             function-struct
account 			model.Account 		 	model.S000001Account
article 			model.Article 		 	model.S000001Article
article_comment 	model.ArticleComment 	model.S000001ArticleComment
data-struct: Structure for writing and reading data.
function-struct: Extended functions of the current table, including column properties and methods related to operating the current table.

*/

type DB struct {
	*model.Database
}

func NewDB() (*DB, error) {
	driverName := "postgres"
	dataSourceName := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", "postgres", "postgres", "example")
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetConnMaxLifetime(time.Minute * 3)
	way := hey.NewWay(db)

	cfg := way.GetCfg()
	cfg.Manual = hey.Postgresql()
	cfg.Manual.Replace = hey.NewReplace()

	colorWrite := logger.NewMyWrite(nil)
	colorWrite.PriorityWrite(colorWrite.ColorWrite)

	// Record SQL log.
	lg := logger.NewLogger(colorWrite)
	way.SetLogger(lg)

	// Customize SQL execution time warning threshold.
	cfg.WarnDuration = time.Millisecond * 200

	debug := hey.NewDebugger()
	debug.SetLogger(logger.NewLogger(colorWrite))
	cfg.Debugger = debug

	database, err := model.NewDatabase(context.Background(), way, func(db *model.Database) error {
		return err
	})

	// Delete specific identifiers in SQL statements.
	cfg.Manual.Replace = nil

	tmp := &DB{}
	tmp.Database = database

	return tmp, nil
}

// InsertOne Create a piece of data and get the auto-increment id value.
func (s *DB) InsertOne() error {
	account := &model.Account{}
	id, err := s.ACCOUNT.AddOne(account, nil)
	if err != nil {
		return err
	}
	account.Id = int(id)
	return nil
}

// Insert Create one or more data.
func (s *DB) Insert() (int64, error) {
	accounts := make([]*model.Account, 0)
	return s.ACCOUNT.Insert(accounts, nil)
}

// InsertWithColumn Create one or more data with the specified columns.
func (s *DB) InsertWithColumn() (int64, error) {
	accounts := make([]*model.Account, 0)
	return s.ACCOUNT.Insert(accounts, func(add *hey.Add) {
		t := s.ACCOUNT
		// Just set the list of allowed columns.
		add.ExceptPermit(func(except hey.UpsertColumns, permit hey.UpsertColumns) {
			permit.Add(t.EMAIL, t.USERNAME, t.PASSWORD, t.STATUS, t.IP, t.CREATED_AT, t.UPDATED_AT)
		})
		// Just set the list of disallowed columns.
		add.ExceptPermit(func(except hey.UpsertColumns, permit hey.UpsertColumns) {
			except.Add(t.ID, t.DELETED_AT)
		})
	})
}

// InsertWithMap Insert data through map.
func (s *DB) InsertWithMap() (int64, error) {
	t := s.ACCOUNT
	now := time.Now()
	return t.Insert(map[string]any{
		t.EMAIL:      "xxx@gmail.com",
		t.USERNAME:   "xxx",
		t.CREATED_AT: now.Unix(),
		t.UPDATED_AT: now.Unix(),
	}, nil)
}

// InsertWithCustom Insert data through custom.
func (s *DB) InsertWithCustom() (int64, error) {
	t := s.ACCOUNT
	now := time.Now()
	return t.Add().
		ColumnValue(t.EMAIL, "xxx@gmail.com").
		ColumnValue(t.USERNAME, "xxx").
		ColumnValue(t.CREATED_AT, now.Unix()).
		ColumnValue(t.UPDATED_AT, now.Unix()).
		Add()
}

// InsertWithQuery Use query results as the value of the insert statement.
func (s *DB) InsertWithQuery() (int64, error) {
	t := s.ACCOUNT
	columns := []string{t.EMAIL, t.USERNAME}
	return t.Add().
		// Table("another_table").
		CmderValues(
			t.Get().Where(func(f hey.Filter) {
				f.LessThanEqual(t.ID, 10)
			}).Select(columns...).Desc(t.ID).Limit(10),
			columns,
		).Add()
}

// Filter Dynamic condition filtering for any combination concatenation.
func (s *DB) Filter() {
	t := s.ACCOUNT

	// Simple conditional filtering and splicing.
	f := t.Filter()
	f.Equal(t.USERNAME, "xxx").
		Between(t.BALANCE, 100, 1000).
		GreaterThan(t.DELETED_AT, 0).
		Between(t.STATUS, 1, 3)
	t.Debug(f)

	// Implementing complex conditional filtering logic.
	f.Clean()
	f.Equal(t.USERNAME, "xxx")
	f.Group(func(g hey.Filter) {
		g.Group(func(g hey.Filter) { g.Between(t.STATUS, 1, 3).GreaterThan(t.BALANCE, 100) })
		g.OrGroup(func(g hey.Filter) { g.Between(t.STATUS, 6, 8).GreaterThan(t.BALANCE, 500) })
	})
	f.IsNotNull(t.UPDATED_AT)
	f.Like(t.EMAIL, "%gmail.com")
	t.Debug(f)

	// column1 IN (SELECT column1 FROM xxx [ WHERE ] [ ORDER BY ] [ LIMIT ])
	f.Clean()
	f.Equal(t.USERNAME, "xxx")
	queryFilter := f.New()
	queryFilter.Between(t.STATUS, 1, 3).GreaterThan(t.BALANCE, 1200)
	f.InQuery(t.USERNAME, t.Get().Select(t.USERNAME).Where(func(f hey.Filter) { f.Use(queryFilter) }).Desc(t.ID).Limit(100))
	t.Debug(f)

	// ( column1, column2, column3 ) IN ( SELECT column1, column2, column3 FROM xxx [ WHERE ] [ ORDER BY ] [ LIMIT ] )
	f.Clean()
	cols := []string{t.EMAIL, t.USERNAME, t.BALANCE}
	f.InColsQuery(cols, t.Get().Select(cols...).Where(func(f hey.Filter) { f.Use(queryFilter) }).Desc(t.ID).Limit(20))
	f.AnyQuantifier(func(tmp hey.Quantifier) {
		tmp.GreaterThan(t.BALANCE, t.Get().Select(t.BALANCE).Where(func(f hey.Filter) {
			f.Equal(t.STATUS, 3)
		}).Limit(10))
	})
	t.Debug(f)

	// Filter allows you to extend methods that are not currently supported.
	// The specific usage is as follows:
	iLike := func(column string, value any) hey.Filter {
		filter := t.Filter()
		way := filter.GetWay()
		return filter.And(fmt.Sprintf("%s ILIKE ?", way.Replace(column)), value)
	}
	f1 := f.New()
	f1.Use(iLike(t.USERNAME, "%xx%"))
	t.Debug(f1)

	// Combine multiple filters at will.
	t.Debug(f.New(f, f1))

	// More formats are as follows:
	fa, fb, fc := f.New(), f.New(), f.New() // ...
	t.Debug(f.New(fa, fb, fc))
}

func (s *DB) accountDeleteIdEqual(id int) func(del *hey.Del, where hey.Filter) {
	t := s.ACCOUNT
	return func(del *hey.Del, where hey.Filter) {
		where.Equal(t.ID, id)
	}
}

// Delete Deleting data.
func (s *DB) Delete() (int64, error) {
	t := s.ACCOUNT
	id := 2
	_, _ = t.Delete(s.accountDeleteIdEqual(id))
	_, _ = t.PrimaryKeyDelete(&model.DELETEAccount{Id: &id}, nil)
	return t.DeleteByColumn(t.EMAIL, "xxx@gmail.com", func(del *hey.Del, where hey.Filter) {
		where.GreaterThan(t.DELETED_AT, 0)
	})
}

// Update Updating data.
func (s *DB) Update() (int64, error) {
	t := s.ACCOUNT
	id := 2
	_, _ = t.Update(func(mod *hey.Mod, where hey.Filter) {
		where.Equal(t.ID, id)
		mod.Incr(t.BALANCE, 100)
		mod.Set(t.IP, "::1")

		// Custom update expressions.
		status := mod.GetWay().Replace(t.STATUS)
		mod.Expr(fmt.Sprintf("%s = %s * 2 + 1", status, status))
	})

	_, _ = t.UpdateByColumn(
		t.EMAIL,
		"xxx@gmail.com",
		map[string]any{
			t.USERNAME: "yyy",
		},
		func(mod *hey.Mod, where hey.Filter) {
			// Optional additional filtering conditions.
			// where.Equal(t.DELETED_AT, 0)
		},
	)

	origin := &model.Account{}
	latest := *origin
	latest.Status = origin.Status + 1
	_, _ = t.Update(func(mod *hey.Mod, where hey.Filter) {
		where.Equal(t.ID, id)
		mod.Compare(origin, &latest) // Auto update: status = origin.Status + 1
	})

	return t.Update(func(mod *hey.Mod, where hey.Filter) {
		where.Equal(t.ID, id)

		mod.ExceptPermit(func(except hey.UpsertColumns, permit hey.UpsertColumns) {
			// Customize the list of columns that are allowed and not allowed to be updated.
		})

		mod.Update(&model.UPDATEAccount{
			// nil value will not be updated
			// ...
		})
	})
}

// Select Query data.
func (s *DB) Select() {
	t1 := s.ACCOUNT
	t2 := s.ARTICLE
	t3 := s.ARTICLE_COMMENT

	{
		now := time.Now()
		// table alias name
		a, b, c := hey.AliasA, hey.AliasB, hey.AliasC
		way := t1.Way()
		// alias name table
		ta, tb, tc := way.TA(), way.TB(), way.TC()
		get := t1.Get().Alias(a).Join(func(j hey.QueryJoin) {
			// current master table
			m := j.GetMaster()

			j1 := j.NewTable(t2.Table(), b)
			j2 := j.NewTable(t3.Table(), c)

			j.LeftJoin(nil, j1, j.OnEqual(t1.ID, t2.ACCOUNT_ID))
			j.LeftJoin(j1, j2, j.OnEqual(t2.ID, t3.ARTICLE_ID))

			// select columns
			j.SelectGroupsColumns(
				j.TableSelect(j2, t3.Column()...),
			)
			j.SelectTableColumnAlias(j1, t2.TITLE, "article_title")
			j.SelectTableColumnAlias(m, t1.EMAIL, "account_email", t1.USERNAME, "account_username")

		}).Where(func(f hey.Filter) {
			f.Equal(ta.Column(t1.STATUS), 1)
			f.Equal(tb.Column(t2.UPDATED_AT), 0)
			f.GreaterThan(tc.Column(t3.CREATED_AT), now.Unix()-86400)
		}).
			// Desc(ta.Column(t1.ID)).
			Order(
				fmt.Sprintf("%s:d", t3.ID),
				// replace columns
				map[string]string{
					t3.ID: tc.Column(t3.ID),
				},
			).Limit(10)
		t1.Debug(get)
	}

	{
		result := &model.Account{}
		err := t1.Get().Select(t1.Column()...).Where(func(f hey.Filter) {
			f.Equal(t1.DELETED_AT, 0)
		}).Desc(t1.ID).Limit(1).Offset(5).Get(result)
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", result)
	}
}

// Transaction For example of Transaction.
func (s *DB) Transaction(ctx context.Context) error {
	t1 := s.ACCOUNT
	t2 := s.ARTICLE
	t3 := s.ARTICLE_COMMENT
	updateErr := func(rows int64, err error) error {
		if err != nil {
			return err
		}
		if rows <= 0 {
			return errors.New("update failed")
		}
		return nil
	}
	return t1.Way().Transaction(ctx, func(tx *hey.Way) error {
		tx.TransactionMessage("try transaction")
		err := updateErr(
			tx.Mod(t1.Table()).Where(func(f hey.Filter) {
				f.Equal(t1.ID, 1)
			}).Set(t1.UPDATED_AT, tx.Now().Unix()).Mod(),
		)
		if err != nil {
			return err
		}
		err = updateErr(
			t2.Mod(tx).Where(func(f hey.Filter) {
				f.Equal(t2.ID, 1)
			}).Set(t2.UPDATED_AT, tx.Now().Unix()).Mod(),
		)
		if err != nil {
			return err
		}
		err = updateErr(
			t2.Update(func(mod *hey.Mod, where hey.Filter) {
				mod.SetWay(tx)
				where.Equal(t2.ID, 2)
				mod.Set(t2.UPDATED_AT, tx.Now().Unix())
			}),
		)
		if err != nil {
			return err
		}
		err = updateErr(
			tx.Del(t3.Table()).Where(func(f hey.Filter) {
				f.Equal(t3.ID, 1)
			}).Del(),
		)
		if err != nil {
			return err
		}
		return nil
	})
}

// Transaction1 Minimize other code in transactions.
func (s *DB) Transaction1() error {
	t1 := s.ACCOUNT
	t2 := s.ARTICLE
	t3 := s.ARTICLE_COMMENT

	way := t1.Way()

	add := way.Add(t1.Table()).
		Create(map[string]interface{}{})

	del := way.Del(t2.Table()).
		Where(func(f hey.Filter) {
			f.Equal(t2.ID, 1)
		})

	mod := way.Mod(t3.Table()).
		Set(t3.CONTENT, "123").
		Where(func(f hey.Filter) {
			f.Equal(t3.ID, 1)
		})

	get := way.Get(t1.Table()).
		Select(t1.EMAIL, t1.USERNAME).
		Where(func(f hey.Filter) {
			f.Equal(t1.ID, 1)
		}).
		Limit(1).
		Desc(t1.ID)

	{
		// handle object add, del, mod, get ...
		// TODO ...
	}

	// Generate SQL statements first.
	prepare1, args1 := add.Cmd()
	prepare2, args2 := del.Cmd()
	prepare3, args3 := mod.Cmd()
	prepare4, args4 := get.Cmd()

	var had *model.Account

	err := way.Transaction(nil, func(tx *hey.Way) error {
		rows, err := tx.Exec(prepare1, args1...)
		if err != nil {
			return err
		}
		if rows == 0 {
			return errors.New("insert failed")
		}

		had, err = hey.RowsScanStructOne[model.Account](
			nil, tx,
			func(rows *sql.Rows, v *model.Account) error {
				return rows.Scan(&v.Email, &v.Username /*, .... */)
			},
			prepare4,
			args4...,
		)
		if err != nil {
			if errors.Is(err, hey.ErrNoRows) {
				// todo
				// return errors.New("record does not exists")
			} else {
				return err
			}
		}
		{
			// use had todo ...
			// Replace the parameters in the following SQL statement to be executed ?
			_ = had
		}

		rows, err = tx.Exec(prepare2, args2...)
		if err != nil {
			return err
		}
		if rows == 0 {
			return errors.New("delete failed")
		}

		rows, err = tx.Exec(prepare3, args3...)
		if err != nil {
			return err
		}
		if rows == 0 {
			return errors.New("update failed")
		}

		return nil
	})
	if err != nil {
		return err
	}

	// TODO other logic codes.

	return nil
}

func (s *DB) RunTest() error {
	// INSERT
	_ = s.InsertOne()
	_, _ = s.Insert()
	_, _ = s.InsertWithColumn()
	_, _ = s.InsertWithMap()
	_, _ = s.InsertWithCustom()
	_, _ = s.InsertWithQuery()

	// Filter WHERE | HAVING
	s.Filter()

	// DELETE
	_, _ = s.Delete()

	// UPDATE
	_, _ = s.Update()

	// SELECT
	s.Select()

	// TRANSACTION
	_ = s.Transaction(context.Background())

	// TRANSACTION Optimized transaction calls.
	_ = s.Transaction1()

	return nil
}
