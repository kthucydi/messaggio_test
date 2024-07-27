-- +goose Up
CREATE TABLE messages (
   id bigserial PRIMARY KEY, -- message ID
   message text NOT NULL, -- message body
   processed BOOLEAN DEFAULT 0 -- processed flag
);

-- +goose Down
DROP TABLE messages;
