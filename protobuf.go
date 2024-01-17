// Copyright 2023 Justin Simmons.
//
// This file is part of sqlmap.
// sqlmap is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// sqlmap is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
// You should have received a copy of the GNU Lesser General Public License along with sqlmap. If not, see <https://www.gnu.org/licenses/>.

package sqlmap

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// NullTimeFromTimestamp converts an timestamppb.Timestamp pointer to a sql.NullTime type.
// This will save you a conversion when dealing with gRPC requests.
func NullTimeFromTimestamp(t *timestamppb.Timestamp) sql.NullTime {
	nTime := sql.NullTime{}

	if t != nil && t.IsValid() {
		nTime.Time = t.AsTime()
		nTime.Valid = true
	}

	return nTime
}

// UnwrapTime unwraps the sql null time to a timestamppb.Timestamp pointer.
func UnwrapTimestamp(t sql.NullTime) *timestamppb.Timestamp {
	tm := UnwrapTime(t)

	if tm.IsZero() {
		return nil
	}

	return timestamppb.New(tm)
}
