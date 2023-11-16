// Copyright 2023 Justin Simmons.
//
// This file is part of sqlmap.
// sqlmap is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// sqlmap is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
// You should have received a copy of the GNU Lesser General Public License along with sqlmap. If not, see <https://www.gnu.org/licenses/>.

package sqlmap

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UnwrapString unwraps the sql null string to a string pointer.
func UnwrapString(s sql.NullString) *string {
	var str *string

	if s.Valid {
		str = &s.String
	}

	return str
}

// UnwrapTime unwraps the sql null time to a time.Time struct.
// If the value is null the function will return an empty time.Time struct.
func UnwrapTime(t sql.NullTime) time.Time {
	var tm time.Time

	if t.Valid {
		tm = t.Time
	}

	return tm
}

// UnwrapTime unwraps the sql null time to a timestamppb.Timestamp pointer.
func UnwrapTimestamp(t sql.NullTime) *timestamppb.Timestamp {
	tm := UnwrapTime(t)

	if tm.IsZero() {
		return nil
	}

	return timestamppb.New(tm)
}

// UnwrapInt64 unwraps the sql.NullInt64 to a int64 pointer.
func UnwrapInt64(in sql.NullInt64) *int64 {
	var nInt *int64

	if in.Valid {
		nInt = &in.Int64
	}

	return nInt
}

// UnwrapInt32 unwraps the sql.NullInt32 to a int32 pointer.
func UnwrapInt32(in sql.NullInt32) *int32 {
	var nInt *int32

	if in.Valid {
		nInt = &in.Int32
	}

	return nInt
}

// UnwrapInt16 unwraps the sql.NullInt16 to a int16 pointer.
func UnwrapInt16(in sql.NullInt16) *int16 {
	var nInt *int16

	if in.Valid {
		nInt = &in.Int16
	}

	return nInt
}

// UnwrapByte unwraps the sql.NullByte to a byte pointer.
func UnwrapByte(in sql.NullByte) *byte {
	var nByte *byte

	if in.Valid {
		nByte = &in.Byte
	}

	return nByte
}

// UnwrapFloat64 unwraps the sql.NullFloat64 to a float64 pointer.
func UnwrapFloat64(in sql.NullFloat64) *float64 {
	var nFloat *float64

	if in.Valid {
		nFloat = &in.Float64
	}

	return nFloat
}

// UnwrapBoolean unwraps the sql.NullBool to a boolean pointer.
func UnwrapBoolean(in sql.NullBool) *bool {
	var nBool *bool

	if in.Valid {
		nBool = &in.Bool
	}

	return nBool
}

// UnwrapUUID unwraps the github.com/google/uuid.NullUUID to a uuid.UUID pointer.
func UnwrapUUID(in uuid.NullUUID) *uuid.UUID {
	var nID *uuid.UUID

	if in.Valid {
		nID = in.UUID
	}

	return nID
}
