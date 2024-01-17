package sqlmap

import (
	"bytes"
	"testing"

	"github.com/google/uuid"
)

// compareUUID ensures the uuid.UUID pointer and uuid.NullUUID are equal.
func compareUUID(t *testing.T, expected *uuid.UUID, actual uuid.NullUUID) {
	t.Helper()

	if expected == nil {
		if actual.Valid {
			t.Fatalf("expected is nil, actual should not be valid")
		}

		return
	}

	if !bytes.Equal(expected[:], actual.UUID[:]) {
		t.Fatalf("result: '%v' does not equal expected: '%v'", actual.UUID, *expected)
	}
}

func TestNullUUID(t *testing.T) {
	t.Run("successfully handle non-pointer UUID values ", func(t *testing.T) {
		v := uuid.New()

		compareUUID(t, &v, NullUUID(v))
	})

	t.Run("successfully handle null UUID pointers", func(t *testing.T) {
		var v *uuid.UUID

		compareUUID(t, v, NullUUID(v))
	})

	t.Run("successfully handle non-null UUID pointers", func(t *testing.T) {
		v := uuid.New()

		compareUUID(t, &v, NullUUID(&v))
	})
}

func TestUnwrapUUID(t *testing.T) {
	testCases := []struct {
		name  string
		input uuid.NullUUID
	}{
		{
			name: "should successfully unwrap uuid.NullUUID with valid value",
			input: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: true,
			},
		},
		{
			name: "should successfully unwrap uuid.NullUUID with null value",
			input: uuid.NullUUID{
				UUID:  uuid.Nil,
				Valid: false,
			},
		},
		{
			name: "should successfully unwrap invalid uuid.NullUUID with value",
			input: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UnwrapUUID(tc.input)

			isZero := bytes.Equal(uuid.Nil[:], result[:])

			if !tc.input.Valid {
				if !isZero {
					t.Fatalf("result should be zero value, got: '%v'", result)
				}

				return
			}

			if isZero {
				t.Fatalf("result should not be null, expected: '%v'", tc.input.UUID)
			}

			if !bytes.Equal(tc.input.UUID[:], result[:]) {
				t.Fatalf("result mismatch got '%v', expected: '%v'", result, tc.input.UUID)
			}
		})
	}
}

func TestUnwrapUUIDPtr(t *testing.T) {
	testCases := []struct {
		name  string
		input uuid.NullUUID
	}{
		{
			name: "should successfully unwrap uuid.NullUUID with valid value",
			input: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: true,
			},
		},
		{
			name: "should successfully unwrap uuid.NullUUID with null value",
			input: uuid.NullUUID{
				UUID:  uuid.Nil,
				Valid: false,
			},
		},
		{
			name: "should successfully unwrap invalid uuid.NullUUID with value",
			input: uuid.NullUUID{
				UUID:  uuid.New(),
				Valid: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UnwrapUUIDPtr(tc.input)

			if !tc.input.Valid {
				if result != nil {
					t.Fatalf("result should be null, got: '%v'", result)
				}

				return
			}

			if result == nil {
				t.Fatalf("result should not be null, expected: '%v'", tc.input.UUID)
			}

			if !bytes.Equal(tc.input.UUID[:], result[:]) {
				t.Fatalf("result mismatch got '%v', expected: '%v'", result, tc.input.UUID)
			}
		})
	}
}
