package sqlmap

import (
	"database/sql"
	"strings"
	"testing"
	"time"
)

func TestUnwrapString(t *testing.T) {
	testCases := []struct {
		name  string
		input sql.NullString
	}{
		{
			name: "should successfully unwrap sql.NullString with valid value",
			input: sql.NullString{
				String: "foobar",
				Valid:  true,
			},
		},
		{
			name: "should successfully unwrap sql.NullString with null value",
			input: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
		{
			name: "should successfully unwrap invalid sql.NullString with value",
			input: sql.NullString{
				String: "foobar",
				Valid:  false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UnwrapString(tc.input)

			if !tc.input.Valid {
				if result != nil {
					t.Fatalf("result should be null, got: '%s'", *result)
				}

				return
			}

			if result == nil {
				t.Fatalf("result should not be null, expected: '%s'", tc.input.String)
			}

			if strings.Compare(*result, tc.input.String) != 0 {
				t.Fatalf("result mismatch got '%s', expected: '%s'", *result, tc.input.String)
			}
		})
	}
}

func TestUnwrapTime(t *testing.T) {
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
			result := UnwrapTime(tc.input)

			if !tc.input.Valid {
				if !result.IsZero() {
					t.Fatalf("result should be zero time, got: '%v'", result)
				}

				return
			}

			if result.IsZero() {
				t.Fatalf("result should not be zero time, expected: '%s'", tc.input.Time)
			}

			if !result.Equal(tc.input.Time) {
				t.Fatalf("result mismatch got '%v', expected: '%v'", result, tc.input.Time)
			}
		})
	}
}

func TestUnwrapTimePtr(t *testing.T) {
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
			result := UnwrapTimePtr(tc.input)

			if !tc.input.Valid {
				if result != nil {
					t.Fatalf("result should be nil, got: '%v'", result)
				}

				return
			}

			if result == nil {
				t.Fatalf("result should not be nil, expected: '%s'", tc.input.Time)
			}

			if !result.Equal(tc.input.Time) {
				t.Fatalf("result mismatch got '%v', expected: '%v'", result, tc.input.Time)
			}
		})
	}
}

func TestUnwrapInt64(t *testing.T) {
	testCases := []struct {
		name  string
		input sql.NullInt64
	}{
		{
			name: "should successfully unwrap sql.NullInt64 with valid value",
			input: sql.NullInt64{
				Int64: 42,
				Valid: true,
			},
		},
		{
			name: "should successfully unwrap sql.NullInt64 with null value",
			input: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
		},
		{
			name: "should successfully unwrap invalid sql.NullInt64 with value",
			input: sql.NullInt64{
				Int64: 42,
				Valid: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UnwrapInt64(tc.input)

			if !tc.input.Valid {
				if result != nil {
					t.Fatalf("result should be null, got: '%d'", *result)
				}

				return
			}

			if result == nil {
				t.Fatalf("result should not be null, expected: '%d'", tc.input.Int64)
			}

			if *result != tc.input.Int64 {
				t.Fatalf("result mismatch got '%d', expected: '%d'", *result, tc.input.Int64)
			}
		})
	}
}

func TestUnwrapInt32(t *testing.T) {
	testCases := []struct {
		name  string
		input sql.NullInt32
	}{
		{
			name: "should successfully unwrap sql.NullInt32 with valid value",
			input: sql.NullInt32{
				Int32: 42,
				Valid: true,
			},
		},
		{
			name: "should successfully unwrap sql.NullInt32 with null value",
			input: sql.NullInt32{
				Int32: 0,
				Valid: false,
			},
		},
		{
			name: "should successfully unwrap invalid sql.NullInt32 with value",
			input: sql.NullInt32{
				Int32: 42,
				Valid: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UnwrapInt32(tc.input)

			if !tc.input.Valid {
				if result != nil {
					t.Fatalf("result should be null, got: '%d'", *result)
				}

				return
			}

			if result == nil {
				t.Fatalf("result should not be null, expected: '%d'", tc.input.Int32)
			}

			if *result != tc.input.Int32 {
				t.Fatalf("result mismatch got '%d', expected: '%d'", *result, tc.input.Int32)
			}
		})
	}
}

