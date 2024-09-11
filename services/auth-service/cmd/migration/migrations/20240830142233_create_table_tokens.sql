-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tokens (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"field" VARCHAR(255),
	"title"       VARCHAR(255) NOT NULL,
    "author"      VARCHAR(255) NOT NULL,
    "genre"       VARCHAR(100),
    "isbn"        VARCHAR(20) NOT NULL,
    "available"   BOOLEAN NOT NULL,
    "rent_price"  NUMERIC(10, 2) NOT NULL
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tokens;
-- +goose StatementEnd
