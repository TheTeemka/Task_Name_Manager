-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS people (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    surname VARCHAR(20) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL,
    nationality VARCHAR(10) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT unique_name_surname UNIQUE (name, surname)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS people;
-- +goose StatementEnd
