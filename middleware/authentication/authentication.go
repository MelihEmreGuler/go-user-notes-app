package authentication

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/MelihEmreGuler/go-user-notes-app/models"
	"github.com/MelihEmreGuler/go-user-notes-app/repository"
	"strings"
)

// SignUp signs up the user
func SignUp(username, email, password string) error {

	//checks if the username is valid
	if valid, err := usernameValid(username); !valid {
		return err
	}
	//checks if the email is valid
	if valid, err := emailValid(email); !valid {
		return err
	}

	//hashes the password
	hashedPassword := hashPassword(password)

	//inserts the user to the database
	return repository.R.InsertUser(username, email, hashedPassword)
}

// SignIn signs in the user with username or email
func SignIn(usernameOrEmail, password string) (*models.User, error) {

	var user *models.User
	var err error
	isUsername := true

	//checks if the usernameOrEmail is email
	if strings.Contains(usernameOrEmail, "@") {
		isUsername = false
	}

	//brings the user from the database
	if isUsername {
		user, err = repository.R.SelectUserByUsername(usernameOrEmail)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = repository.R.SelectUserByEmail(usernameOrEmail)
		if err != nil {
			return nil, err
		}
	}

	//checks if the user exists
	if user == nil {
		return user, errors.New("user does not exist")
	}

	//checks if the password is correct
	if checkPassword(password, user.PasswordHash) {
		fmt.Println("password is correct for user: " + usernameOrEmail)
		return user, nil // password is correct
	} else {
		return user, errors.New(usernameOrEmail + " password is incorrect")
	}
}

// hashPassword hashes the password using sha256
func hashPassword(password string) string {
	hashed := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hashed[:])
}

// checkPassword checks if the input password is correct
func checkPassword(inputPassword, storedHashedPassword string) bool {
	//brings the hashed version of the input password
	hashedInputPassword := hashPassword(inputPassword)

	//compares the hashed version of the input password with the stored hashed password
	if hashedInputPassword == storedHashedPassword {
		return true // password is correct
	}
	return false // password is incorrect
}

func usernameValid(username string) (bool, error) {

	//checks if the username is empty
	if username == "" {
		return false, errors.New("username cannot be empty")
	}

	//checks if the username is longer than 3 characters
	if len(username) < 3 {
		return false, errors.New("username must be longer than 3 characters")
	}

	//checks if the username contains only letters and numbers
	for _, ch := range username {
		if !(ch >= 'a' && ch <= 'z') && !(ch >= 'A' && ch <= 'Z') && !(ch >= '0' && ch <= '9') {
			return false, errors.New("username contains invalid characters")
		}
	}

	return true, nil
}

func emailValid(email string) (bool, error) {
	//checks if the email is empty
	if email == "" {
		return false, errors.New("email cannot be empty")
	}

	//checks if the email is longer than 3 characters
	if len(email) < 3 {
		return false, errors.New("email must be longer than 3 characters")
	}

	//checks if the email contains "@" character and only one and not at the beginning or end.
	atCount := 0
	for i, ch := range email {
		if ch == '@' {
			atCount++
			if i == 0 || i == len(email)-1 {
				return false, errors.New("email cannot start or end with @")
			}
		}
	}

	//checks last 4 characters of the email for ".com"
	if email[len(email)-4:] != ".com" {
		return false, errors.New("email must end with .com")
	}

	return true, nil
}
