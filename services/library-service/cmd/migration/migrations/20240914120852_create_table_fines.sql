-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS fine (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"lending_id" INT REFERENCES lending("id"),
	"amount" INT,
	"snap_url" VARCHAR(255),
	"payment_status" VARCHAR(255),
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS fine;
-- +goose StatementEnd
