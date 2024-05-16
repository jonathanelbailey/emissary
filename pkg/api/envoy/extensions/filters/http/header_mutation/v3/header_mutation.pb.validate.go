//go:build !disable_pgv

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/header_mutation/v3/header_mutation.proto

package header_mutationv3

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

// Validate checks the field values on Mutations with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Mutations) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Mutations with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MutationsMultiError, or nil
// if none found.
func (m *Mutations) ValidateAll() error {
	return m.validate(true)
}

func (m *Mutations) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRequestMutations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MutationsValidationError{
						field:  fmt.Sprintf("RequestMutations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MutationsValidationError{
						field:  fmt.Sprintf("RequestMutations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MutationsValidationError{
					field:  fmt.Sprintf("RequestMutations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetResponseMutations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MutationsValidationError{
						field:  fmt.Sprintf("ResponseMutations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MutationsValidationError{
						field:  fmt.Sprintf("ResponseMutations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MutationsValidationError{
					field:  fmt.Sprintf("ResponseMutations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MutationsMultiError(errors)
	}

	return nil
}

// MutationsMultiError is an error wrapping multiple validation errors returned
// by Mutations.ValidateAll() if the designated constraints aren't met.
type MutationsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MutationsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MutationsMultiError) AllErrors() []error { return m }

// MutationsValidationError is the validation error returned by
// Mutations.Validate if the designated constraints aren't met.
type MutationsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MutationsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MutationsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MutationsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MutationsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MutationsValidationError) ErrorName() string { return "MutationsValidationError" }

// Error satisfies the builtin error interface
func (e MutationsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMutations.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MutationsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MutationsValidationError{}

// Validate checks the field values on HeaderMutationPerRoute with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HeaderMutationPerRoute) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HeaderMutationPerRoute with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HeaderMutationPerRouteMultiError, or nil if none found.
func (m *HeaderMutationPerRoute) ValidateAll() error {
	return m.validate(true)
}

func (m *HeaderMutationPerRoute) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMutations()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HeaderMutationPerRouteValidationError{
					field:  "Mutations",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HeaderMutationPerRouteValidationError{
					field:  "Mutations",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMutations()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeaderMutationPerRouteValidationError{
				field:  "Mutations",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HeaderMutationPerRouteMultiError(errors)
	}

	return nil
}

// HeaderMutationPerRouteMultiError is an error wrapping multiple validation
// errors returned by HeaderMutationPerRoute.ValidateAll() if the designated
// constraints aren't met.
type HeaderMutationPerRouteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HeaderMutationPerRouteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HeaderMutationPerRouteMultiError) AllErrors() []error { return m }

// HeaderMutationPerRouteValidationError is the validation error returned by
// HeaderMutationPerRoute.Validate if the designated constraints aren't met.
type HeaderMutationPerRouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderMutationPerRouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderMutationPerRouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderMutationPerRouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderMutationPerRouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderMutationPerRouteValidationError) ErrorName() string {
	return "HeaderMutationPerRouteValidationError"
}

// Error satisfies the builtin error interface
func (e HeaderMutationPerRouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderMutationPerRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderMutationPerRouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderMutationPerRouteValidationError{}

// Validate checks the field values on HeaderMutation with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HeaderMutation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HeaderMutation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HeaderMutationMultiError,
// or nil if none found.
func (m *HeaderMutation) ValidateAll() error {
	return m.validate(true)
}

func (m *HeaderMutation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMutations()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HeaderMutationValidationError{
					field:  "Mutations",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HeaderMutationValidationError{
					field:  "Mutations",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMutations()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeaderMutationValidationError{
				field:  "Mutations",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MostSpecificHeaderMutationsWins

	if len(errors) > 0 {
		return HeaderMutationMultiError(errors)
	}

	return nil
}

// HeaderMutationMultiError is an error wrapping multiple validation errors
// returned by HeaderMutation.ValidateAll() if the designated constraints
// aren't met.
type HeaderMutationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HeaderMutationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HeaderMutationMultiError) AllErrors() []error { return m }

// HeaderMutationValidationError is the validation error returned by
// HeaderMutation.Validate if the designated constraints aren't met.
type HeaderMutationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderMutationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderMutationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderMutationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderMutationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderMutationValidationError) ErrorName() string { return "HeaderMutationValidationError" }

// Error satisfies the builtin error interface
func (e HeaderMutationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderMutation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderMutationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderMutationValidationError{}
