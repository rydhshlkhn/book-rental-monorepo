-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS lending (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"book_item_id" INT,
	"user_id" INT,
	"due_date" TIMESTAMPTZ(6),
	"return_date" TIMESTAMPTZ(6),
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lending;
-- +goose StatementEnd
