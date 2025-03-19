/* company (company) */
DROP TABLE IF EXISTS "company";
CREATE SEQUENCE IF NOT EXISTS company_id_seq START 1;
CREATE TABLE IF NOT EXISTS "company" (
  "id" integer NOT NULL DEFAULT nextval('company_id_seq'::regclass),
  "pid" integer NOT NULL DEFAULT 0,
  "name" character varying(128) NOT NULL DEFAULT ''::character varying,
  "country" character varying(128) NOT NULL DEFAULT ''::character varying,
  "city" character varying(128) NOT NULL DEFAULT ''::character varying,
  "address" character varying(128) NOT NULL DEFAULT ''::character varying,
  "logo" character varying(255) NOT NULL DEFAULT ''::character varying,
  "state" integer NOT NULL DEFAULT 0,
  "remark" text NOT NULL DEFAULT ''::text,
  "created_at" bigint NOT NULL DEFAULT 0,
  "updated_at" bigint NOT NULL DEFAULT 0,
  "deleted_at" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);
COMMENT ON TABLE "company" IS 'company';
COMMENT ON COLUMN "company"."id" IS 'id';
COMMENT ON COLUMN "company"."pid" IS 'pid';
COMMENT ON COLUMN "company"."name" IS 'name';
COMMENT ON COLUMN "company"."country" IS 'country';
COMMENT ON COLUMN "company"."city" IS 'city';
COMMENT ON COLUMN "company"."address" IS 'address';
COMMENT ON COLUMN "company"."logo" IS 'logo';
COMMENT ON COLUMN "company"."state" IS 'state';
COMMENT ON COLUMN "company"."remark" IS 'remark';
COMMENT ON COLUMN "company"."created_at" IS 'created_at';
COMMENT ON COLUMN "company"."updated_at" IS 'updated_at';
COMMENT ON COLUMN "company"."deleted_at" IS 'deleted_at';



/* employee (employee) */
DROP TABLE IF EXISTS "employee";
CREATE SEQUENCE IF NOT EXISTS employee_id_seq START 1;
CREATE TABLE IF NOT EXISTS "employee" (
  "id" integer NOT NULL DEFAULT nextval('employee_id_seq'::regclass),
  "company_id" integer NOT NULL DEFAULT 0,
  "name" character varying(32) NOT NULL DEFAULT ''::character varying,
  "age" integer NOT NULL DEFAULT 0,
  "birthday" character varying(10) NOT NULL DEFAULT '0000-00-00'::character varying,
  "gender" character varying(16) NOT NULL DEFAULT 'unknown'::character varying,
  "height" numeric NOT NULL DEFAULT 0,
  "weight" numeric NOT NULL DEFAULT 0,
  "health" numeric NOT NULL DEFAULT 0,
  "salary" numeric NOT NULL DEFAULT 0,
  "department" character varying(32) NOT NULL DEFAULT ''::character varying,
  "state" integer NOT NULL DEFAULT 0,
  "remark" text NOT NULL DEFAULT ''::text,
  "created_at" bigint NOT NULL DEFAULT 0,
  "updated_at" bigint NOT NULL DEFAULT 0,
  "deleted_at" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);
COMMENT ON TABLE "employee" IS 'employee';
COMMENT ON COLUMN "employee"."id" IS 'id';
COMMENT ON COLUMN "employee"."company_id" IS 'company_id';
COMMENT ON COLUMN "employee"."name" IS 'name';
COMMENT ON COLUMN "employee"."age" IS 'age';
COMMENT ON COLUMN "employee"."birthday" IS 'birthday';
COMMENT ON COLUMN "employee"."gender" IS 'gender unknown OR male OR female';
COMMENT ON COLUMN "employee"."height" IS 'height unit: cm';
COMMENT ON COLUMN "employee"."weight" IS 'weight unit: kg';
COMMENT ON COLUMN "employee"."health" IS 'health value';
COMMENT ON COLUMN "employee"."salary" IS 'salary';
COMMENT ON COLUMN "employee"."department" IS 'department';
COMMENT ON COLUMN "employee"."state" IS 'state';
COMMENT ON COLUMN "employee"."remark" IS 'remark';
COMMENT ON COLUMN "employee"."created_at" IS 'created_at';
COMMENT ON COLUMN "employee"."updated_at" IS 'updated_at';
COMMENT ON COLUMN "employee"."deleted_at" IS 'deleted_at';