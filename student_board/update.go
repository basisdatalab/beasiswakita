package student_board

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

	board.UpdatedAt = time.Now()

	col, err := beasiswakita.Transaction.Update(&board)
	if err != nil {
		return board, err
	}

	if col == 0 {
		return board, errors.New("Board not found")
	}

	return board, nil
}
