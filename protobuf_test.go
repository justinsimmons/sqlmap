// Copyright 2023, 2024 Justin Simmons.
//
// This file is part of sqlmap.
// sqlmap is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// sqlmap is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
// You should have received a copy of the GNU Lesser General Public License along with sqlmap. If not, see <https://www.gnu.org/licenses/>.

package sqlmap

import (
	"database/sql"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNullTimeFromTimestamp(t *testing.T) {
	testCases := []struct {
		name  string
		input *timestamppb.Timestamp
	}{
		{
			name:  "successfully handle non-pointer timestamppb.Timestamp values ",
			input: timestamppb.Now(),
		},
		{
			name:  "successfully handle null timestamppb.Timestamp pointers",
			input: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NullTimeFromTimestamp(tc.input)

			if tc.input == nil {
				if result.Valid {
					t.Fatalf("result should be invalid, got: '%v'", result)
				}

				return
			}

			if !result.Valid {
				t.Fatalf("result should be valid value, got: '%v'", result)
			}

			if !result.Time.Equal(tc.input.AsTime()) {
				t.Fatalf("result mismatch got '%v', expected: '%v'", result.Time, tc.input.AsTime())
			}
		})
	}
}

func TestUnwrapTimestamp(t *testing.T) {
	testCases := []struct {
		name  string
		input sql.NullTime
	}{
		{
			name: "should successfully unwrap sql.NullTime with valid value",
			input: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		},
		{
			name: "should successfully unwrap sql.NullTime with null value",
			input: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
		},
		{
			name: "should successfully unwrap invalid sql.NullTime with value",
			input: sql.NullTime{
				Time:  time.Now(),
				Valid: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UnwrapTimestamp(tc.input)

			if !tc.input.Valid {
				if err := result.CheckValid(); err == nil {
					t.Fatalf("timestamp should not be valid, got: '%v'", result)
				}

				return
			}

			if err := result.CheckValid(); err != nil {
				t.Fatalf("result should be a valid timestamp, expected: '%s' got: '%v'", tc.input.Time, err)
			}

			if !result.AsTime().Equal(tc.input.Time) {
				t.Fatalf("result mismatch got '%v', expected: '%v'", result.AsTime(), tc.input.Time)
			}
		})
	}
}
