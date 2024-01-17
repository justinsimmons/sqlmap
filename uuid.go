// Copyright 2023 Justin Simmons.
//
// This file is part of sqlmap.
// sqlmap is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// sqlmap is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
// You should have received a copy of the GNU Lesser General Public License along with sqlmap. If not, see <https://www.gnu.org/licenses/>.

package sqlmap

import (
	"github.com/google/uuid"
)

// NullUUID converts a UUID pointer to a uuid.NullUUID type.
func NullUUID(id *uuid.UUID) uuid.NullUUID {
	nUUID := uuid.NullUUID{}

	if id != nil {
		nUUID.UUID = *id
		nUUID.Valid = true
	}

	return nUUID
}

// UnwrapUUID unwraps the sql null UUID to a uuid.UUID struct.
// If the value is null the function will return an empty uuid.UUID struct.
// Use UnwrapUUIDPointer() if you want a pointer value.
func UnwrapUUID(id uuid.NullUUID) uuid.UUID {
	if !id.Valid {
		return uuid.Nil
	}

	return id.UUID
}

// UnwrapUUIDPtr unwraps the sql null UUID to a uuid.UUID pointer.
func UnwrapUUIDPtr(id uuid.NullUUID) *uuid.UUID {
	if !id.Valid {
		return nil
	}

	return &id.UUID
}
