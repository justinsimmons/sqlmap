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
)

// UnwrapString unwraps the sql null string to a string pointer.
func UnwrapString(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}

	return &s.String
}

// UnwrapTime unwraps the sql null time to a time.Time struct.
// If the value is null the function will return an empty time.Time struct.
func UnwrapTime(t sql.NullTime) time.Time {
	if !t.Valid {
		return time.Time{}
	}

	return t.Time
}

// UnwrapTimePtr unwraps the sql null time to a time.Time pointer.
func UnwrapTimePtr(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}

	return &t.Time
}

// UnwrapInt64 unwraps the sql.NullInt64 to a int64 pointer.
func UnwrapInt64(i sql.NullInt64) *int64 {
	if !i.Valid {
		return nil
	}

	return &i.Int64
}

// UnwrapInt32 unwraps the sql.NullInt32 to a int32 pointer.
func UnwrapInt32(i sql.NullInt32) *int32 {
	if !i.Valid {
		return nil
	}

	return &i.Int32
}

// UnwrapInt16 unwraps the sql.NullInt16 to a int16 pointer.
func UnwrapInt16(i sql.NullInt16) *int16 {
	if !i.Valid {
		return nil
	}

	return &i.Int16
}

// UnwrapByte unwraps the sql.NullByte to a byte pointer.
func UnwrapByte(b sql.NullByte) *byte {
	if !b.Valid {
		return nil
	}

	return &b.Byte
}

// UnwrapFloat64 unwraps the sql.NullFloat64 to a float64 pointer.
func UnwrapFloat64(f sql.NullFloat64) *float64 {
	if !f.Valid {
		return nil
	}

	return &f.Float64
}

// UnwrapBoolean unwraps the sql.NullBool to a boolean pointer.
func UnwrapBoolean(b sql.NullBool) *bool {
	if !b.Valid {
		return nil
	}

	return &b.Bool
}
