package main

import (
	"fmt"
	"time"
)

type Book struct {
	Title      string
	Author     string
	Publisher  string
	ReleasedAt time.Time
	ISBN       string
}

// 普通はAuthorを別の構造体として宣言する
type Book2 struct {
	Title  string
	Author struct {
		FirstName string
		LastName  string
	}
	Publisher  string
	ReleasedAt time.Time
}

func StructSample() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	book1 := Book{
		Title:      "Real World HTTP",
		Author:     "渋川よしき",
		Publisher:  "オライリー・ジャパン",
		ISBN:       "4873119030",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book1.Title)

	book2 := Book2{
		Title: "Real World HTTP",
		Author: struct {
			FirstName string
			LastName  string
		}{
			FirstName: "よしき",
			LastName:  "渋川",
		},
		Publisher:  "オライリー・ジャパン",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book2)

	//テストとか書く時は以下のように使い捨ての構造体を書く
	book := struct {
		Title      string
		Author     string
		Publisher  string
		ISBN       string
		ReleasedAt time.Time
	}{
		Title:      "Real World HTTP",
		Author:     "渋川よしき",
		Publisher:  "オライリー・ジャパン",
		ISBN:       "4873119030",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book)
}
