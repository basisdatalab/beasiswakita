package scholarship

import (
	"errors"

	"github.com/harkce/beasiswakita"
)

func UpdateScholarship(s beasiswakita.Scholarship) (beasiswakita.Scholarship, error) {
	var currentS beasiswakita.Scholarship
	err := beasiswakita.DbMap.SelectOne(currentS, "select * from scholarships where id = ?", s.ID)
	if err != nil {
		return s, nil
	}

	currentS.Name = s.Name
	currentS.Country = s.Country
	currentS.Flag = s.Flag
	currentS.State = s.State
	currentS.StartDate = s.StartDate
	currentS.EndDate = s.EndDate
	currentS.Description = s.Description
	currentS.Requirement = s.Requirement
	currentS.UpdatedAt = s.UpdatedAt

	col, err := beasiswakita.Transaction.Update(&currentS)
	if err != nil {
		return s, err
	}

	if col == 0 {
		return s, errors.New("Scholarship not found")
	}

	return currentS, nil
}
