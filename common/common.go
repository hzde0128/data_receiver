package common

import (
	"time"

	"github.com/satori/go.uuid"
)

func GetUUID() (string, error) {
	u := uuid.NewV4()

	return u.String(), nil
}
