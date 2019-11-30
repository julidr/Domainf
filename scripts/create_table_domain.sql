CREATE TABLE domain (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    host STRING(100) NOT NULL,
    servers STRING[],
    ssl_grade STRING(5) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);