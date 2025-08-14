-- +goose Up
-- +goose StatementBegin
CREATE TABLE short_url (
    id BIGSERIAL PRIMARY KEY,
    slug VARCHAR(32) NOT NULL,
    target_url TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX idx_short_url_slug ON short_url(slug);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_short_url_slug;
DROP TABLE IF EXISTS short_url;
-- +goose StatementEnd
