// Конвертер обычных ошибок в grpc формат
package service

import (
	"errors"

	pkg_errors "github.com/dragun-igor/img-strg/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCError interface {
	Error() string
	GRPCStatus() *status.Status
	Unwrap() error
}

type grpcError struct {
	err    error
	status *status.Status
}

func convert(err error) GRPCError {
	if v, ok := err.(GRPCError); ok { //nolint:errorlint
		return v
	}
	switch {
	case errors.Is(err, pkg_errors.ErrFileNotFound):
		return newGRPCError(err, codes.InvalidArgument)
	default:
		return newGRPCError(err, codes.Internal)
	}
}

func newGRPCError(err error, code codes.Code) grpcError {
	return grpcError{
		err:    err,
		status: status.New(code, err.Error()),
	}
}

func (e grpcError) Error() string {
	return e.err.Error()
}

func (e grpcError) GRPCStatus() *status.Status {
	return e.status
}

func (e grpcError) Unwrap() error {
	return e.err
}
