package sqlmap

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
	"time"
)

// compareString ensures the string pointer and sql.NullString are equal.
func compareString(t *testing.T, expected *string, actual sql.NullString) {
	t.Helper()

	if expected == nil {
		if actual.Valid {
			t.Fatalf("expected is nil, actual should not be valid")
		}

		return
	}

	if strings.Compare(*expected, actual.String) != 0 {
		t.Fatalf("result: '%s' does not equal expected: '%s'", actual.String, *expected)
	}
}

// compareInt64 ensures the int64 pointer and sql.NullInt64 are equal.
func compareInt64(t *testing.T, expected *int64, actual sql.NullInt64) {
	t.Helper()

	if expected == nil {
		if actual.Valid {
			t.Fatalf("expected is nil, actual should not be valid")
		}

		return
	}

	if *expected != actual.Int64 {
		t.Fatalf("result: '%d' does not equal expected: '%d'", actual.Int64, *expected)
	}
}

// compareInt32 ensures the int32 pointer and sql.NullInt32 are equal.
func compareInt32(t *testing.T, expected *int32, actual sql.NullInt32) {
	t.Helper()

	if expected == nil {
		if actual.Valid {
			t.Fatalf("expected is nil, actual should not be valid")
		}

		return
	}

	if *expected != actual.Int32 {
		t.Fatalf("result: '%d' does not equal expected: '%d'", actual.Int32, *expected)
	}
}

// compareByte ensures the byte pointer and sql.NullByte are equal.
func compareByte(t *testing.T, expected *byte, actual sql.NullByte) {
	t.Helper()

	if expected == nil {
		if actual.Valid {
			t.Fatalf("expected is nil, actual should not be valid")
		}

		return
	}

	if *expected != actual.Byte {
		t.Fatalf("result: '%v' does not equal expected: '%v'", actual.Byte, *expected)
	}
}

// compareInt16 ensures the int16 pointer and sql.NullInt16 are equal.
func compareInt16(t *testing.T, expected *int16, actual sql.NullInt16) {
	t.Helper()

	if expected == nil {
		if actual.Valid {
			t.Fatalf("expected is nil, actual should not be valid")
		}

		return
	}

	if *expected != actual.Int16 {
		t.Fatalf("result: '%d' does not equal expected: '%d'", actual.Int16, *expected)
	}
}

// compareFloat64 ensures the float64 pointer and sql.NullFloat64 are equal.
func compareFloat64(t *testing.T, expected *float64, actual sql.NullFloat64) {
	t.Helper()

	if expected == nil {
		if actual.Valid {
			t.Fatalf("expected is nil, actual should not be valid")
		}

		return
	}

	if *expected != actual.Float64 {
		t.Fatalf("result: '%f' does not equal expected: '%f'", actual.Float64, *expected)
	}
}

// compareBoolean ensures the bool pointer and sql.NullBool are equal.
func compareBoolean(t *testing.T, expected *bool, actual sql.NullBool) {
	t.Helper()

	if expected == nil {
		if actual.Valid {
			t.Fatalf("expected is nil, actual should not be valid")
		}

		return
	}

	if *expected != actual.Bool {
		t.Fatalf("result: '%v' does not equal expected: '%v'", actual.Bool, *expected)
	}
}

// compareTime ensures the time.Time pointer and sql.NullTime are equal.
func compareTime(t *testing.T, expected *time.Time, actual sql.NullTime) {
	t.Helper()

	if expected == nil {
		if actual.Valid {
			t.Fatalf("expected is nil, actual should not be valid")
		}

		return
	}

	if !expected.Equal(actual.Time) {
		t.Fatalf("result: '%v' does not equal expected: '%v'", actual.Time, *expected)
	}
}

func TestNullString(t *testing.T) {
	t.Run("successfully handle non-pointer string values", func(t *testing.T) {
		v := "testing123"

		compareString(t, &v, NullString(v))
	})

	t.Run("successfully handle null string pointers", func(t *testing.T) {
		var s *string

		compareString(t, s, NullString(s))
	})

	t.Run("successfully handle non-null string pointers", func(t *testing.T) {
		s := "foo"

		compareString(t, &s, NullString(&s))
	})
}

func TestNullInt64(t *testing.T) {
	t.Run("successfully handle non-pointer int64 values ", func(t *testing.T) {
		var v int64 = 42

		compareInt64(t, &v, NullInt64(v))
	})

	t.Run("successfully handle null int64 pointers", func(t *testing.T) {
		var v *int64

		compareInt64(t, v, NullInt64(v))
	})

	t.Run("successfully handle non-null int64 pointers", func(t *testing.T) {
		var v int64 = 42

		compareInt64(t, &v, NullInt64(&v))
	})
}

