CREATE TABLE family_list
(
    fl_id       SERIAL PRIMARY KEY,
    cst_id      INT         NOT NULL,
    fl_relation VARCHAR(50) NOT NULL,
    fl_name     VARCHAR(50) NOT NULL,
    fl_dob      DATE        NOT NULL,
    FOREIGN KEY (cst_id) REFERENCES customers (cst_id) ON DELETE CASCADE ON UPDATE CASCADE
);