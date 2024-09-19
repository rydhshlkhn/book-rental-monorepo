-- +goose Up
-- +goose StatementBegin

-- Create transactions table
CREATE TABLE IF NOT EXISTS transaction (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"transaction_time" TIMESTAMPTZ(6),
	"transaction_status" VARCHAR(50),
	"transaction_id" VARCHAR(100),
	"status_message" VARCHAR(255),
	"status_code" VARCHAR(10),
	"signature_key" TEXT,
	"settlement_time" TIMESTAMPTZ(6),
	"payment_type" VARCHAR(50),
	"order_id" VARCHAR(100),
	"merchant_id" VARCHAR(100),
	"gross_amount" DECIMAL(20, 2),
	"fraud_status" VARCHAR(50),
	"currency" VARCHAR(10),
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6)
);

-- Create va_numbers table (One-to-One with transactions)
CREATE TABLE IF NOT EXISTS va_number (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"transaction_id" INT REFERENCES transaction(id) ON DELETE CASCADE ON UPDATE CASCADE,
	"va_number" VARCHAR(50),
	"bank" VARCHAR(50),
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6)
);

-- Create payment_amounts table (One-to-One with transactions)
CREATE TABLE IF NOT EXISTS payment_amount (
	"id" SERIAL NOT NULL PRIMARY KEY,
	"transaction_id" INT REFERENCES transaction(id) ON DELETE CASCADE ON UPDATE CASCADE,
	"paid_at" TIMESTAMPTZ(6),
	"amount" DECIMAL(20, 2),
	"created_at" TIMESTAMPTZ(6),
	"updated_at" TIMESTAMPTZ(6)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop payment_amounts table
DROP TABLE IF EXISTS payment_amount;

-- Drop va_numbers table
DROP TABLE IF EXISTS va_number;

-- Drop transactions table
DROP TABLE IF EXISTS transaction;

-- +goose StatementEnd
