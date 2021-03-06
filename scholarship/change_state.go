package scholarship

import (
	"errors"
	"time"

	"github.com/basisdatalab/beasiswakita"
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

	result, err := beasiswakita.Transaction.Exec("update scholarships set state = ? where id = ?", state, ID)
	if err != nil {
		return err
	}

	col, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if col == 0 {
		return errors.New("Scholarship not found")
	}

	return nil
}
