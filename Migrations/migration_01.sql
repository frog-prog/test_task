CREATE TABLE IF NOT EXISTS tasks (
    id INT generated always as identity,
    text TEXT,
    done  BOOLEAN default false,
    date DATE,
    PRIMARY KEY (id)
)