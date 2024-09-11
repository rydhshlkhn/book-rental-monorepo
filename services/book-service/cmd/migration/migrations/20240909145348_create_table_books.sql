-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "isbn" VARCHAR(13) NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "subject" VARCHAR(255),
    "publisher" VARCHAR(255),
    "language" VARCHAR(100),
    "number_of_page" INT,
    "created_at" TIMESTAMPTZ(6),
    "updated_at" TIMESTAMPTZ(6)
);

CREATE TABLE IF NOT EXISTS book_formats (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMPTZ(6),
    "updated_at" TIMESTAMPTZ(6)
);

CREATE TABLE IF NOT EXISTS book_statuses (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMPTZ(6),
    "updated_at" TIMESTAMPTZ(6)
);

CREATE TABLE IF NOT EXISTS book_items (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "barcode" VARCHAR(255) NOT NULL,
    "is_reference_only" BOOLEAN NOT NULL,
    "borrowed" TIMESTAMPTZ(6),
    "due_date" TIMESTAMPTZ(6),
    "format_id" INT REFERENCES book_formats("id"),
    "status_id" INT REFERENCES book_statuses("id"),
    "date_of_purchase" TIMESTAMPTZ(6),
    "publication_date" TIMESTAMPTZ(6),
    "created_at" TIMESTAMPTZ(6),
    "updated_at" TIMESTAMPTZ(6),
    "book_id" INT,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE SET NULL ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS book_items;
DROP TABLE IF EXISTS book_formats;
DROP TABLE IF EXISTS book_statuses;
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
