package board

import (
	"github.com/harkce/beasiswakita"
)

func GetBoards(ID int) ([]beasiswakita.StudentBoard, error) {
	var boards []beasiswakita.StudentBoard

	_, err := beasiswakita.DbMap.Select(&boards, "select * from student_boards where user_id = ?", ID)
	if err != nil {
		return boards, err
	}

	return boards, nil
}
