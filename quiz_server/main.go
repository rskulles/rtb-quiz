package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	bytes, _ := id.MarshalBinary()
	fmt.Printf("%v\n%v\n", id, bytes)
}
