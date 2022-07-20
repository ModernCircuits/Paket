package innosetup

import (
	"strings"

	"github.com/google/uuid"
)

func NewInnoSetupGUID() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return strings.ToUpper(id.String()), nil
}
