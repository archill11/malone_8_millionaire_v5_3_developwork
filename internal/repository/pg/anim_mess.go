package pg

import (
	"encoding/json"
	"fmt"
	"myapp/internal/entity"
)

func (s *Database) AddNewAminMess(txt_id, txt_mess string) error {
	q := `
		INSERT INTO anim_mess (txt_id, txt_mess)
			VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`
	_, err := s.Exec(q, txt_id, txt_mess)
	if err != nil {
		return fmt.Errorf("AddNewAminMess Exec err: %s", err)
	}
	return nil
}

func (s *Database) GetAminMessByTxtId(txt_id string) (entity.AnimMess, error) {
	q := `
		SELECT coalesce((
			SELECT to_json(c)
	  		FROM anim_mess as c
	  		WHERE txt_id = $1
		), '{}'::json)
	`
	var u entity.AnimMess
	var data []byte
	err := s.QueryRow(q, txt_id).Scan(&data)
	if err != nil {
		return u, fmt.Errorf("GetAminMessByTxtId Scan: %v", err)
	}
	if err := json.Unmarshal(data, &u); err != nil {
		return u, fmt.Errorf("GetAminMessByTxtId Unmarshal: %v", err)
	}
	return u, nil
}

func (s *Database) EditAnimMessText(txt_id, txt_mess string) error {
	q := `UPDATE anim_mess SET txt_mess = $1 WHERE txt_id = $2`
	_, err := s.Exec(q, txt_mess, txt_id)
	if err != nil {
		return fmt.Errorf("EditAnimMessText Exec err: %v", err)
	}
	return nil
}
