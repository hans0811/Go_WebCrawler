package services

import (
	"context"
	"reflect"
)

// import "fmt"

func Populate(target interface{}) error {
	return PopulateForContext(cpntext.Background(), target)
}

func PopulateForContext(c context.Context, target interface{})(err error){
	return PopulateForContextWithExtras(c, target, make(map[reflect.Type]reflect.Value))
}