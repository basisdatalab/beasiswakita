package beasiswakita

import (
	"errors"
	"time"
)

var scholarshipCategory = map[int]string{
	1: "Beasiswa Penuh",
	2: "Beasiswa Atlet",
}

var boardState = map[int]string{
	1: "Whislist",
	2: "Applied",
	3: "In Progress",
	4: "Rejected",
	5: "Accepted",
}

type StudentBoard struct {
	ID            int       `db:"id" json:"id"`
	Name          string    `db:"name" json:"name"`
	Category      int       `db:"category" json:"category"`
	Description   string    `db:"description" json:"description"`
	State         int       `db:"state" json:"state"`
	UserID        int       `db:"user_id" json:"user_id"`
	ScholarshipID *int      `db:"scholarship_id" json:"scholarship_id"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

func (s *StudentBoard) Validate() error {
	if _, ok := scholarshipCategory[s.Category]; !ok {
		return errors.New("Category not found")
	}

	if _, ok := boardState[s.State]; !ok {
		return errors.New("Invalid state")
	}

	return nil
}
