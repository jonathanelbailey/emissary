//go:build !disable_pgv

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/listener/local_ratelimit/v3/local_ratelimit.proto

package local_ratelimitv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on LocalRateLimit with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LocalRateLimit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LocalRateLimit with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LocalRateLimitMultiError,
// or nil if none found.
func (m *LocalRateLimit) ValidateAll() error {
	return m.validate(true)
}

func (m *LocalRateLimit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := LocalRateLimitValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTokenBucket() == nil {
		err := LocalRateLimitValidationError{
			field:  "TokenBucket",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTokenBucket()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "TokenBucket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "TokenBucket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTokenBucket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "TokenBucket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRuntimeEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "RuntimeEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "RuntimeEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRuntimeEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "RuntimeEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LocalRateLimitMultiError(errors)
	}

	return nil
}

// LocalRateLimitMultiError is an error wrapping multiple validation errors
// returned by LocalRateLimit.ValidateAll() if the designated constraints
// aren't met.
type LocalRateLimitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocalRateLimitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocalRateLimitMultiError) AllErrors() []error { return m }

// LocalRateLimitValidationError is the validation error returned by
// LocalRateLimit.Validate if the designated constraints aren't met.
type LocalRateLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalRateLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalRateLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalRateLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalRateLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalRateLimitValidationError) ErrorName() string { return "LocalRateLimitValidationError" }

// Error satisfies the builtin error interface
func (e LocalRateLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalRateLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalRateLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalRateLimitValidationError{}
