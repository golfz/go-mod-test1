package input

// Input is a data (struct) from controller to interactor <|-- usecase

type Person struct {
	Id       int64
	Name     string
	Birthday string
}
