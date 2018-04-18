package authentication_test

import (
	"os"
	"testing"

	"github.com/harkce/beasiswakita/authentication"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
)

func TestMain(m *testing.M) {
	gotenv.Load(os.Getenv("GOPATH") + "/src/github.com/harkce/beasiswakita/.env")
}

func TestAuthenticate(t *testing.T) {
	password := "s0m3r4nd0mp455w0rd"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword := string(hash)

	err := authentication.Authenticate(password, hashedPassword)
	if err != nil {
		t.Error("Should have no error")
	}

	wrongPassword := "wr0ngp455w0rd"

	err = authentication.Authenticate(wrongPassword, hashedPassword)
	if err == nil {
		t.Error("Should have error")
	}
}

func TestAuthorize(t *testing.T) {
	owner := authentication.Owner{
		ID:           1,
		EmailAddress: "habibgarut@gmail.com",
		Role:         "student",
	}
	issuer := "habibfikri"

	_, err := authentication.Authorize(owner, issuer)
	if err != nil {
		t.Error("Should have no error")
	}
}

func TestToken(t *testing.T) {
	owner := authentication.Owner{
		ID:           1,
		EmailAddress: "habibgarut@gmail.com",
		Role:         "student",
	}
	issuer := "habibfikri"

	token, err := authentication.Authorize(owner, issuer)
	if err != nil {
		t.Error("Should have no error")
	}

	bearerToken := "Bearer " + token

	_, err = authentication.Token(bearerToken, []string{})
	if err != nil {
		t.Error("Should not have error")
	}

	wrongToken := "Token " + token

	_, err = authentication.Token(wrongToken, []string{})
	if err == nil {
		t.Error("Should have error")
	}

	_, err = authentication.Token(bearerToken, []string{"organization"})
	if err == nil {
		t.Error("Should have error")
	}

	os.Setenv("APP_KEY", "wr0ngs1gn1ngk3y")
	_, err = authentication.Token(bearerToken, []string{})
	if err == nil {
		t.Error("Should have error")
	}
}
