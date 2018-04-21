package user

import (
	"time"

	"github.com/harkce/beasiswakita"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(user beasiswakita.User) (beasiswakita.User, error) {
	err := user.ValidatePassword()
	if err != nil {
		return user, err
	}

	err = user.Validate()
	if err != nil {
		return user, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hashed)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = beasiswakita.Transaction.Insert(&user)
	if err != nil {
		return user, err
	}

	user.Password = ""
	user.PasswordConfirm = ""

	if user.Role == "organization" {
		err = AddOrganization(user.ID, user.UserData)
	} else {
		err = AddStudent(user.ID, user.UserData)
	}
	if err != nil {
		return user, err
	}

	return user, nil
}

func AddOrganization(userID int, data map[string]interface{}) error {
	var organization beasiswakita.Organization
	err := organization.Parse(data)
	if err != nil {
		return err
	}

	organization.CreatedAt = time.Now()
	organization.UpdatedAt = time.Now()

	organization.UserID = userID
	err = beasiswakita.Transaction.Insert(&organization)
	if err != nil {
		return err
	}

	return nil
}

func AddStudent(userID int, data map[string]interface{}) error {
	var student beasiswakita.Student
	err := student.Parse(data)
	if err != nil {
		return err
	}

	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	student.UserID = userID
	err = beasiswakita.Transaction.Insert(&student)
	if err != nil {
		return err
	}

	return nil
}
