package param

// Param is a data (struct) from request, we use it in controller

type Person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}
