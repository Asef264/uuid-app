package main

import (
	"fmt"
	"github.com/google/uuid"
)

func uuidcode() string {
	val := uuid.New().String()
	fmt.Println("a new uuid code is :", val)
	return val

}

func main() {
	
	uuidcode()
	
}