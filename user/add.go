package user

import (
	"time"

	"github.com/basisdatalab/beasiswakita"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(user beasiswakita.User) (beasiswakita.User, int, error) {
	err := user.ValidatePassword()
	if err != nil {
		return user, 0, err
	}

	err = user.Validate()
	if err != nil {
		return user, 0, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, 0, err
	}

	user.Password = string(hashed)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = beasiswakita.Transaction.Insert(&user)
	if err != nil {
		return user, 0, err
	}

	user.Password = ""
	user.PasswordConfirm = ""
	var profileID int

	if user.Role == "organization" {
		profileID, err = AddOrganization(user.ID, user.UserData)
	} else {
		profileID, err = AddStudent(user.ID, user.UserData)
	}
	if err != nil {
		return user, 0, err
	}

	return user, profileID, nil
}

func AddOrganization(userID int, data map[string]interface{}) (int, error) {
	var organization beasiswakita.Organization
	err := organization.Parse(data)
	if err != nil {
		return 0, err
	}

	organization.CreatedAt = time.Now()
	organization.UpdatedAt = time.Now()

	organization.UserID = userID
	err = beasiswakita.Transaction.Insert(&organization)
	if err != nil {
		return 0, err
	}

	return organization.ID, nil
}

func AddStudent(userID int, data map[string]interface{}) (int, error) {
	var student beasiswakita.Student
	err := student.Parse(data)
	if err != nil {
		return 0, err
	}

	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	student.UserID = userID
	err = beasiswakita.Transaction.Insert(&student)
	if err != nil {
		return 0, err
	}

	return student.ID, nil
}
