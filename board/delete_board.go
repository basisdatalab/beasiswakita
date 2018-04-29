package board

import (
	"errors"

	"github.com/basisdatalab/beasiswakita"
)

func DeleteBoard(board beasiswakita.StudentBoard) (beasiswakita.StudentBoard, error) {
	err := board.Validate()
	if err != nil {
		return board, err
	}

	var currentBoard beasiswakita.StudentBoard
	err = beasiswakita.DbMap.SelectOne(&currentBoard, "select * from student_boards where id = ?", board.ID)
	if err != nil {
		return board, err
	}

	col, err := beasiswakita.Transaction.Delete(&board)
	if err != nil {
		return board, err
	}

	if col == 0 {
		return board, errors.New("Board not found")
	}

	return currentBoard, nil
}
