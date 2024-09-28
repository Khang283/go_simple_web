package types

import "time"

type Class struct {
	StartDate   time.Time
	EndDate     time.Time
	Id          string
	Name        string
	Code        string
	Description string
	Type        string
}
