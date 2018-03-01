package server

import (
	"fmt"

	"github.com/JREAMLU/core/guid"
)

// UUID uuid
func UUID() {
	fmt.Println("++++++++++++: ", guid.NewObjectID().Hex())
}
