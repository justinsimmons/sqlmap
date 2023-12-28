package sqlmap

import (
	"github.com/google/uuid"
)

// ToNullUUID converts a UUID pointer to a uuid.NullUUID type.
func ToNullUUID(id *uuid.UUID) uuid.NullUUID {
	nUUID := uuid.NullUUID{}

	if id != nil {
		nUUID.UUID = *id
		nUUID.Valid = true
	}

	return nUUID
}
