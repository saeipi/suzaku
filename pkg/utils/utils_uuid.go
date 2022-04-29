package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func GetUUID() (id string) {
	id = fmt.Sprintf("%s", uuid.NewV4())
	return
}
