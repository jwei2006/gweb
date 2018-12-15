package util

import (
	"github.com/satori/go.uuid"
	"strings"
)

func NewUuid() string {
	return strings.Replace(uuid.Must(uuid.NewV4()).String(), "-", "", -1)
}