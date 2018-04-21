package student_board

import (
	"time"

	"github.com/harkce/beasiswakita"
)

func CreateBoard(board beasiswakita.StudentBoard) (beasiswakita.StudentBoard, error) {
	err := board.Validate()
	if err != nil {
		return board, err
	}

	board.CreatedAt = time.Now()
	board.UpdatedAt = time.Now()

	err = beasiswakita.Transaction.Insert(&board)
	if err != nil {
		return board, err
	}

	return board, nil
}
