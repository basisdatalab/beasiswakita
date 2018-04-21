package student_board

import (
	"errors"

	"github.com/harkce/beasiswakita"
)

func DeleteBoard(board beasiswakita.StudentBoard) (beasiswakita.StudentBoard, error) {
	col, err := beasiswakita.Transaction.Delete(&board)
	if err != nil {
		return board, err
	}

	if col == 0 {
		return board, errors.New("Board not found")
	}

	return board, nil
}
