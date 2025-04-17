package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cd365/hey-example/db/model"
	"github.com/cd365/hey-example/db/model/biz"
	"github.com/cd365/hey-example/db/model/tab"
	"github.com/cd365/hey/v3"
	"github.com/cd365/logger/v8"
	_ "github.com/lib/pq"
	"os"
	"time"
)

var (
	way        *hey.Way
	schema     *tab.Schema
	executeDdl bool // Whether to execute the DDL for deleting/creating tables ?
)

var initSql = `DROP TABLE IF EXISTS company;
CREATE TABLE IF NOT EXISTS company (
    id serial NOT NULL,
    pid int NOT NULL DEFAULT 0,
    name varchar(128) NOT NULL DEFAULT '',
    country varchar(128) NOT NULL DEFAULT '',
    city varchar(128) NOT NULL DEFAULT '',
    address varchar(128) NOT NULL DEFAULT '',
    logo varchar(255) NOT NULL DEFAULT '',
    state int NOT NULL DEFAULT 0,
    remark text NOT NULL DEFAULT '',
    created_at bigint NOT NULL DEFAULT 0,
    updated_at bigint NOT NULL DEFAULT 0,
    deleted_at bigint NOT NULL DEFAULT 0,
    CONSTRAINT company_primary_key PRIMARY KEY (id)
) TABLESPACE pg_default;
ALTER TABLE IF EXISTS company OWNER to postgres;
COMMENT ON TABLE company IS 'company';
COMMENT ON COLUMN company.id IS 'id';
COMMENT ON COLUMN company.pid IS 'pid';
COMMENT ON COLUMN company.name IS 'name';
COMMENT ON COLUMN company.country IS 'country';
COMMENT ON COLUMN company.city IS 'city';
COMMENT ON COLUMN company.address IS 'address';
COMMENT ON COLUMN company.logo IS 'logo';
COMMENT ON COLUMN company.state IS 'state';
COMMENT ON COLUMN company.remark IS 'remark';
COMMENT ON COLUMN company.created_at IS 'created_at';
COMMENT ON COLUMN company.updated_at IS 'updated_at';
COMMENT ON COLUMN company.deleted_at IS 'deleted_at';
DROP TABLE IF EXISTS employee;
CREATE TABLE IF NOT EXISTS employee (
   id serial NOT NULL,
   company_id int NOT NULL DEFAULT 0,
   name varchar(32) NOT NULL DEFAULT '',
   age int NOT NULL DEFAULT 0,
   birthday varchar(10) NOT NULL DEFAULT '0000-00-00',
   gender varchar(16) NOT NULL DEFAULT 'unknown',
   height decimal(5,2) NOT NULL DEFAULT 0,
   weight decimal(5,2) NOT NULL DEFAULT 0,
   health decimal(5,2) NOT NULL DEFAULT 0,
   salary decimal(5,2) NOT NULL DEFAULT 0,
   department varchar(32) NOT NULL DEFAULT '',
   state int NOT NULL DEFAULT 0,
   remark text NOT NULL DEFAULT '',
   created_at bigint NOT NULL DEFAULT 0,
   updated_at bigint NOT NULL DEFAULT 0,
   deleted_at bigint NOT NULL DEFAULT 0,
   CONSTRAINT employee_primary_key PRIMARY KEY (id)
) TABLESPACE pg_default;
ALTER TABLE IF EXISTS employee OWNER to postgres;
COMMENT ON TABLE employee IS 'employee';
COMMENT ON COLUMN employee.id IS 'id';
COMMENT ON COLUMN employee.company_id IS 'company_id';
COMMENT ON COLUMN employee.name IS 'name';
COMMENT ON COLUMN employee.age IS 'age';
COMMENT ON COLUMN employee.birthday IS 'birthday';
COMMENT ON COLUMN employee.gender IS 'gender unknown OR male OR female';
COMMENT ON COLUMN employee.height IS 'height unit: cm';
COMMENT ON COLUMN employee.weight IS 'weight unit: kg';
COMMENT ON COLUMN employee.health IS 'health value';
COMMENT ON COLUMN employee.salary IS 'salary';
COMMENT ON COLUMN employee.department IS 'department';
COMMENT ON COLUMN employee.state IS 'state';
COMMENT ON COLUMN employee.remark IS 'remark';
COMMENT ON COLUMN employee.created_at IS 'created_at';
COMMENT ON COLUMN employee.updated_at IS 'updated_at';
COMMENT ON COLUMN employee.deleted_at IS 'deleted_at';`

