// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/basic/v1/operation_log.proto

package v1

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
)

// Validate checks the field values on ListOperationLogUserRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListOperationLogUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserName

	return nil
}

// ListOperationLogUserRequestValidationError is the validation error returned
// by ListOperationLogUserRequest.Validate if the designated constraints
// aren't met.
type ListOperationLogUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOperationLogUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOperationLogUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOperationLogUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOperationLogUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOperationLogUserRequestValidationError) ErrorName() string {
	return "ListOperationLogUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListOperationLogUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOperationLogUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOperationLogUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOperationLogUserRequestValidationError{}

// Validate checks the field values on ListOperationLogUserReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListOperationLogUserReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListOperationLogUserReplyValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListOperationLogUserReplyValidationError is the validation error returned by
// ListOperationLogUserReply.Validate if the designated constraints aren't met.
type ListOperationLogUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOperationLogUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOperationLogUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOperationLogUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOperationLogUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOperationLogUserReplyValidationError) ErrorName() string {
	return "ListOperationLogUserReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListOperationLogUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOperationLogUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOperationLogUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOperationLogUserReplyValidationError{}

// Validate checks the field values on ListOperationLogRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListOperationLogRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListOperationLogRequestValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	// no validation rules for CorpId

	if v, ok := interface{}(m.GetOperateStartAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListOperationLogRequestValidationError{
				field:  "OperateStartAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOperateEndAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListOperationLogRequestValidationError{
				field:  "OperateEndAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AppId

	// no validation rules for OperationName

	return nil
}

// ListOperationLogRequestValidationError is the validation error returned by
// ListOperationLogRequest.Validate if the designated constraints aren't met.
type ListOperationLogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOperationLogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOperationLogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOperationLogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOperationLogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOperationLogRequestValidationError) ErrorName() string {
	return "ListOperationLogRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListOperationLogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOperationLogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOperationLogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOperationLogRequestValidationError{}

// Validate checks the field values on ListOperationLogReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListOperationLogReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListOperationLogReplyValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	return nil
}

// ListOperationLogReplyValidationError is the validation error returned by
// ListOperationLogReply.Validate if the designated constraints aren't met.
type ListOperationLogReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOperationLogReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOperationLogReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOperationLogReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOperationLogReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOperationLogReplyValidationError) ErrorName() string {
	return "ListOperationLogReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListOperationLogReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOperationLogReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOperationLogReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOperationLogReplyValidationError{}

// Validate checks the field values on Log with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Log) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for AppId

	// no validation rules for CorpId

	// no validation rules for CorpName

	// no validation rules for UserId

	// no validation rules for UserName

	// no validation rules for OperationName

	// no validation rules for OperationType

	// no validation rules for OperationModule

	// no validation rules for Detail

	if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LogValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LogValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	// no validation rules for Reason

	return nil
}

// LogValidationError is the validation error returned by Log.Validate if the
// designated constraints aren't met.
type LogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogValidationError) ErrorName() string { return "LogValidationError" }

// Error satisfies the builtin error interface
func (e LogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogValidationError{}

// Validate checks the field values on AddRequest with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AddRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetLog()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddRequestValidationError{
				field:  "Log",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AddRequestValidationError is the validation error returned by
// AddRequest.Validate if the designated constraints aren't met.
type AddRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddRequestValidationError) ErrorName() string { return "AddRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddRequestValidationError{}

// Validate checks the field values on AddReply with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AddReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// AddReplyValidationError is the validation error returned by
// AddReply.Validate if the designated constraints aren't met.
type AddReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddReplyValidationError) ErrorName() string { return "AddReplyValidationError" }

// Error satisfies the builtin error interface
func (e AddReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddReplyValidationError{}

// Validate checks the field values on ListOperationLogUserReply_User with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListOperationLogUserReply_User) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	// no validation rules for UserName

	return nil
}

// ListOperationLogUserReply_UserValidationError is the validation error
// returned by ListOperationLogUserReply_User.Validate if the designated
// constraints aren't met.
type ListOperationLogUserReply_UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOperationLogUserReply_UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOperationLogUserReply_UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOperationLogUserReply_UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOperationLogUserReply_UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOperationLogUserReply_UserValidationError) ErrorName() string {
	return "ListOperationLogUserReply_UserValidationError"
}

// Error satisfies the builtin error interface
func (e ListOperationLogUserReply_UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOperationLogUserReply_User.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOperationLogUserReply_UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOperationLogUserReply_UserValidationError{}