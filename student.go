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
	SchoolZipcode *string   `db:"school_zipcode" json:"school_zipcode,omitempty"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	UserID        int       `db:"user_id" json:"-"`
}

func (s *Student) Parse(data map[string]interface{}) error {
	time, err := time.Parse("2006-01-02", data["birthdate"].(string))
	if err != nil {
		return err
	}
	s.Name = data["name"].(string)
	s.Birthdate = time
	s.Address = data["address"].(string)
	s.City = data["city"].(string)
	s.Region = data["region"].(string)
	s.Country = data["country"].(string)
	s.Zipcode = data["zipcode"].(string)
	s.Education = data["education"].(string)
	s.SchoolName = data["school_name"].(string)
	s.SchoolAddress = data["school_address"].(string)
	s.SchoolCity = data["school_city"].(string)
	s.SchoolRegion = data["school_region"].(string)
	s.SchoolZipcode = data["school_zipcode"].(*string)

	return nil
}

func (s *Student) Map() map[string]interface{} {
	data := make(map[string]interface{})
	data["id"] = s.ID
	data["name"] = s.Name
	data["address"] = s.Address
	data["city"] = s.City
	data["region"] = s.Region
	data["country"] = s.Country
	data["zipcode"] = s.Zipcode
	data["education"] = s.Education
	data["school_name"] = s.SchoolName
	data["school_address"] = s.SchoolAddress
	data["school_city"] = s.SchoolCity
	data["school_region"] = s.SchoolRegion
	data["school_zipcode"] = s.SchoolZipcode

	return data
}
