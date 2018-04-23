package scholarship

import (
	"github.com/harkce/beasiswakita"
)

func CreateBoard(s beasiswakita.Scholarship) beasiswakita.StudentBoard {
	var board beasiswakita.StudentBoard

	board.Name = s.Name
	board.Category = 1
	board.Description = s.Description
	board.State = 1
	board.ScholarshipID = &s.ID

	return board
}
