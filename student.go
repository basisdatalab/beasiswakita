package beasiswakita

import (
	"time"
)

type Student struct {
	ID            int       `db:"id" json:"id"`
	Name          string    `db:"name" json:"name"`
	Birthdate     time.Time `db:"birthdate" json:"birthdate"`
	Address       string    `db:"address" json:"address"`
	City          string    `db:"city" json:"city"`
	Region        string    `db:"region" json:"region"`
	Country       string    `db:"country" json:"country"`
	Zipcode       string    `db:"zipcode" json:"zipcode,omitempty"`
	Education     string    `db:"education" json:"education"`
	SchoolName    string    `db:"school_name" json:"school_name"`
	SchoolAddress string    `db:"school_address" json:"school_address"`
	SchoolCity    string    `db:"school_city" json:"school_city"`
	SchoolRegion  string    `db:"school_region" json:"school_region"`
	SchoolZipcode string    `db:"school_zipcode" json:"school_zipcode,omitempty"`
	UserID        int       `db:"user_id" json:"-"`
}

func (s *Student) Parse(data map[string]string) error {
	time, err := time.Parse("2006-01-02", data["birthdate"])
	if err != nil {
		return err
	}
	s.Name = data["name"]
	s.Birthdate = time
	s.Address = data["address"]
	s.City = data["city"]
	s.Region = data["region"]
	s.Country = data["country"]
	s.Zipcode = data["zipcode"]
	s.Education = data["education"]
	s.SchoolName = data["school_name"]
	s.SchoolAddress = data["school_address"]
	s.SchoolCity = data["school_city"]
	s.SchoolRegion = data["school_region"]
	s.SchoolZipcode = data["school_zipcode"]

	return nil
}
