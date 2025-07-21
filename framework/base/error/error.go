package error

import (
	stderr "errors"
	"fmt"

	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	MessageField  string     `json:"msg,omitempty"`
	KeyField      string     `json:"key,omitempty"`
	CodeField     int64      `json:"code,omitempty"`
	GRPCCodeField codes.Code `json:"-"`
	DebugField    string     `json:"-"`
	DataField     interface{}
	err           error
	IsRetry       bool `json:"is_retry,omitempty"`

	// Further error details
	DetailsField map[string]interface{} `json:"details,omitempty"`
	// Request ID
	RIDField string `json:"rid,omitempty"`
}

func NewErr(err error, args ...interface{}) *Error {
	eErr := &Error{
		MessageField: err.Error(),
		err:          err,
	}
	if len(args) > 0 {
		eErr.KeyField = utils.ItoString(args[0])
	}
	if len(args) > 1 {
		eErr.DataField = args[1]
	}

	if len(args) > 2 {
		retriable, ok := args[2].(bool)
		if ok {
			eErr.IsRetry = retriable
		}
	}

	return eErr
}

func (e *Error) IsRetriable() bool {
	if e == nil {
		return false
	}

	return e.IsRetry
}

func New(msg string, args ...interface{}) *Error {
	eErr := &Error{
		MessageField: msg,
		err:          errors.New(msg),
	}

	if len(args) > 0 {
		eErr.KeyField = utils.ItoString(args[0])
	}

	if len(args) > 1 {
		eErr.DataField = args[1]
	}

	if len(args) > 2 {
		retriable, ok := args[2].(bool)
		if ok {
			eErr.IsRetry = retriable
		}
	}
	return eErr
}

func (e Error) Code() int64 {
	return e.CodeField
}

func (e Error) Msg() string {
	return e.MessageField
}

func (e Error) Data() interface{} {
	return e.DataField
}

func (e Error) Key() string {
	return e.KeyField
}

func (e Error) RID() string {
	return e.RIDField
}

func (e Error) Debug() string {
	return e.DebugField
}

func (e Error) Error() string {
	return e.MessageField
}

func (e Error) Details() map[string]interface{} {
	return e.DetailsField
}

func (e Error) GRPCStatus() *status.Status {
	s := status.New(e.GRPCCodeField, e.Error())

	st := e.StackTrace()
	var stackEntries []string
	if st != nil {
		stackEntries = make([]string, len(st))
		for i, f := range st {
			stackEntries[i] = fmt.Sprintf("%+v", f)
		}
	}

	details := make([]proto.Message, 0, 3)

	if e.Debug() != "" || st != nil {
		details = append(details, &errdetails.DebugInfo{
			StackEntries: stackEntries,
			Detail:       e.Debug(),
		})
	}

	if e.Error() != "" {
		details = append(details, &errdetails.ErrorInfo{
			Reason: e.Error(),
		})
	}

	if e.RID() != "" {
		details = append(details, &errdetails.RequestInfo{
			RequestId: e.RID(),
		})
	}

	if e.GRPCCodeField == codes.InvalidArgument && e.err != nil {
		if fvs := e.fieldViolations(); len(fvs) > 0 {
			details = append(details, &errdetails.BadRequest{
				FieldViolations: fvs,
			})
		}
	}

	s, err := s.WithDetails(details...)
	if err != nil {
		panic(err)
	}

	return s
}

func (e Error) WithMsg(msg string) *Error {
	e.MessageField = msg
	return &e
}

func (e Error) WithKey(key string) *Error {
	e.KeyField = key
	return &e
}

func (e Error) WithCode(code int64) *Error {
	e.CodeField = code
	return &e
}

func (e Error) WithData(data interface{}) *Error {
	e.DataField = data
	return &e
}

func (e Error) WithDebug(debug string) *Error {
	e.DebugField = debug
	return &e
}

func (e Error) WithError(err *Error) *Error {
	err.CodeField = e.CodeField
	err.GRPCCodeField = e.GRPCCodeField
	return err
}

func (e Error) WithRID(rid string) *Error {
	e.RIDField = rid
	return &e
}

func (e *Error) Wrap(err error) {
	e.err = err
}

func (e *Error) WithTrace(err error) *Error {
	if st := stackTracer(nil); !stderr.As(e.err, &st) {
		e.Wrap(errors.WithStack(err))
	} else {
		e.Wrap(err)
	}
	return e
}

func (e Error) Is(err error) bool {
	var te Error
	switch {
	case errors.As(err, te):
		return e.CodeField == te.CodeField &&
			e.GRPCCodeField == te.GRPCCodeField
	case errors.As(err, &te):
		return e.CodeField == te.CodeField &&
			e.GRPCCodeField == te.GRPCCodeField
	default:
		return false
	}
}

// StackTrace returns the error's stack trace.
func (e *Error) StackTrace() (trace errors.StackTrace) {
	if e.err == e {
		return
	}

	if st := stackTracer(nil); stderr.As(e.err, &st) {
		trace = st.StackTrace()
	}

	return
}

// fieldViolationError is an interface implemented by proto-gen-validate.
type fieldViolationError interface {
	Field() string
	Reason() string
	Cause() error
}

type multiError interface {
	AllErrors() []error
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func rootCauses(err fieldViolationError) []fieldViolationError {
	if err == nil {
		return []fieldViolationError{}
	}

	switch e := err.Cause().(type) {
	case fieldViolationError:
		return rootCauses(e)

	case multiError:
		var causes []fieldViolationError
		for _, e := range e.AllErrors() {
			if fvErr, ok := e.(fieldViolationError); ok {
				causes = append(causes, rootCauses(fvErr)...)
			}
		}

		return causes
	}

	return []fieldViolationError{err}
}

func (e Error) fieldViolations() (fv []*errdetails.BadRequest_FieldViolation) {
	err, ok := e.err.(multiError)
	if !ok {
		return
	}

	for _, e := range err.AllErrors() {
		if fvErr, ok := e.(fieldViolationError); ok {
			// We only want to show the root cause of the error.
			for _, cause := range rootCauses(fvErr) {
				fv = append(fv, &errdetails.BadRequest_FieldViolation{
					Field:       cause.Field(),
					Description: cause.Reason(),
				})
			}
		}
	}

	return
}
