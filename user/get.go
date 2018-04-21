package user

import (
	"github.com/harkce/beasiswakita"
)

func GetUser(ID int) (interface{}, error) {
	var user beasiswakita.User
	err := beasiswakita.DbMap.SelectOne(&user, "select * from users where id = ?", ID)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	if user.Role == "organization" {
		var organization beasiswakita.Organization
		err = beasiswakita.DbMap.SelectOne(&organization, "select * from organizations where user_id = ?", ID)
		if err != nil {
			return nil, err
		}
		user.UserData = organization.Map()
	} else {
		var student beasiswakita.Student
		err = beasiswakita.DbMap.SelectOne(&student, "select * from students where user_id = ?", ID)
		if err != nil {
			return nil, err
		}
		user.UserData = student.Map()
	}

	return user, nil
}
