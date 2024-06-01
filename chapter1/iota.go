package main

import "fmt"

//go:generate go run github.com/dmarkham/enumer -type=CarType 
//go:generate go run github.com/dmarkham/enumer -type=CarOption

type CarType int

const (
	Sedan CarType = iota + 1
	Hatchback
	MPV
	SUV
	Crossover
	Coupe
	Convertible
)

type CarOption int

const (
    GPS CarOption = 1 << iota
    AWD
    SunRoof
    HeatedSeat
    DriverAssist
)

func HasSunRoof(o CarOption) bool {
    return o&SunRoof != 0
}

func CallSedan() {
    var t CarType
    t = Sedan
    fmt.Printf("Hey %v\n", t)
    o, err  := CarOptionString("GPS")
    if err != nil {
        fmt.Println("Failed to execute CarOptionString()")
        return
    }
    fmt.Printf("GPS is %v\n", o)
}
