package input

// ViewModel is a data (struct) for view

type Person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}
