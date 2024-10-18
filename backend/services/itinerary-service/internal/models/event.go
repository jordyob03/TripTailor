package models

type Event struct {
	Id          int
	Name        string
	Date        string
	Time        string
	Location    string
	Price       float64
	Description string
	Photos      []string
}
