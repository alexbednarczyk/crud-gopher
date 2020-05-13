CREATE TABLE
IF NOT EXISTS users
(
    guid uuid,
    display_name varchar
(255),
    first_name varchar
(255),    
    last_name varchar
(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ);