package main

import "fmt"

type errDatabase int

func (e errDatabase) Error() string {
    return fmt.Sprintln("Database error")
}

const (
    ErrDatabase errDatabase = 0
)
