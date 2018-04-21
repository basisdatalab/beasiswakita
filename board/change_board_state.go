package board

import (
	"errors"

	"github.com/harkce/beasiswakita"
)

func ChangeBoardState(ID int, state int) error {
	var board beasiswakita.StudentBoard

	err := beasiswakita.DbMap.SelectOne(&board, "select * from student_boards where id = ?", ID)
	if err != nil {
		return err
	}

	board.State = state
	err = board.Validate()
	if err != nil {
		return err
	}

	col, err := beasiswakita.Transaction.Update(&board)
	if err != nil {
		return err
	}

	if col == 0 {
		return errors.New("Board not found")
	}

	return nil
}