func TestUnwrapInt16(t *testing.T) {
	testCases := []struct {
		name  string
		input sql.NullInt16
	}{
		{
			name: "should successfully unwrap sql.NullInt16 with valid value",
			input: sql.NullInt16{
				Int16: 42,
				Valid: true,
			},
		},
		{
			name: "should successfully unwrap sql.NullInt16 with null value",
			input: sql.NullInt16{
				Int16: 0,
				Valid: false,
			},
		},
		{
			name: "should successfully unwrap invalid sql.NullInt16 with value",
			input: sql.NullInt16{
				Int16: 42,
				Valid: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UnwrapInt16(tc.input)

			if !tc.input.Valid {
				if result != nil {
					t.Fatalf("result should be null, got: '%d'", *result)
				}

				return
			}

			if result == nil {
				t.Fatalf("result should not be null, expected: '%d'", tc.input.Int16)
			}

			if *result != tc.input.Int16 {
				t.Fatalf("result mismatch got '%d', expected: '%d'", *result, tc.input.Int16)
			}
		})
	}
}

func TestUnwrapByte(t *testing.T) {
	testCases := []struct {
		name  string
		input sql.NullByte
	}{
		{
			name: "should successfully unwrap sql.NullByte with valid value",
			input: sql.NullByte{
				Byte:  1,
				Valid: true,
			},
		},
		{
			name: "should successfully unwrap sql.NullByte with null value",
			input: sql.NullByte{
				Byte:  0,
				Valid: false,
			},
		},
		{
			name: "should successfully unwrap invalid sql.NullByte with value",
			input: sql.NullByte{
				Byte:  1,
				Valid: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UnwrapByte(tc.input)

			if !tc.input.Valid {
				if result != nil {
					t.Fatalf("result should be null, got: '%v'", *result)
				}

				return
			}

			if result == nil {
				t.Fatalf("result should not be null, expected: '%v'", tc.input.Byte)
			}

			if *result != tc.input.Byte {
				t.Fatalf("result mismatch got '%v', expected: '%v'", *result, tc.input.Byte)
			}
		})
	}
}

func TestUnwrapFloat64(t *testing.T) {
	testCases := []struct {
		name  string
		input sql.NullFloat64
	}{
		{
			name: "should successfully unwrap sql.NullFloat64 with valid value",
			input: sql.NullFloat64{
				Float64: 1,
				Valid:   true,
			},
		},
		{
			name: "should successfully unwrap sql.NullFloat64 with null value",
			input: sql.NullFloat64{
				Float64: 0,
				Valid:   false,
			},
		},
		{
			name: "should successfully unwrap invalid sql.NullFloat64 with value",
			input: sql.NullFloat64{
				Float64: 1,
				Valid:   false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UnwrapFloat64(tc.input)

			if !tc.input.Valid {
				if result != nil {
					t.Fatalf("result should be null, got: '%v'", *result)
				}

				return
			}

			if result == nil {
				t.Fatalf("result should not be null, expected: '%v'", tc.input.Float64)
			}

			if *result != tc.input.Float64 {
				t.Fatalf("result mismatch got '%v', expected: '%v'", *result, tc.input.Float64)
			}
		})
	}
}

func TestUnwrapBoolean(t *testing.T) {
	testCases := []struct {
		name  string
		input sql.NullBool
	}{
		{
			name: "should successfully unwrap sql.NullBool with valid value",
			input: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		},
		{
			name: "should successfully unwrap sql.NullBool with null value",
			input: sql.NullBool{
				Bool:  false,
				Valid: false,
			},
		},
		{
			name: "should successfully unwrap invalid sql.NullBool with value",
			input: sql.NullBool{
				Bool:  true,
				Valid: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UnwrapBoolean(tc.input)

			if !tc.input.Valid {
				if result != nil {
					t.Fatalf("result should be null, got: '%v'", *result)
				}

				return
			}

			if result == nil {
				t.Fatalf("result should not be null, expected: '%v'", tc.input.Bool)
			}

			if *result != tc.input.Bool {
				t.Fatalf("result mismatch got '%v', expected: '%v'", *result, tc.input.Bool)
			}
		})
	}
}
