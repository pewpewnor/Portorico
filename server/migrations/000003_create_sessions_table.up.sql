CREATE TABLE IF NOT EXISTS sessions (
    token CHAR(32) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id)
) INHERITS (base);