package entity

type Person struct {
	Id       int64
	Name     string
	Birthday string
}

func (e Person) GetAge() int64 {
	// example
	return 36
}