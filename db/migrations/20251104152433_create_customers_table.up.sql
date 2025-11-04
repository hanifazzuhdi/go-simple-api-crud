CREATE TABLE customers
(
    cst_id         SERIAL PRIMARY KEY,
    cst_name       CHAR(50)    NOT NULL,
    cst_dob        DATE        NOT NULL,
    cst_phoneNum   varchar(20) NOT NULL,
    cst_email      varchar(50) NOT NULL,
    nationality_id INT         NOT NULL,
    FOREIGN KEY (nationality_id) REFERENCES nationalities (nationality_id) ON DELETE RESTRICT ON UPDATE CASCADE
);