CREATE TABLE IF NOT EXISTS websites (
    name VARCHAR(255) UNIQUE,
    template_name VARCHAR(255),
    description TEXT,
    visitors_this_month INTEGER,
    content JSON,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
) INHERITS (base);