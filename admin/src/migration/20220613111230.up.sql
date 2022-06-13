--每次改 scheme 不一定能對上程式裡的 struct
CREATE TABLE ami
(
    id SERIAL PRIMARY KEY,
    organize VARCHAR(20),
    nickname VARCHAR(20),
    username VARCHAR(20),
    password CHAR(20) --此專案尚未針對資安
);
