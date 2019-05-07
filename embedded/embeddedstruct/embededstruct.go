package embeddedstruct

import (
	"strings"
	"time"
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

// AdminUser 72 bytes
type AdminUser struct {
	User
	AdminSince time.Time
}

// MemberUser 32 bytes
type MemberUser struct {
	*User
	JoinAt time.Time
}

// RegularUser 32 bytes
type RegularUser struct {
	*User
}

func (adminUser *AdminUser) ToString() string {
	return adminUser.User.ToString() + ", AdminSince=" + adminUser.AdminSince.String()
	// comment above lines and uncomment following line to see differences
	//return fmt.Sprintf("%+v", adminUser)
}

func (memberUser *MemberUser) ToString() string {
	userToString := "<no info>, "
	if memberUser.User != nil {
		userToString = memberUser.User.ToString() + ", "
	}

	return userToString + "JoinSince=" + memberUser.JoinAt.String()
	// comment above lines and uncomment following line to see differences
	//return fmt.Sprintf("%+v", memberUser)
}
