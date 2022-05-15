package utils

import (
	"fmt"
	"strconv"
)

type Member struct {
	UserID  string
	UserNo  int
	Name    string
	Address string
}

type VIPMember struct {
	Member
	Level int
}

func UserSt() {
	member := Member{"guest", 1213, "김수로", "경기도"}
	var member2 []Member
	member2 = append(member2, member)

	vip := VIPMember{
		member,
		3,
	}

	var vip2 []VIPMember

	vip2 = append(vip2, vip)
	vip2 = append(vip2, vip)
	vip2 = append(vip2, vip)

	for i, a := range vip2 {
		fmt.Println(i, a)
	}

	fmt.Printf("유저: %s ID: %s 주소 %s\n", member.Name, member.UserID, member.Address)
	fmt.Printf("VIP 유저: %s ID: %s 나이 %d VIP 레벨: %d 유저 레벨:%s\n",
		vip.Name,
		vip.UserID,
		vip.UserNo,
		vip.Level,
		vip.Member.Name,
	)
}

func Attach(data string) (string, int) {

	var result string
	var data2 string = "world"

	if data == "hello" {
		result = data + " " + data2
	} else {
		result = " " + data2
	}

	if data == "hello" {

		fmt.Println("hello")

	} else if data == "bye" {

		fmt.Println("bye")

	} else {

		fmt.Println("Good bye")

	}

	return result, len(result)

}

func main() {

	var data string = "bye"
	ret, index := Attach(data)
	UserSt()

	fmt.Println(ret + ": length : " + strconv.Itoa(index))

	if rets := tf(true); rets == true {
		fmt.Println("Success")
	}

	var data3 [5]string = [5]string{"haha", "hihi", "huhu"}
	data3[3] = "hoho"

	for i, a := range data3 {
		fmt.Println(i, a)
	}

	fmt.Print(data3)

	switch rets := tf(false); rets {
	case true:
		fmt.Println("hello")
	case false:
		fmt.Println("bye")
	default:
		fmt.Println("Good Bye")
	}
	for a := 0; a < 10; a++ {
		fmt.Println(a, data)
	}

}

func tf(para bool) bool {
	return para
}
