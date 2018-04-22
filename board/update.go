package board

import (
	"errors"
	"time"

	"github.com/harkce/beasiswakita"
)

func UpdateBoard(board beasiswakita.StudentBoard) (beasiswakita.StudentBoard, error) {
	err := board.Validate()
	if err != nil {
		return board, err
	}

	var currentBoard beasiswakita.StudentBoard
	err = beasiswakita.DbMap.SelectOne(&currentBoard, "select * from student_boards where id = ?", board.ID)
	if err != nil {
		return board, err
	}

	currentBoard.Name = board.Name
	currentBoard.Category = board.Category
	currentBoard.Description = board.Description
	currentBoard.UpdatedAt = time.Now()

	col, err := beasiswakita.Transaction.Update(&currentBoard)
	if err != nil {
		return board, err
	}

	if col == 0 {
		return board, errors.New("Board not found")
	}

	return currentBoard, nil
}
