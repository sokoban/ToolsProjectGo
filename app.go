package main

import (
	"fmt"
	"strconv"
)

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
