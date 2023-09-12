// Copyright 2023 Justin Simmons
//
// This file is part of sqlmap.
// sqlmap is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// sqlmap is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
// You should have received a copy of the GNU Lesser General Public License along with sqlmap. If not, see <https://www.gnu.org/licenses/>.

package sqlmap

import (
	"database/sql"
	"time"

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
