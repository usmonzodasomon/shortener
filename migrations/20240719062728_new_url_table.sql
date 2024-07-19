-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS urls(
    id        SERIAL PRIMARY KEY ,
    url       TEXT NOT NULL,
    short_url TEXT NOT NULL,
    count     INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS urls;
-- +goose StatementEnd
