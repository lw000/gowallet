package errors

import (
	"fmt"
	"time"
)

var (
	// 参数无效
	InvalidArgument = &Error{code: 0, what: "invalid argument", detail: "invalid argument"}
)

var (
	// 非法数据包
	ErrorIllegalPacket = &Error{code: 0, what: "[sg error]", detail: "非法数据包[proto]"}
)

// Error ...
type Error struct {
	when   time.Time
	code   int32
	what   string
	detail string
}

// New ...
func New(code int32, what, detail string) *Error {
	return &Error{
		when:   time.Now(),
		code:   code,
		what:   what,
		detail: detail,
	}
}

// NewError ...
func NewError(err *Error) *Error {
	err.SetWhen(time.Now())
	return err
}

// SetWhen ...
func (e *Error) SetWhen(when time.Time) {
	e.when = when
}

// When ...
func (e Error) When() time.Time {
	return e.when
}

// WhenString ...
func (e Error) WhenString() string {
	return e.when.Format("2006-01-02 15:04:05")
}

// Detail ...
func (e Error) Detail() string {
	return e.detail
}

// What ...
func (e Error) What() string {
	return e.what
}

// Code ...
func (e Error) Code() int32 {
	return e.code
}

// Error ...
func (e *Error) Error() string {
	return e.String()
}

// String ...
func (e Error) String() string {
	return fmt.Sprintf("code:%d | what:%s | detail:%s", e.code, e.what, e.detail)
}
