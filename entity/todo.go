package entity

import "time"

type Todo struct {
	Todo_Id    uint
	Title      string
	Completed  bool
	Created_At time.Time
	Updated_At time.Time
}
