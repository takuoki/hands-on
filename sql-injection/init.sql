CREATE TABLE  employees (
    id INTEGER NOT NULL,
    name VARCHAR(64) NOT NULL,
    position VARCHAR(64),
    department VARCHAR(64),
    salary NUMERIC(10,2),
    secret BOOLEAN,
    PRIMARY KEY (id)
);

INSERT INTO employees VALUES
(1, '大庭 勇吉', '一般社員', '総務', 300000.00, FALSE),
(2, '野崎 竜夫', '係長', '経理', 400000.00, FALSE),
(3, '生田 信吉', '課長', '総務', 500000.00, FALSE),
(4, '浜野 一子', '部長', '経理', 820000.00, TRUE),
(5, '川崎 知治', '部長', '人事', 800000.00, TRUE);
