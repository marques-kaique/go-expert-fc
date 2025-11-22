package entity

// por não estar em internal, outras projeto podem utilizar esse pacote
import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(id string) (ID, error) {
	// parse the string into a UUID
	newUuid, err := uuid.Parse(id)
	return ID(newUuid), err
}
