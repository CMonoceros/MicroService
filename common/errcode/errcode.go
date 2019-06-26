package errcode

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

var (
	_codes = map[int]string{} // register codes.
)

// New new a ecode.Codes by int value.
// NOTE: ecode must unique in global, the New will check repeat and then panic.
func New(e int, msg string) ErrCode {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e, msg)
}

func add(e int, msg string) ErrCode {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = msg
	return ErrCode{
		code:    e,
		message: msg,
	}
}

// Codes ecode error interface which has a code & message.
type Codes interface {
	// sometimes Error return Code in string form
	// NOTE: don't use Error in monitor report even it also work for now
	Error() string
	// Code get error code.
	Code() int
	// Message get code message.
	Message() string
	//Detail get error detail,it may be nil.
	Details() []interface{}
}

// A Code is an int error code spec.
type ErrCode struct {
	code    int
	message string
}

func (e ErrCode) Error() string {
	return strconv.FormatInt(int64(e.code), 10)
}

// Code return error code
func (e ErrCode) Code() int { return e.code }

// Message return error message
func (e ErrCode) Message() string {
	return e.message
}

// Details return details.
func (e ErrCode) Details() []interface{} { return nil }

// String parse code string to error.
func String(e string) ErrCode {
	if e == "" {
		return OK
	}
	// try error string
	i, err := strconv.Atoi(e)
	if err != nil {
		return ServerErr
	}
	if _, ok := _codes[i]; !ok {
		panic(fmt.Sprintf("ecode: %d not exist", i))
	}

	return ErrCode{
		code:    i,
		message: _codes[i],
	}
}

// Cause cause from error to ecode.
func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
}

// Equal equal a and b by code int.
func Equal(a, b Codes) bool {
	if a == nil {
		a = OK
	}
	if b == nil {
		b = OK
	}
	return a.Code() == b.Code()
}

// EqualError equal error
func EqualError(code Codes, err error) bool {
	return Cause(err).Code() == code.Code()
}
