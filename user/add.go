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

func AddOrganization(userID int, data map[string]string) error {
	return nil
}

func AddStudent(userID int, data map[string]string) error {
	var student beasiswakita.Student
	err := student.Parse(data)
	if err != nil {
		return err
	}

	student.UserID = userID
	err = beasiswakita.Transaction.Insert(&student)
	if err != nil {
		return err
	}

	return nil
}
