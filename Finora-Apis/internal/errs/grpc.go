package errs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPC preserves gRPC status errors; unknown errors become Internal without leaking details.
func GRPC(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := status.FromError(err); ok {
		return err
	}
	return status.Error(codes.Internal, "internal error")
}
