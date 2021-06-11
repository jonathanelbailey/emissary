// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/matcher/v4alpha/regex.proto

package envoy_type_matcher_v4alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _regex_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on RegexMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RegexMatcher) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRegex()) < 1 {
		return RegexMatcherValidationError{
			field:  "Regex",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.EngineType.(type) {

	case *RegexMatcher_GoogleRe2:

		if m.GetGoogleRe2() == nil {
			return RegexMatcherValidationError{
				field:  "GoogleRe2",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetGoogleRe2()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegexMatcherValidationError{
					field:  "GoogleRe2",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return RegexMatcherValidationError{
			field:  "EngineType",
			reason: "value is required",
		}

	}

	return nil
}

// RegexMatcherValidationError is the validation error returned by
// RegexMatcher.Validate if the designated constraints aren't met.
type RegexMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegexMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegexMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegexMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegexMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegexMatcherValidationError) ErrorName() string { return "RegexMatcherValidationError" }

// Error satisfies the builtin error interface
func (e RegexMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegexMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegexMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegexMatcherValidationError{}

// Validate checks the field values on RegexMatchAndSubstitute with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegexMatchAndSubstitute) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPattern() == nil {
		return RegexMatchAndSubstituteValidationError{
			field:  "Pattern",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPattern()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegexMatchAndSubstituteValidationError{
				field:  "Pattern",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Substitution

	return nil
}

// RegexMatchAndSubstituteValidationError is the validation error returned by
// RegexMatchAndSubstitute.Validate if the designated constraints aren't met.
type RegexMatchAndSubstituteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegexMatchAndSubstituteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegexMatchAndSubstituteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegexMatchAndSubstituteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegexMatchAndSubstituteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegexMatchAndSubstituteValidationError) ErrorName() string {
	return "RegexMatchAndSubstituteValidationError"
}

// Error satisfies the builtin error interface
func (e RegexMatchAndSubstituteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegexMatchAndSubstitute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegexMatchAndSubstituteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegexMatchAndSubstituteValidationError{}

// Validate checks the field values on RegexMatcher_GoogleRE2 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegexMatcher_GoogleRE2) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedMaxProgramSize()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegexMatcher_GoogleRE2ValidationError{
				field:  "HiddenEnvoyDeprecatedMaxProgramSize",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RegexMatcher_GoogleRE2ValidationError is the validation error returned by
// RegexMatcher_GoogleRE2.Validate if the designated constraints aren't met.
type RegexMatcher_GoogleRE2ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegexMatcher_GoogleRE2ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegexMatcher_GoogleRE2ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegexMatcher_GoogleRE2ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegexMatcher_GoogleRE2ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegexMatcher_GoogleRE2ValidationError) ErrorName() string {
	return "RegexMatcher_GoogleRE2ValidationError"
}

// Error satisfies the builtin error interface
func (e RegexMatcher_GoogleRE2ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegexMatcher_GoogleRE2.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegexMatcher_GoogleRE2ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegexMatcher_GoogleRE2ValidationError{}