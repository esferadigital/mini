-- +goose Up
-- +goose StatementBegin
CREATE TABLE visit (
  id BIGSERIAL PRIMARY KEY,
  short_url_id BIGINT NOT NULL,
  occurred_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  user_agent TEXT,

  CONSTRAINT fk_visit_short_url
    FOREIGN KEY (short_url_id)
    REFERENCES short_url(id)
    ON DELETE CASCADE
);

-- visit lookups for a given short_url
CREATE INDEX idx_visit_short_url_id ON visit(short_url_id);

-- visit lookups for a given short_url, but time-ordered
CREATE INDEX idx_visit_short_url_occurred_at ON visit(short_url_id, occurred_at DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS visit;
-- +goose StatementEnd
