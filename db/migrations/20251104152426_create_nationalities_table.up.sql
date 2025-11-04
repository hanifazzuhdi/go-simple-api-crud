CREATE TABLE nationalities
(
    nationality_id   SERIAL PRIMARY KEY,
    nationality_name VARCHAR(50) NOT NULL,
    nationality_code CHAR(2)     NOT NULL
);