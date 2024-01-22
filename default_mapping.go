// Copyright 2023, 2024 Justin Simmons.
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

// NullString converts the native go string type to a sql.NullString type.
func NullString[T string | *string](s T) sql.NullString {
	nString := sql.NullString{}

	switch v := any(s).(type) {
	case string:
		nString.String = v
		nString.Valid = true
	case *string:
		if v != nil {
			nString.String = *v
			nString.Valid = true
		}
	}

	return nString
}

// NullInt64 converts the native go int64 type to a sql.NullInt64 type.
func NullInt64[T int64 | *int64](i T) sql.NullInt64 {
	nInt := sql.NullInt64{}

	switch v := any(i).(type) {
	case int64:
		nInt.Int64 = v
		nInt.Valid = true
	case *int64:
		if v != nil {
			nInt.Int64 = *v
			nInt.Valid = true
		}
	}

	return nInt
}

// NullInt32 converts the native go int32 type to a sql.NullInt32 type.
func NullInt32[T int32 | *int32](i T) sql.NullInt32 {
	nInt := sql.NullInt32{}

	switch v := any(i).(type) {
	case int32:
		nInt.Int32 = v
		nInt.Valid = true
	case *int32:
		if v != nil {
			nInt.Int32 = *v
			nInt.Valid = true
		}
	}

	return nInt
}

// NullInt16 converts the native go int16 type to a sql.NullInt16 type.
func NullInt16[T int16 | *int16](i T) sql.NullInt16 {
	nInt := sql.NullInt16{}

	switch v := any(i).(type) {
	case int16:
		nInt.Int16 = v
		nInt.Valid = true
	case *int16:
		if v != nil {
			nInt.Int16 = *v
			nInt.Valid = true
		}
	}

	return nInt
}

// ToNullByte converts the native go byte type to a sql.NullByte type.
func NullByte[T byte | *byte](b T) sql.NullByte {
	nByte := sql.NullByte{}

	switch v := any(b).(type) {
	case byte:
		nByte.Byte = v
		nByte.Valid = true
	case *byte:
		if v != nil {
			nByte.Byte = *v
			nByte.Valid = true
		}
	}

	return nByte
}

// NullFloat64 converts the native go float64 type to a sql.NullFloat64 type.
func NullFloat64[T float64 | *float64](f T) sql.NullFloat64 {
	nFloat := sql.NullFloat64{}

	switch v := any(f).(type) {
	case float64:
		nFloat.Float64 = v
		nFloat.Valid = true
	case *float64:
		if v != nil {
			nFloat.Float64 = *v
			nFloat.Valid = true
		}
	}

	return nFloat
}

// NullBoolean converts the native go boolean type to a sql.NullBool type.
func NullBoolean[T bool | *bool](b T) sql.NullBool {
	nBool := sql.NullBool{}

	switch v := any(b).(type) {
	case bool:
		nBool.Bool = v
		nBool.Valid = true
	case *bool:
		if v != nil {
			nBool.Bool = *v
			nBool.Valid = true
		}
	}

	return nBool
}

// NullTime converts an time pointer to a sql NullTime type.
// This function will treats the Zero time as valid.
func NullTime[T time.Time | *time.Time](t T) sql.NullTime {
	nTime := sql.NullTime{}

	switch v := any(t).(type) {
	case time.Time:
		nTime.Time = v
		nTime.Valid = true
	case *time.Time:
		if v != nil {
			nTime.Time = *v
			nTime.Valid = true
		}
	}

	return nTime
}

// Serial is a notational convenience for creating unique identifier columns.
// It is an auto-incrementing integer starting from zero.
// https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-SERIAL
func Serial(in string) (uint64, error) {
	out, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		err = fmt.Errorf("sqlmap: failed to parse '%s' as unsigned int", in)
	}

	return out, err
}
