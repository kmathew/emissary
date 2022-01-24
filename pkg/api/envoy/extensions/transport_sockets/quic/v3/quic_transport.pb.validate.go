// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/quic/v3/quic_transport.proto

package quicv3

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

// Validate checks the field values on QuicDownstreamTransport with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuicDownstreamTransport) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuicDownstreamTransport with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuicDownstreamTransportMultiError, or nil if none found.
func (m *QuicDownstreamTransport) ValidateAll() error {
	return m.validate(true)
}

func (m *QuicDownstreamTransport) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetDownstreamTlsContext() == nil {
		err := QuicDownstreamTransportValidationError{
			field:  "DownstreamTlsContext",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDownstreamTlsContext()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuicDownstreamTransportValidationError{
					field:  "DownstreamTlsContext",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuicDownstreamTransportValidationError{
					field:  "DownstreamTlsContext",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDownstreamTlsContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicDownstreamTransportValidationError{
				field:  "DownstreamTlsContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return QuicDownstreamTransportMultiError(errors)
	}
	return nil
}

// QuicDownstreamTransportMultiError is an error wrapping multiple validation
// errors returned by QuicDownstreamTransport.ValidateAll() if the designated
// constraints aren't met.
type QuicDownstreamTransportMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuicDownstreamTransportMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuicDownstreamTransportMultiError) AllErrors() []error { return m }

// QuicDownstreamTransportValidationError is the validation error returned by
// QuicDownstreamTransport.Validate if the designated constraints aren't met.
type QuicDownstreamTransportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuicDownstreamTransportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuicDownstreamTransportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuicDownstreamTransportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuicDownstreamTransportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuicDownstreamTransportValidationError) ErrorName() string {
	return "QuicDownstreamTransportValidationError"
}

// Error satisfies the builtin error interface
func (e QuicDownstreamTransportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuicDownstreamTransport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuicDownstreamTransportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuicDownstreamTransportValidationError{}

// Validate checks the field values on QuicUpstreamTransport with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuicUpstreamTransport) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuicUpstreamTransport with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuicUpstreamTransportMultiError, or nil if none found.
func (m *QuicUpstreamTransport) ValidateAll() error {
	return m.validate(true)
}

func (m *QuicUpstreamTransport) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUpstreamTlsContext() == nil {
		err := QuicUpstreamTransportValidationError{
			field:  "UpstreamTlsContext",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetUpstreamTlsContext()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuicUpstreamTransportValidationError{
					field:  "UpstreamTlsContext",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuicUpstreamTransportValidationError{
					field:  "UpstreamTlsContext",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpstreamTlsContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicUpstreamTransportValidationError{
				field:  "UpstreamTlsContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return QuicUpstreamTransportMultiError(errors)
	}
	return nil
}

// QuicUpstreamTransportMultiError is an error wrapping multiple validation
// errors returned by QuicUpstreamTransport.ValidateAll() if the designated
// constraints aren't met.
type QuicUpstreamTransportMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuicUpstreamTransportMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuicUpstreamTransportMultiError) AllErrors() []error { return m }

// QuicUpstreamTransportValidationError is the validation error returned by
// QuicUpstreamTransport.Validate if the designated constraints aren't met.
type QuicUpstreamTransportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuicUpstreamTransportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuicUpstreamTransportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuicUpstreamTransportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuicUpstreamTransportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuicUpstreamTransportValidationError) ErrorName() string {
	return "QuicUpstreamTransportValidationError"
}

// Error satisfies the builtin error interface
func (e QuicUpstreamTransportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuicUpstreamTransport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuicUpstreamTransportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuicUpstreamTransportValidationError{}
