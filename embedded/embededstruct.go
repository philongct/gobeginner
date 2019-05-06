package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/philongct/golessions/embedded/user"
)

type Stringable interface {
    ToString() string
}

func stringit(s Stringable) string {
    return s.ToString()
}

func main() {
	admin := user.AdminUser{user.User{"admin", "Foo", "Bar"}, time.Now()}
	fmt.Println(stringit(&admin))

	member := user.MemberUser{&user.User{"member", "Fizz", "Buzz"}, time.Now()}
	fmt.Println(stringit(&member))

	memberNil := user.MemberUser{nil, time.Now()}
	fmt.Println(stringit(&memberNil))

	fmt.Println()

	fmt.Printf("admin.GetSize(): %d, member.GetSize(): %d, memberNil.GetSize(): %d\n",
		admin.GetSize(), member.GetSize(), memberNil.GetSize())
	fmt.Printf("Size of admin: %d, Size of member: %d, Size of memberNil: %d\n",
		unsafe.Sizeof(admin), unsafe.Sizeof(member), unsafe.Sizeof(memberNil))

	fmt.Println()

	regularUser := user.RegularUser{nil}
	fmt.Printf("%v Size: %d", stringit(&regularUser), unsafe.Sizeof(regularUser))
}
