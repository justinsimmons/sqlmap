// Copyright 2023 Justin Simmons.
//
// This file is part of sqlmap.
// sqlmap is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// sqlmap is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
// You should have received a copy of the GNU Lesser General Public License along with sqlmap. If not, see <https://www.gnu.org/licenses/>.

package sqlmap

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

// ToNullString converts a string pointer to a sql.NullString type.
func ToNullString(in *string) sql.NullString {
	nString := sql.NullString{}

	if in != nil {
		nString.String = *in
		nString.Valid = true
	}

	return nString
}

// ToNullInt64 converts an int64 pointer to a sql.NullInt64 type.
func ToNullInt64(in *int64) sql.NullInt64 {
	nInt := sql.NullInt64{}

	if in != nil {
		nInt.Int64 = *in
		nInt.Valid = true
	}

	return nInt
}

// ToNullInt32 converts an int32 pointer to a sql.NullInt32 type.
func ToNullInt32(in *int32) sql.NullInt32 {
	nInt := sql.NullInt32{}

	if in != nil {
		nInt.Int32 = *in
		nInt.Valid = true
	}

	return nInt
}

// ToNullInt16 converts an int16 pointer to a sql.NullInt16 type.
// TODO: find a clever way to do this with generics
func ToNullInt16(in *int16) sql.NullInt16 {
	nInt := sql.NullInt16{}

	if in != nil {
		nInt.Int16 = *in
		nInt.Valid = true
	}

	return nInt
}

// ToNullByte converts a byte pointer to a sql.NullByte type.
func ToNullByte(in *byte) sql.NullByte {
	nByte := sql.NullByte{}

	if in != nil {
		nByte.Byte = *in
		nByte.Valid = true
	}

	return nByte
}

// ToNullFloat64 converts a float64 pointer to a sql.NullFloat64 type.
func ToNullFloat64(in *float64) sql.NullFloat64 {
	nFloat := sql.NullFloat64{}

	if in != nil {
		nFloat.Float64 = *in
		nFloat.Valid = true
	}

	return nFloat
}

// ToNullBoolean converts an boolean pointer to a sql.NullBool type.
func ToNullBoolean(b *bool) sql.NullBool {
	nBool := sql.NullBool{}

	if b != nil {
		nBool.Bool = *b
		nBool.Valid = true
	}

	return nBool
}

// ToNullTime converts an time pointer to a sql NullTime type.
// This function will mark zero time as valid.
func ToNullTime(t *time.Time) sql.NullTime {
	nTime := sql.NullTime{}

	if t != nil {
		nTime.Time = *t
		nTime.Valid = true
	}

	return nTime
}

// Serial is a notational convenience for creating unique identifier columns.
// It is an auto-incrementing integer starting from zero.
// https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-SERIAL
func ToSerial(in string) (uint64, error) {
	out, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		err = fmt.Errorf("sqlmap: failed to parse '%s' as unsigned int", in)
	}

	return out, err
}
