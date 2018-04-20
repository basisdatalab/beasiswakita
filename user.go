package beasiswakita

import (
	"errors"
	"time"
)

// User hold record of all registered user
type User struct {
	ID              int               `db:"id" json:"id"`
	EmailAddress    string            `db:"email_address" json:"email_address"`
	Password        string            `db:"password" json:"password,omitempty"`
	PasswordConfirm string            `db:"-" json:"password_confirm,omitempty"`
	Role            string            `db:"role" json:"role"`
	UserData        map[string]string `db:"-" json:"user_data"`
	CreatedAt       time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time         `db:"updated_at" json:"updated_at"`
}

func (u *User) Validate() error {
	if u.Role != "organization" && u.Role != "student" {
		return errors.New("Invalid role")
	}
	return nil
}

func (u *User) ValidatePassword() error {
	if len(u.Password) < 6 || len(u.Password) > 16 {
		return errors.New("Password must between 6-16 characters")
	}

	if u.Password != u.PasswordConfirm {
		return errors.New("Password not match")
	}

	return nil
}
