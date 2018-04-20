package beasiswakita

type Organization struct {
	ID                int    `db:"id" json:"id"`
	Name              string `db:"name" json:"name"`
	Positions         string `db:"positions" json:"positions"`
	OrganizationName  string `db:"organization_name" json:"organization_name"`
	OrganizationEmail string `db:"organization_email" json:"organization_email"`
	CompanyAddress    string `db:"company_address" json:"company_address"`
	City              string `db:"city" json:"city"`
	Region            string `db:"region" json:"region"`
	Country           string `db:"country" json:"country"`
	Zipcode           string `db:"zipcode" json:"zipcode,omitempty"`
	Website           string `db:"website" json:"website,omitempty"`
}
