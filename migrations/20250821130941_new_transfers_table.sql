-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transfers (
    from_id UUID NOT NULL,
    to_id UUID NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transfers;
-- +goose StatementEnd
