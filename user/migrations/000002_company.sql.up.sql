CREATE TABLE IF NOT EXISTS company(
    id bigserial PRIMARY KEY,
    name VARCHAR (150) UNIQUE NOT NULL
);