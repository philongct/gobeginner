package embeddedstruct

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

type Stringable interface {
	ToString() string
}

func stringit(s Stringable) string {
	return s.ToString()
}

// go test -run TestEmbeddedStruct
func TestEmbeddedStruct(t *testing.T) {
	u := User{"admin", "Foo", "Bar"}

	admin := AdminUser{u, time.Now()}
	fmt.Println(stringit(&admin))

	member := MemberUser{&u, time.Now()}
	fmt.Println(stringit(&member))

	memberNil := MemberUser{nil, time.Now()}
	fmt.Println(stringit(&memberNil))

	fmt.Print("\n------ Change first name, only memberUser changed ------\n\n")

	u.UserName = "member"
	fmt.Println(stringit(&admin))
	fmt.Println(stringit(&member))

	fmt.Print("\n------ Sizes ------\n\n")

	fmt.Printf("admin.GetSize(): %d, member.GetSize(): %d, memberNil.GetSize(): %d\n",
		admin.GetSize(), member.GetSize(), memberNil.GetSize())
	fmt.Printf("Size of admin: %d, Size of member: %d, Size of memberNil: %d\n",
		unsafe.Sizeof(admin), unsafe.Sizeof(member), unsafe.Sizeof(memberNil))

	fmt.Print("\n----- Auto handling of ToString() ------\n\n")

	regularUser := RegularUser{nil}
	fmt.Printf("%v Size: %d\n", stringit(&regularUser), unsafe.Sizeof(regularUser))
}