func TestNullInt32(t *testing.T) {
	t.Run("successfully handle non-pointer int32 values ", func(t *testing.T) {
		var v int32 = 11

		compareInt32(t, &v, NullInt32(v))
	})

	t.Run("successfully handle null int32 pointers", func(t *testing.T) {
		var v *int32

		compareInt32(t, v, NullInt32(v))
	})

	t.Run("successfully handle non-null int32 pointers", func(t *testing.T) {
		var v int32 = 11

		compareInt32(t, &v, NullInt32(&v))
	})
}

func TestNullInt16(t *testing.T) {
	t.Run("successfully handle non-pointer int16 values ", func(t *testing.T) {
		var v int16 = 11

		compareInt16(t, &v, NullInt16(v))
	})

	t.Run("successfully handle null int16 pointers", func(t *testing.T) {
		var v *int16

		compareInt16(t, v, NullInt16(v))
	})

	t.Run("successfully handle non-null int16 pointers", func(t *testing.T) {
		var v int16 = 11

		compareInt16(t, &v, NullInt16(&v))
	})
}

func TestNullByte(t *testing.T) {
	t.Run("successfully handle non-pointer byte values ", func(t *testing.T) {
		var v byte = 1

		compareByte(t, &v, NullByte(v))
	})

	t.Run("successfully handle null byte pointers", func(t *testing.T) {
		var v *byte

		compareByte(t, v, NullByte(v))
	})

	t.Run("successfully handle non-null byte pointers", func(t *testing.T) {
		var v byte = 1

		compareByte(t, &v, NullByte(&v))
	})
}

func TestNullFloat64(t *testing.T) {
	t.Run("successfully handle non-pointer float64 values ", func(t *testing.T) {
		var v float64 = 1

		compareFloat64(t, &v, NullFloat64(v))
	})

	t.Run("successfully handle null float64 pointers", func(t *testing.T) {
		var v *float64

		compareFloat64(t, v, NullFloat64(v))
	})

	t.Run("successfully handle non-null float64 pointers", func(t *testing.T) {
		var v float64 = 1

		compareFloat64(t, &v, NullFloat64(&v))
	})
}

func TestNullBoolean(t *testing.T) {
	t.Run("successfully handle non-pointer bool values ", func(t *testing.T) {
		v := true

		compareBoolean(t, &v, NullBoolean(v))
	})

	t.Run("successfully handle null bool pointers", func(t *testing.T) {
		var v *bool

		compareBoolean(t, v, NullBoolean(v))
	})

	t.Run("successfully handle non-null bool pointers", func(t *testing.T) {
		v := true

		compareBoolean(t, &v, NullBoolean(&v))
	})
}

func TestNullTime(t *testing.T) {
	t.Run("successfully handle non-pointer time.Time values ", func(t *testing.T) {
		v := time.Now()

		compareTime(t, &v, NullTime(v))
	})

	t.Run("successfully handle null time.Time pointers", func(t *testing.T) {
		var v *time.Time

		compareTime(t, v, NullTime(v))
	})

	t.Run("successfully handle non-null time.Time pointers", func(t *testing.T) {
		v := time.Now()

		compareTime(t, &v, NullTime(&v))
	})
}

func TestSerial(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput uint64
		expectedError  error
	}{
		{
			name:           "should parse serialized integer",
			input:          "42",
			expectedOutput: 42,
			expectedError:  nil,
		},
		{
			name:           "should parse random string",
			input:          "foobar",
			expectedOutput: 0,
			expectedError:  fmt.Errorf("sqlmap: failed to parse '%s' as unsigned int", "foobar"),
		},
		{
			name:           "should fail to parse float",
			input:          "1.1111",
			expectedOutput: 0,
			expectedError:  fmt.Errorf("sqlmap: failed to parse '%s' as unsigned int", "1.1111"),
		},
		{
			name:           "should fail to parse negative integer",
			input:          "-1",
			expectedOutput: 0,
			expectedError:  fmt.Errorf("sqlmap: failed to parse '%s' as unsigned int", "-1"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Serial(tc.input)
			if err != nil {
				if tc.expectedError != nil {
					// Assert the errors match.
					if strings.Compare(err.Error(), tc.expectedError.Error()) != 0 {
						t.Errorf("error mismatch; got '%v', expected: '%v'", err, tc.expectedError)
					}

					return
				}

				t.Fatalf("function should not return error, got error: '%v'", err)
			}

			if result != tc.expectedOutput {
				t.Fatalf("result: '%d' does not equal expected: '%d'", result, tc.expectedOutput)
			}
		})
	}
}
