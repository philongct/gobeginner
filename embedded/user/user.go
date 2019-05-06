package user

import (
	"strings"
	"unsafe"
)

type User struct {
	UserName string

	FirstName string
	LastName  string
}

func (user *User) GetSize() uintptr {
	// Returns 48 because User has 3 fields, each has 16 bytes
	return unsafe.Sizeof(*user)
}

func (user *User) GetFullName() string {
	return user.FirstName + " " + user.LastName
}

func (user *User) ToString() string {
	if user == nil {
		return "<nil>"
	}

	return strings.Join([]string{
		"UserName=" + user.UserName,
		"fullName=" + user.GetFullName(),
	}, ", ")
}
