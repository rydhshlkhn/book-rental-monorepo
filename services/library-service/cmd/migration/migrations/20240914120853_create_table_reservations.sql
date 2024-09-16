-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reservation_status (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMPTZ(6),
    "updated_at" TIMESTAMPTZ(6)
);

CREATE TABLE IF NOT EXISTS reservation (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"book_item_id" INT,
	"status_id" INT REFERENCES reservation_status("id"),
	"user_id" INT,
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6),
	CONSTRAINT unique_book_item_user UNIQUE ("book_item_id", "user_id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reservation_status;
DROP TABLE IF EXISTS reservation;
-- +goose StatementEnd
