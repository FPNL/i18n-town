--每次改 scheme 不一定能對上程式裡的 struct
CREATE TABLE word
(
    id   SERIAL PRIMARY KEY,
    tag  VARCHAR(20),
    lang VARCHAR(10),
    word text,
    UNIQUE (tag, lang)
);
