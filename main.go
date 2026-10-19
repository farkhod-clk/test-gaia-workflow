package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println(uuid.NewString(), errors.New("ok"))
}