func initDatabase() error {
	driverName := "postgres"
	dataSourceName := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", "postgres", "postgres", "hey")
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(8)
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetConnMaxLifetime(time.Minute * 3)
	way = hey.NewWay(db)

	cfg := way.GetCfg()
	// Use the Helper for the corresponding database.
	cfg.Helper = hey.NewPostgresHelper(driverName, dataSourceName)

	// Record SQL log.
	lg := logger.NewLogger(nil)
	way.SetLogger(lg)

	// Customize SQL execution time warning threshold.
	cfg.WarnDuration = time.Millisecond * 200

	way.SetCfg(*cfg)

	// Execute create table DDL.
	if executeDdl {
		if _, err = db.Exec(initSql); err != nil {
			return err
		}
	}

	dbSchema, err := model.NewDatabase(context.Background(), way, func(db *model.Database) error {
		return err
	})
	bizSchema, err := biz.NewSchema(dbSchema, func(db *model.Database, schema *biz.Schema) error {
		return nil
	})
	schema = tab.NewSchema(bizSchema)

	// Delete specific identifiers in SQL statements.
	cfg = way.GetCfg()
	cfg.Replacer.DelAll()
	way.SetCfg(*cfg)

	return nil
}

func init() {
	if err := initDatabase(); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func main() {
	Insert()
	InsertOne()
	InsertFromQuery()
	Filter()
	Delete()
	Update()
	Select()
	Transaction()
}

func Insert() {
	t := schema.Company
	m := t.Model()
	now := way.Now()
	creates := []*model.Company{
		{
			Name:      "abc",
			State:     1,
			CreatedAt: now.Unix(),
			UpdatedAt: now.Unix(),
		},
		{
			Name:      "def",
			State:     1,
			CreatedAt: now.Unix(),
			UpdatedAt: now.Unix(),
		},
	}

	{
		add := way.Add(m.Table()).Comment("create one or more rows").ExceptPermit(func(except hey.UpsertColumns, permit hey.UpsertColumns) {
			except.Add(m.PrimaryKey()) /* ignore primary key */
		}).Create(creates)
		way.Debugger(add)
		// add.Add() /* execute INSERT */
	}

}

func InsertOne() {
	t := schema.Company
	m := t.Model()
	now := way.Now()

	create := &model.Company{
		Name:      "abc",
		State:     1,
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
	}

	{
		// INSERT ONE RETURNING primary key value.
		add := way.Add(m.Table()).
			Comment("create one return primary key value").
			ExceptPermit(func(except hey.UpsertColumns, permit hey.UpsertColumns) {
				except.Add(m.PrimaryKey()) /* ignore primary key */
			}).
			Create(create)
		id, err := add.AddOne(func(cmder hey.Cmder) hey.Cmder {
			prepare, args := cmder.Cmd()
			cmder = hey.NewCmder(fmt.Sprintf("%s RETURNING %s", prepare, m.PrimaryKey()), args)
			// way.Debugger(cmder)
			return cmder
		}, func(ctx context.Context, stmt *hey.Stmt, args []interface{}) (id int64, err error) {
			err = stmt.QueryRow(func(rows *sql.Row) error { return rows.Scan(&id) }, args...)
			return
		})
		if err != nil {
			return
		}
		fmt.Println(id)
	}

}

func InsertFromQuery() {
	t := schema.Company
	m := t.Model()
	filter := way.F().GreaterThanEqual(m.ID, 1)
	get := way.Get(m.Table()).Select(m.NAME, m.CITY).Where(t.Filter(filter)).Asc(m.ID).Limit(5)
	add := way.Add(m.Table()).Comment("insert query results into the table").CmderValues(get, []string{m.NAME, m.CITY})
	way.Debugger(add)
}

// Filter is often used in WHERE and HAVING.
func Filter() {
	t := schema.Employee
	m := t.Model()

	// Simple conditional filtering and splicing.
	f := way.F()
	f.Equal(m.GENDER, "female").
		Between(m.AGE, 18, 20).
		GreaterThan(m.HEIGHT, 160).
		Between(m.WEIGHT, 90, 110)
	way.Debugger(f)

	// Implementing complex conditional filtering logic.
	f = way.F()
	f.Equal(m.GENDER, "male")
	f.Group(func(g hey.Filter) {
		g.Group(func(g hey.Filter) { g.Between(m.AGE, 18, 20).GreaterThan(m.HEIGHT, 170) })
		g.OrGroup(func(g hey.Filter) { g.Between(m.WEIGHT, 115, 130).GreaterThan(m.HEIGHT, 175) })
	})
	f.IsNotNull(m.REMARK)
	f.Like(m.NAME, "A%")
	way.Debugger(f)

	// column1 IN ( SELECT column1 FROM xxx [ WHERE ] [ ORDER BY ] [ LIMIT ] )
	f = way.F()
	f.Equal(m.GENDER, "male")
	queryFilter := way.F()
	queryFilter.Between(m.WEIGHT, 115, 130).GreaterThan(m.HEIGHT, 175)
	f.InQuery(m.AGE, m.Get().Select(m.AGE).Where(t.Filter(queryFilter)).Desc(m.ID).Limit(100))
	way.Debugger(f)

	// ( column1, column2, column3 ) IN ( SELECT column1, column2, column3 FROM xxx [ WHERE ] [ ORDER BY ] [ LIMIT ] )
	f = way.F()
	f.InColsQuery([]string{m.NAME, m.AGE, m.GENDER}, m.Get().Select(m.NAME, m.AGE, m.GENDER).Where(t.Filter(queryFilter)).Desc(m.ID).Limit(20))
	way.Debugger(f)
}

func Delete() {
	t := schema.Employee
	m := t.Model()

	filter := way.F()
	filter.Between(m.ID, 2, 5)
	// way.Debugger(filter)
	del := way.Del(m.Table()).Comment("delete data").Where(func(where hey.Filter) { where.Use(filter) })
	way.Debugger(del)
	// del.Del() /* execute DELETE */

	// For complex WHERE conditions, refer to Filter()

	// For executing DELETE statements, locking is not recommended. It is recommended to use optimistic locking.
}

func Update() {
	t := schema.Employee
	m := t.Model()

	filter := way.F()
	filter.Equal(m.ID, 8)
	mod := way.Mod(m.Table()).Comment("update data").Where(func(where hey.Filter) { where.Use(filter) }).Update(map[string]interface{}{
		m.NAME: "Jack",
	})
	way.Debugger(mod)
	// mod.Mod() /* execute UPDATE */

	mod = way.Mod(m.Table()).Comment("update data").Where(func(where hey.Filter) { where.Use(filter) }).Update(m.ColumnValue(m.NAME, "JACK", m.AGE, 21))
	way.Debugger(mod)
	// mod.Mod() /* execute UPDATE */

	mod.ExceptPermit(func(except hey.UpsertColumns, permit hey.UpsertColumns) {
		// Set a list of fields that are not allowed to be updated.
		except.Add(m.PrimaryKey(), m.CREATED_AT, m.DELETED_AT)
		except.Add(m.ColumnCreatedAt()...)
		except.Add(m.ColumnDeletedAt()...)

		// You can also set a list of fields that are only allowed to be updated.
		// permit.Add(m.NAME, m.AGE, m.WEIGHT, m.HEIGHT, m.HEALTH)
		permit.Add(m.NAME, m.AGE, m.WEIGHT, m.HEIGHT, m.HEALTH, m.SALARY)

		// Special note: It is not recommended to use a and b at the same time.
	})

	// salary = salary + 1
	mod.Incr(m.SALARY, 1)

	mod.Set(m.HEALTH, 90.05)
	way.Debugger(mod)

	// For complex WHERE conditions, refer to Filter()

	// For executing UPDATE statements, locking is not recommended. It is recommended to use optimistic locking.
}

func Select() {
	t := schema.Employee
	m := t.Model()
	t1 := schema.Company
	m1 := t1.Model()

	// Plan A for query.
	{
		a := hey.AliasA
		b := hey.AliasB
		get := way.Get(m.Table()).Alias(a).Comment("query data plan A").Join(func(j hey.QueryJoin) {
			r1 := j.NewTable(m1.Table(), b)
			j.LeftJoin(nil, r1, j.OnEqual(fmt.Sprintf("%s.%s", a, m.COMPANY_ID), fmt.Sprintf("%s.%s", b, m1.ID)))
			j.Queries().AddAll(fmt.Sprintf("b.%s AS %s", m1.NAME, "company_name"))
			j.Queries().AddAll("a.*")
		})
		get.Desc("a.id")
		get.Limit(100)
		way.Debugger(get)
	}

	// Plan B for query, it is strongly recommended to avoid using quoted strings as much as possible.
	{
		a := m.AliasA()
		b := m1.AliasB()
		tc := way.T()
		get := way.Get(a.Table()).Alias(a.Alias()).Comment("query data plan B").Join(func(j hey.QueryJoin) {
			r0 := j.GetMaster()
			r1 := j.NewTable(b.Table(), b.Alias())
			j.LeftJoin(r0, r1, j.OnEqual(a.COMPANY_ID, b.ID))
			// j.Queries().AddAll("a.*")
			j.Queries().AddAll(a.Column()...).
				AddAll(tc.Column(b.NAME, "company_name"))
		})
		get.Desc(a.ID)
		get.Limit(100)
		way.Debugger(get)
	}

	{
		filter := way.F()
		filter.GreaterThan(m.ID, 0)

		// SELECT ONE ROW
		employee, err := m.SelectOne(filter, func(get *hey.Get) {
			get.Comment("select one row")
		})
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", employee)

		// SELECT MORE ROWS
		employees, err := m.SelectAll(filter, func(get *hey.Get) {
			get.Comment("select more rows")
			get.Desc(m.ID).Limit(10)
		})
		if err != nil {
			return
		}
		for _, tmp := range employees {
			fmt.Printf("%+v\n", tmp)
		}
	}

}

func Transaction() {
	t := schema.Employee
	m := t.Model()
	t1 := schema.Company
	m1 := t1.Model()

	// Set the transaction timeout period.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var msg error
	err := way.Transaction(ctx, func(tx *hey.Way) error {
		tx.TransactionMessage("try transaction")
		// Method A
		rows, err := tx.Mod(m.Table()).Comment("update 1 in a transaction").Where(t.Filter(m.Filter().Equal(m.ID, 1))).Incr(m.AGE, 1).Mod()
		if err != nil {
			return err
		}
		if rows == 0 {
			msg = fmt.Errorf("`%s` update failed", m.Table())
			return msg
		}

		// Method B
		rows, err = t1.Update(tx, func(f hey.Filter, u *hey.Mod) {
			f.Equal(m1.ID, 1)
			u.Comment("update 2 in a transaction")
			u.Set(m1.CITY, "Tokyo")
		})
		if err != nil {
			return err
		}
		if rows == 0 {
			msg = fmt.Errorf("`%s` update failed", m1.Table())
			return msg
		}

		// Compare method A and method B, which one is easier to write ?
		return nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}

}
