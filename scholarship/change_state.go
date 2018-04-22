package scholarship

import (
	"errors"
	"time"

	"github.com/harkce/beasiswakita"
)

func ChangeState(ID int, state int) error {
	var s beasiswakita.Scholarship

	err := beasiswakita.DbMap.SelectOne(&s, "select * from scholarships where id = ?", ID)
	if err != nil {
		return err
	}

	s.State = state
	err = s.Validate()
	if err != nil {
		return err
	}

	s.UpdatedAt = time.Now()

	col, err := beasiswakita.Transaction.Update(&s)
	if err != nil {
		return err
	}

	if col == 0 {
		return errors.New("Scholarship not found")
	}

	return nil
}
