package user

import (
	"time"
)

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
