package main

import (
	"encoding/json"
	"fmt"
)

type CreditCard string

type ConfidentialCustomer struct {
	CustomerID int
	CreditCard CreditCard
}

func (c CreditCard) String() string {
	return fmt.Sprintln("xxxx-xxxx-xxxx")
}

func (c CreditCard) GoString() string {
	return fmt.Sprintln("xxxx-xxxx-xxxx")
}

func GoStringerAndStringerInterface() {
	c := ConfidentialCustomer{
		CustomerID: 1,
		CreditCard: "4111-1111-2222",
	}

	//fmtで出力する時はマスクされる
	fmt.Println(c)
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%#v\n", c)

	//JSONとしてプログラム内では元のデータを扱える
	bytes, _ := json.Marshal(c)
	fmt.Println("JSON: ", string(bytes))
}
