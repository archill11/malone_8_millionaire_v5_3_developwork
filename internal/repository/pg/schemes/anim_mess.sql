CREATE TABLE IF NOT EXISTS anim_mess (
    id       SERIAL,
    txt_id   TEXT   DEFAULT '',
    txt_mess TEXT   DEFAULT '',

    PRIMARY KEY (id, txt_id)
);