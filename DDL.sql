DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS article;
DROP TABLE IF EXISTS article_comment;

CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    balance DECIMAL(38,8) NOT NULL DEFAULT 0,
    password VARCHAR(64) NOT NULL DEFAULT '',
    status INT NOT NULL DEFAULT 0,
    ip VARCHAR(39) NOT NULL DEFAULT '',
    created_at BIGINT NOT NULL DEFAULT 0,
    updated_at BIGINT NOT NULL DEFAULT 0,
    deleted_at bigint NOT NULL DEFAULT 0
);
CREATE INDEX account_status ON account (status);
CREATE INDEX account_created_at ON account (created_at);
COMMENT ON COLUMN account.id IS 'id comment';
COMMENT ON COLUMN account.email IS 'email comment';
COMMENT ON COLUMN account.username IS 'username comment';
COMMENT ON COLUMN account.balance IS 'balance comment';
COMMENT ON COLUMN account.password IS 'balance password';
COMMENT ON COLUMN account.status IS 'status comment';
COMMENT ON COLUMN account.ip IS 'ip comment';
COMMENT ON COLUMN account.created_at IS 'created_at comment';
COMMENT ON COLUMN account.updated_at IS 'updated_at comment';
COMMENT ON COLUMN account.deleted_at IS 'deleted_at comment';
COMMENT ON TABLE account IS 'account comment';

CREATE TABLE IF NOT EXISTS article (
    id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL DEFAULT 0,
    title VARCHAR(255) NOT NULL DEFAULT '',
    content TEXT NOT NULL DEFAULT '',
    stars BIGINT NOT NULL DEFAULT 0,
    created_at BIGINT NOT NULL DEFAULT 0,
    updated_at BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX article_account_id ON article (account_id);
CREATE INDEX article_title ON article (title);
CREATE INDEX account_stars ON article (stars);
CREATE INDEX article_created_at ON article (created_at);
COMMENT ON COLUMN article.id IS 'id comment';
COMMENT ON COLUMN article.account_id IS 'account_id comment';
COMMENT ON COLUMN article.title IS 'title comment';
COMMENT ON COLUMN article.content IS 'content comment';
COMMENT ON COLUMN article.stars IS 'stars comment';
COMMENT ON COLUMN article.created_at IS 'created_at comment';
COMMENT ON COLUMN article.updated_at IS 'updated_at comment';
COMMENT ON TABLE article IS 'article comment';

CREATE TABLE IF NOT EXISTS article_comment (
    id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL DEFAULT 0,
    article_id INTEGER NOT NULL DEFAULT 0,
    content TEXT NOT NULL DEFAULT '',
    created_at BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX article_comment_account_id ON article_comment (account_id);
CREATE INDEX article_comment_article_id ON article_comment (article_id);
CREATE INDEX article_comment_created_at ON article_comment (created_at);
COMMENT ON COLUMN article_comment.id IS 'id comment';
COMMENT ON COLUMN article_comment.account_id IS 'account_id comment';
COMMENT ON COLUMN article_comment.article_id IS 'article_id comment';
COMMENT ON COLUMN article_comment.content IS 'content comment';
COMMENT ON COLUMN article_comment.created_at IS 'created_at comment';
COMMENT ON TABLE article_comment IS 'article_comment comment';
