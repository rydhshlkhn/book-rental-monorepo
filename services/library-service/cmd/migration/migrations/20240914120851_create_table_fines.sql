-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS fine (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"book_item_id" INT,
	"user_id" INT,
	"amount" INT,
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS fine;
-- +goose StatementEnd
