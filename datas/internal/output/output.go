package output

import "time"

// Output is a data (struct) from usecase to presenter

type Person struct {
	Id       int64
	Name     string
	Birthday time.Time

}
