-- +goose Up
-- +goose StatementBegin

-- +goose StatementEnd

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE PRODUCTS (
    ID UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    PRODUCT_CODE VARCHAR NOT NULL,
    NAME VARCHAR NOT NULL,
    QUANTITY INTEGER,
    PRICE NUMERIC(12,2) NULL DEFAULT 0,
    DISCOUNT NUMERIC(12,2) NULL DEFAULT 0,
    CREATED_BY UUID,
    CREATED_AT timestamp with time zone DEFAULT now() NOT NULL,
    UPDATED_AT timestamp with time zone DEFAULT now() NOT NULL,
    DELETED_AT timestamp with time zone,
    IS_DELETED boolean DEFAULT false
);

CREATE INDEX IDX_PRODUCTS_ID ON PRODUCTS(ID);
CREATE INDEX IDX_PRODUCTS_CREATED_AT ON PRODUCTS(CREATED_AT);
CREATE INDEX IDX_PRODUCTS_PRODUCT_CODE_AT ON PRODUCTS(PRODUCT_CODE);

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS PRODUCTS;
-- +goose StatementEnd
