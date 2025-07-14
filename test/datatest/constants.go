// Package datatest provides shared constants and helpers for test data,
// including predefined UUIDs and reusable error values.
// This package is intended for use in unit tests and should not be imported in production code.
package datatest

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// FakeProductID is a fixed UUIDv4 value used for testing purposes.
	// It simulates a realistic UUID without relying on random generation.
	FakeProductID = uuid.MustParse("4e3d9f02-8a7c-4b72-b10f-3fd88e2ecfaa")

	// ErrUnexpectedDB simulates a generic database error used in test scenarios.
	ErrUnexpectedDB = errors.New("unexpected database error")
)
