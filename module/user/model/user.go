package usermodel

import (
	"Blog-CMS/common"
	"errors"
	"regexp"
	"strings"
	"time"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	Salt            string `json:"-" gorm:"column:salt"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	Phone           string `json:"phone" gorm:"column:phone"`
	Role            string `json:"role" gorm:"column:role"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) GetPassword() string {
	return u.Password
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	Salt            string `json:"-" gorm:"column:salt"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	Phone           string `json:"phone" gorm:"column:phone"`
	Role            string `json:"role" gorm:"column:role"`
}

func (u *UserCreate) Validate() error {

	if err := u.validateEmail(); err != nil {
		return err
	}

	if err := u.ValidatePassword(); err != nil {
		return err
	}

	return nil
}

func (u *UserCreate) validateEmail() error {
	if !strings.HasSuffix(u.Email, "@gmail.com") {
		return ErrNotAnEmail
	}
	return nil
}

func (u *UserCreate) ValidatePassword() error {
	// Check length requirement (8-50 characters)
	if len(u.Password) < 8 || len(u.Password) > 50 {
		return ErrPasswordTooShort
	}

	// Check for at least one uppercase letter (A-Z)
	if matched, _ := regexp.MatchString(`[A-Z]`, u.Password); !matched {
		return ErrPasswordMissingUppercase
	}

	// Check for at least one lowercase letter (a-z)
	if matched, _ := regexp.MatchString(`[a-z]`, u.Password); !matched {
		return ErrPasswordMissingLowercase
	}

	// Check for at least one numeric digit (0-9)
	if matched, _ := regexp.MatchString(`[0-9]`, u.Password); !matched {
		return ErrPasswordMissingNumber
	}

	// Check for at least one special character (!@#$%^&* and others)
	if matched, _ := regexp.MatchString(`[!@#$%^&*()_+\-=\[\]{};':",.<>?/\\|]`, u.Password); !matched {
		return ErrPasswordMissingSpecialChar
	}

	// Ensure the password does not contain any whitespace characters
	if matched, _ := regexp.MatchString(`\s`, u.Password); matched {
		return ErrPasswordContainsWhitespace
	}

	// If all checks pass, return nil (valid password)
	return nil
}

func (UserCreate) TableName() string { return User{}.TableName() }

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserRegister struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}

var (
	loginAttempts = make(map[string]int)
	blockTime     = make(map[string]time.Time)
)

const (
	maxAttempts   = 3
	blockDuration = 5 * time.Minute
)

func (u *UserLogin) ValidateBlock() error {
	if blockUntil, found := blockTime[u.Email]; found {
		if time.Now().Before(blockUntil) {
			return ErrTooManyLoginAttempts
		}

		delete(loginAttempts, u.Email)
		delete(blockTime, u.Email)
	}

	return nil
}

func (u *UserLogin) RegisterFailedAttempt() {
	loginAttempts[u.Email]++
	if loginAttempts[u.Email] >= maxAttempts {
		blockTime[u.Email] = time.Now().Add(blockDuration)
	}
}

func (u *UserLogin) ResetAttempts() {
	delete(loginAttempts, u.Email)
	delete(blockTime, u.Email)
}

type UserChangePd struct {
	Password string `json:"password" gorm:"password"`
}

var (
	ErrEmailnameOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailnameOrPasswordInvalid")

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted")

	ErrNotAnEmail = common.NewCustomError(
		errors.New("Please using an email"),
		"Please using an email",
		"ErrNotAnEmail")

	ErrPasswordTooShort = common.NewCustomError(
		errors.New("password must be between 8 and 50 characters"),
		"password must be between 8 and 50 characters",
		"ErrPasswordTooShort")

	ErrPasswordMissingUppercase = common.NewCustomError(
		errors.New("password must contain at least 1 uppercase letter"),
		"password must contain at least 1 uppercase letter",
		"ErrPasswordMissingUppercase")

	ErrPasswordMissingLowercase = common.NewCustomError(
		errors.New("password must contain at least 1 lowercase letter"),
		"password must contain at least 1 lowercase letter",
		"ErrPasswordMissingLowercase")

	ErrPasswordMissingNumber = common.NewCustomError(
		errors.New("password must contain at least 1 number"),
		"password must contain at least 1 number",
		"ErrPasswordMissingNumber")

	ErrPasswordMissingSpecialChar = common.NewCustomError(
		errors.New("password must contain at least 1 special character (!@#$%^&*)"),
		"password must contain at least 1 special character (!@#$%^&*)",
		"ErrPasswordMissingSpecialChar")

	ErrPasswordContainsWhitespace = common.NewCustomError(
		errors.New("password cannot contain whitespace"),
		"password cannot contain whitespace",
		"ErrPasswordContainsWhitespace")

	ErrTooManyLoginAttempts = common.NewCustomError(
		errors.New("too many failed login attempts, please try again later"),
		"too many failed login attempts, please try again later",
		"ErrTooManyLoginAttempts",
	)
)
