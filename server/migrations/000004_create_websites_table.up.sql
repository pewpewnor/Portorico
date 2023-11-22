CREATE TABLE IF NOT EXISTS websites (
    name VARCHAR(64) UNIQUE,
    template_name VARCHAR(255),
    description VARCHAR(170),
    visitors_this_month INTEGER,
    content JSON,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
) INHERITS (base);