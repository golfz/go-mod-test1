package go_mod_test1

import "github.com/golfz/go-mod-test1/implements/usecases"

func UseCase(name string) *usecases.UseCaseStruct {
	return &usecases.UseCaseStruct{Name: name}
}

