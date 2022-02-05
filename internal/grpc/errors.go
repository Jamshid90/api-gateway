package grpc

import (
	"errors"

	errorpkg "github.com/Jamshid90/api-getawey/internal/errors"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Error(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return nil
	}

	switch st.Code() {

	// error bad request
	case codes.OK:
		return nil

	// error not found
	case codes.NotFound:
		return errorpkg.ErrorNotFound

	// error validation errors
	case codes.InvalidArgument:
		errValidation := errorpkg.NewErrValidation()
		errValidation.Errors = ErrorDetails(st)
		return errValidation

	default:
		var errStr string
		for _, detail := range st.Details() {
			if errorInfo, ok := detail.(*epb.ErrorInfo); ok {
				errStr += " " + errorInfo.Reason
			}
		}
		if len(errStr) != 0 {
			return errors.New(errStr)
		}
		return errorpkg.ErrInternalServer
	}
}

func ErrorDetails(st *status.Status) map[string]string {
	var data = make(map[string]string)
	for _, detail := range st.Details() {
		if badRequest, ok := detail.(*epb.BadRequest); ok {
			for _, violation := range badRequest.GetFieldViolations() {
				data[violation.GetField()] = violation.GetDescription()
			}
		}
	}
	return data
}
