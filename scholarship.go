package beasiswakita

import (
	"time"
)

type Scholarship struct {
	ID             int       `db:"id" json:"id"`
	Name           string    `db:"name" json:"name"`
	Country        string    `db:"country" json:"country"`
	Flag           string    `db:"flag" json:"flag"`
	State          int       `db:"state" json:"state"`
	StartDate      string    `db:"start_date" json:"start_date"`
	EndDate        string    `db:"end_date" json:"end_date"`
	Description    string    `db:"description" json:"description"`
	Requirement    string    `db:"requirement" json:"requirement"`
	OrganizationID int       `db:"organization_id" json:"organization_id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

func (s *Scholarship) Validate() error {
	return nil
}
