/* account (account comment) */
DROP TABLE IF EXISTS "account";
CREATE SEQUENCE IF NOT EXISTS account_id_seq START 1;
CREATE TABLE IF NOT EXISTS "account" (
  "id" integer NOT NULL DEFAULT nextval('account_id_seq'::regclass),
  "email" character varying(100) NOT NULL,
  "username" character varying(50) NOT NULL,
  "balance" numeric NOT NULL DEFAULT 0,
  "password" character varying(64) NOT NULL DEFAULT ''::character varying,
  "status" integer NOT NULL DEFAULT 0,
  "ip" character varying(39) NOT NULL DEFAULT ''::character varying,
  "created_at" bigint NOT NULL DEFAULT 0,
  "updated_at" bigint NOT NULL DEFAULT 0,
  "deleted_at" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  CONSTRAINT "account_email_key" UNIQUE (email),
  CONSTRAINT "account_username_key" UNIQUE (username)
);
CREATE INDEX IF NOT EXISTS account_status ON account USING btree (status);
CREATE INDEX IF NOT EXISTS account_created_at ON account USING btree (created_at);
COMMENT ON TABLE "account" IS 'account comment';
COMMENT ON COLUMN "account"."id" IS 'id comment';
COMMENT ON COLUMN "account"."email" IS 'email comment';
COMMENT ON COLUMN "account"."username" IS 'username comment';
COMMENT ON COLUMN "account"."balance" IS 'balance comment';
COMMENT ON COLUMN "account"."password" IS 'balance password';
COMMENT ON COLUMN "account"."status" IS 'status comment';
COMMENT ON COLUMN "account"."ip" IS 'ip comment';
COMMENT ON COLUMN "account"."created_at" IS 'created_at comment';
COMMENT ON COLUMN "account"."updated_at" IS 'updated_at comment';
COMMENT ON COLUMN "account"."deleted_at" IS 'deleted_at comment';



/* article (article comment) */
DROP TABLE IF EXISTS "article";
CREATE SEQUENCE IF NOT EXISTS article_id_seq START 1;
CREATE TABLE IF NOT EXISTS "article" (
  "id" integer NOT NULL DEFAULT nextval('article_id_seq'::regclass),
  "account_id" integer NOT NULL DEFAULT 0,
  "title" character varying(255) NOT NULL DEFAULT ''::character varying,
  "content" text NOT NULL DEFAULT ''::text,
  "stars" bigint NOT NULL DEFAULT 0,
  "created_at" bigint NOT NULL DEFAULT 0,
  "updated_at" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS article_account_id ON article USING btree (account_id);
CREATE INDEX IF NOT EXISTS article_title ON article USING btree (title);
CREATE INDEX IF NOT EXISTS account_stars ON article USING btree (stars);
CREATE INDEX IF NOT EXISTS article_created_at ON article USING btree (created_at);
COMMENT ON TABLE "article" IS 'article comment';
COMMENT ON COLUMN "article"."id" IS 'id comment';
COMMENT ON COLUMN "article"."account_id" IS 'account_id comment';
COMMENT ON COLUMN "article"."title" IS 'title comment';
COMMENT ON COLUMN "article"."content" IS 'content comment';
COMMENT ON COLUMN "article"."stars" IS 'stars comment';
COMMENT ON COLUMN "article"."created_at" IS 'created_at comment';
COMMENT ON COLUMN "article"."updated_at" IS 'updated_at comment';



/* article_comment (article_comment comment) */
DROP TABLE IF EXISTS "article_comment";
CREATE SEQUENCE IF NOT EXISTS article_comment_id_seq START 1;
CREATE TABLE IF NOT EXISTS "article_comment" (
  "id" integer NOT NULL DEFAULT nextval('article_comment_id_seq'::regclass),
  "account_id" integer NOT NULL DEFAULT 0,
  "article_id" integer NOT NULL DEFAULT 0,
  "content" text NOT NULL DEFAULT ''::text,
  "created_at" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS article_comment_account_id ON article_comment USING btree (account_id);
CREATE INDEX IF NOT EXISTS article_comment_article_id ON article_comment USING btree (article_id);
CREATE INDEX IF NOT EXISTS article_comment_created_at ON article_comment USING btree (created_at);
COMMENT ON TABLE "article_comment" IS 'article_comment comment';
COMMENT ON COLUMN "article_comment"."id" IS 'id comment';
COMMENT ON COLUMN "article_comment"."account_id" IS 'account_id comment';
COMMENT ON COLUMN "article_comment"."article_id" IS 'article_id comment';
COMMENT ON COLUMN "article_comment"."content" IS 'content comment';
COMMENT ON COLUMN "article_comment"."created_at" IS 'created_at comment';