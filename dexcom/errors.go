package dexcom

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRegion = errors.New("invalid region")
	ErrNoReadings    = errors.New("no glucose readings available")
)

var (
	ErrSessionExpired     = errors.New("dexcom: session expired")
	ErrSessionNotFound    = errors.New("dexcom: session id not found")
	ErrSessionInvalid     = errors.New("dexcom: session id invalid")
	ErrAccountAuthFailed  = errors.New("dexcom: account authentication failed")
	ErrAccountMaxAttempts = errors.New("dexcom: account authentication max attempts exceeded")
	ErrArgumentInvalid    = errors.New("dexcom: invalid argument")
	ErrServerInvalidJSON  = errors.New("dexcom: invalid server JSON")
	ErrServerUnexpected   = errors.New("dexcom: unexpected server error")
)

type DexcomError struct {
	Kind string // e.g. "SessionError", "AccountError"
	Err  error  // underlying sentinel error
	Msg  string // optional extra context
}

func (e *DexcomError) Error() string {
	if e.Msg != "" {
		return fmt.Sprintf("%s: %s (%v)", e.Kind, e.Msg, e.Err)
	}
	return fmt.Sprintf("%s: %v", e.Kind, e.Err)
}

func (e *DexcomError) Unwrap() error {
	return e.Err
}
