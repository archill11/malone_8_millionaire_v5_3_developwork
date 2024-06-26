CREATE TABLE IF NOT EXISTS schemes (
    id           TEXT   DEFAULT '',
    sc_name      TEXT   DEFAULT '',
    sc_idx       INT    DEFAULT 0,
    lichka       TEXT DEFAULT '',
    lichka_id    INT  DEFAULT 0,
    link         TEXT DEFAULT '',
    chat_check_link TEXT DEFAULT '',
    chat_check_id INT DEFAULT 0,

    PRIMARY KEY (id)
);

-------------------------------------

ALTER TABLE schemes
  ADD COLUMN IF NOT EXISTS lichka_id INT DEFAULT 0;

ALTER TABLE schemes
  ADD COLUMN IF NOT EXISTS chat_check_id INT DEFAULT 0;

ALTER TABLE schemes
  ADD COLUMN IF NOT EXISTS chat_check_link TEXT DEFAULT '';